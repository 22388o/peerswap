package test

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/invoicesrpc"
	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/sputn1ck/peerswap/lightning"
	"github.com/sputn1ck/peerswap/testframework"
	"github.com/stretchr/testify/suite"
)

type LndHodlTest struct {
	suite.Suite
	assertions *AssertionCounter

	bitcoind     *testframework.BitcoinNode
	lnds         []*testframework.LndNode
	lighntningds []*testframework.CLightningNode
	scid         string
	lcid         uint64

	channelBalances []uint64
	walletBalances  []uint64
}

// TestLndLndSwapsOnBitcoin runs all integration tests concerning
// bitcoin backend and lnd-lnd operation.
func TestLndHodl(t *testing.T) {
	// Long running tests only run in integration test mode.
	testEnabled := os.Getenv("RUN_INTEGRATION_TESTS")
	if testEnabled == "" {
		t.Skip("set RUN_INTEGRATION_TESTS to run this test")
	}
	suite.Run(t, new(LndHodlTest))
}

func (suite *LndHodlTest) SetupSuite() {
	t := suite.T()
	suite.assertions = &AssertionCounter{}

	// Settings
	// Inital channel capacity
	var fundAmt = uint64(math.Pow(10, 7))

	// Get PeerSwap plugin path and test dir
	testDir := t.TempDir()

	// Setup nodes (1 bitcoind, 2 lightningd, 2 peerswapd)
	bitcoind, err := testframework.NewBitcoinNode(testDir, 1)
	if err != nil {
		t.Fatalf("could not create bitcoind %v", err)
	}
	t.Cleanup(bitcoind.Kill)

	var lnds []*testframework.LndNode
	for i := 1; i <= 2; i++ {
		lightningd, err := testframework.NewLndNode(testDir, bitcoind, i)
		if err != nil {
			t.Fatalf("could not create liquidd %v", err)
		}
		t.Cleanup(lightningd.Kill)

		lnds = append(lnds, lightningd)
	}

	// Start nodes
	err = bitcoind.Run(true)
	if err != nil {
		t.Fatalf("bitcoind.Run() got err %v", err)
	}

	for _, lnd := range lnds {
		err = lnd.Run(true, true)
		if err != nil {
			t.Fatalf("lightningd.Run() got err %v", err)
		}
	}

	// Setup channel ([0] fundAmt(10^7) ---- 0 [1])
	scid, err := lnds[1].OpenChannel(lnds[0], fundAmt, true, true, true)
	if err != nil {
		t.Fatalf("lightingds[0].OpenChannel() %v", err)
	}

	lcid, err := lnds[1].ChanIdFromScid(scid)
	if err != nil {
		t.Fatalf("lightingds[0].ChanIdFromScid() %v", err)
	}

	suite.bitcoind = bitcoind
	suite.lnds = lnds
	suite.scid = scid
	suite.lcid = lcid
}

func (suite *LndHodlTest) BeforeTest(suiteName, testName string) {
	fmt.Printf("===RUN %s/%s\n", suiteName, testName)
	// make shure we dont have pending balances
	var err error
	for _, lightningd := range suite.lnds {
		err = testframework.WaitForWithErr(func() (bool, error) {
			hasPending, err := lightningd.HasPendingHtlcOnChannel(suite.scid)
			return !hasPending, err
		}, testframework.TIMEOUT)
	}
	suite.Require().NoError(err)

	var channelBalances []uint64
	var walletBalances []uint64
	for _, lightningd := range suite.lnds {
		b, err := lightningd.GetBtcBalanceSat()
		suite.Require().NoError(err)
		walletBalances = append(walletBalances, b)

		cb, err := lightningd.GetChannelBalanceSat(suite.scid)
		suite.Require().NoError(err)
		channelBalances = append(channelBalances, cb)
	}

	suite.channelBalances = channelBalances
	suite.walletBalances = walletBalances
}

func (suite *LndHodlTest) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	var head = "FAIL"
	if stats.Passed() {
		head = "PASS"
	}
	fmt.Printf("--- %s: %s (%.2fs)\n", head, suiteName, stats.End.Sub(stats.Start).Seconds())
	for _, tStats := range stats.TestStats {
		var head = "FAIL"
		if tStats.Passed {
			head = "PASS"
		}
		fmt.Printf("\t--- %s: %s (%.2fs)\n", head, tStats.TestName, tStats.End.Sub(tStats.Start).Seconds())
	}
}

func (suite *LndHodlTest) Test_HodlInvoice() {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	lightningds := suite.lnds
	bitcoind := suite.bitcoind

	// get ctv
	listChannelsRes, err := lightningds[0].Rpc.ListChannels(ctx, &lnrpc.ListChannelsRequest{})
	suite.Require().NoError(err)
	getChanInfoRes, err := lightningds[1].Rpc.GetChanInfo(ctx, &lnrpc.ChanInfoRequest{ChanId: listChannelsRes.Channels[0].ChanId})
	suite.Require().NoError(err)
	timelock := getChanInfoRes.Node1Policy.TimeLockDelta
	log.Printf("peerswap timelock %v", timelock)
	// create hodl invoice
	invoiceClient := invoicesrpc.NewInvoicesClient(lightningds[0].Conn)

	preimage, err := lightning.GetPreimage()
	suite.Require().NoError(err)
	phash := preimage.Hash()

	holdInvRes, err := invoiceClient.AddHoldInvoice(ctx, &invoicesrpc.AddHoldInvoiceRequest{
		Value:      1000,
		Hash:       phash[:],
		CltvExpiry: 400,
	})
	suite.Require().NoError(err)
	listenHodlInv, err := invoiceClient.SubscribeSingleInvoice(ctx, &invoicesrpc.SubscribeSingleInvoiceRequest{RHash: phash[:]})
	suite.Require().NoError(err)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := listenHodlInv.Recv()
				if err != nil && strings.Contains(err.Error(), "EOF") {
					return
				}
				suite.Require().NoError(err)
				log.Printf("\n peerswap hodl message %v", msg)
			}
		}
	}()

	paymentRes, err := lightningds[1].RpcV2.SendPaymentV2(ctx, &routerrpc.SendPaymentRequest{
		PaymentRequest: holdInvRes.PaymentRequest,
		TimeoutSeconds: 3600 * 24 * 7,
	})
	suite.Require().NoError(err)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := paymentRes.Recv()
				if err != nil && strings.Contains(err.Error(), "EOF") {
					return
				}
				suite.Require().NoError(err)
				log.Printf("\n peerswap payment message %v", msg)
				if msg.Status == lnrpc.Payment_SUCCEEDED {
					return
				}
			}
		}
	}()

	time.Sleep(5 * time.Second)
	log.Printf("\n peerswap payment res %v", paymentRes)

	//err = bitcoind.GenerateBlocks(int(timelock - 5))
	err = bitcoind.GenerateBlocks(390)
	suite.Require().NoError(err)
	listChannelsRes, err = lightningds[0].Rpc.ListChannels(ctx, &lnrpc.ListChannelsRequest{})
	suite.Require().NoError(err)
	log.Printf("\n speerswap channels %v", listChannelsRes.Channels[0].PendingHtlcs)

	time.Sleep(10 * time.Second)
	settleRes, err := invoiceClient.SettleInvoice(ctx, &invoicesrpc.SettleInvoiceMsg{Preimage: preimage[:]})
	suite.Require().NoError(err)
	log.Printf("\n peerswap settleres %v", settleRes)

	time.Sleep(5 * time.Second)
	listChannelsRes, err = lightningds[0].Rpc.ListChannels(ctx, &lnrpc.ListChannelsRequest{})
	suite.Require().NoError(err)
	log.Printf("\n speerswap channels %v", listChannelsRes.Channels[0])
}
