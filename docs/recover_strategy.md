# Recover Strategies

## Overview
This document describes the recover strategies when a peerswap node restarts on an active swap. Strategies divert based on the role and the swap type. The recovery strategies are listed by the current state of the fsm.

## Table of Contents
  * [Swap In - Sender](#swap-in---sender)
    * [State_SwapInSender_CreateSwap](#state_swapinsender_createswap)
    * [State_SwapInSender_SendRequest](#state_swapinsender_sendrequest)
    * [State_SwapInSender_AwaitAgreement](#state_swapinsender_awaitagreement)
    * [State_SwapInSender_BroadcastOpeningTx](#state_swapinsender_broadcastopeningtx)
    * [State_SwapInSender_SendTxBroadcastedMessage](#state_swapinsender_sendtxbroadcastedmessage)
    * [State_SwapInSender_AwaitClaimPayment](#state_swapinsender_awaitclaimpayment)
    * [State_SwapInSender_ClaimSwapCoop](#state_swapinsender_claimswapcoop)
    * [State_SwapInSender_ClaimSwapCsv](#state_swapinsender_claimswapcsv)
  * [Swap In - Receiver](#swap-in---receiver)
    * [State_SwapInReceiver_CreateSwap](#state_swapinreceiver_createswap)
    * [State_SwapInReceiver_SendAgreement](#state_swapinreceiver_sendagreement)
    * [State_SwapInReceiver_AwaitTxBroadcastedMessage](#state_swapinreceiver_awaittxbroadcastedmessage)
    * [State_SwapInReceiver_AwaitTxConfirmation](#state_swapinreceiver_awaittxconfirmation)
    * [State_SwapInReceiver_ValidateTxAndPayClaimInvoice](#state_swapinreceiver_validatetxandpayclaiminvoice)
    * [State_SwapInReceiver_BuildSigHash](#state_swapinreceiver_buildsighash)
    * [State_SwapInReceiver_ClaimSwap](#state_swapinreceiver_claimswap)
    * [State_SwapInReceiver_SendCoopClose](#state_swapinreceiver_sendcoopclose)
  * [Swap Out - Sender](#swap-out---sender)
    * [State_SwapOutSender_CreateSwap](#state_swapoutsender_createswap)
    * [State_SwapOutSender_SendRequest](#state_swapoutsender_sendrequest)
    * [State_SwapOutSender_AwaitFeeResponse](#state_swapoutsender_awaitfeeresponse)
    * [State_SwapOutSender_PayFeeInvoice](#state_swapoutsender_payfeeinvoice)
    * [State_SwapOutSender_AwaitTxBroadcastedMessage](#state_swapoutsender_awaittxbroadcastedmessage)
    * [State_SwapOutSender_ValidateTxAndPayClaimInvoice](#state_swapoutsender_validatetxandpayclaiminvoice)
    * [State_SwapOutSender_ClaimSwap](#state_swapoutsender_claimswap)
    * [State_SwapOutSender_BuildSigHash](#state_swapoutsender_buildsighash)
    * [State_SwapOutSender_SendCoopClose](#state_swapoutsender_sendcoopclose)
  * [Swap Out - Receiver](#swap-out---receiver)
    * [State_SwapOutReceiver_CreateSwap](#state_swapoutreceiver_createswap)
    * [State_SwapOutReceiver_SendFeeInvoice](#state_swapoutreceiver_sendfeeinvoice)
    * [State_SwapOutReceiver_AwaitFeeInvoicePayment](#state_swapoutreceiver_awaitfeeinvoicepayment)
    * [State_SwapOutReceiver_BroadcastOpeningTx](#state_swapoutreceiver_broadcastopeningtx)
    * [State_SwapOutReceiver_SendTxBroadcastedMessage](#state_swapoutreceiver_sendtxbroadcastedmessage)
    * [State_SwapOutReceiver_AwaitClaimInvoicePayment](#state_swapoutreceiver_awaitclaiminvoicepayment)
    * [State_SwapOutReceiver_ClaimSwapCoop](#state_swapoutreceiver_claimswapcoop)
    * [State_SwapOutReceiver_ClaimSwapCsv](#state_swapoutreceiver_claimswapcsv)
    * [State_SwapOutReceiver_Aborted](#state_swapoutreceiver_aborted)
  * [Agnostic States](#agnostic-states)
    * [State_SendCancel](#state_sendcancel)
    * [State_SwapCanceled](#state_swapcanceled)
    * [State_ClaimedPreimage](#state_claimedpreimage)
    * [State_ClaimedCoop](#state_claimedcoop)
    * [State_ClaimedCsv](#state_claimedcsv)

## Swap In - Sender
* ### State_SwapInSender_CreateSwap
    #### Strategy
    The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    Peerswap stopped execution before sending a request to the peer. In this state there was no transaction broadcasted to an onchain mempool. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap.

* ### State_SwapInSender_SendRequest
    #### Strategy
    The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    We don't know wether we successfully sent a request to the peer or not. In this state there was no transaction broadcasted to an onchain mempool. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap. The peer might have canceled the swap already.

* ### State_SwapInSender_AwaitAgreement
    #### Strategy
    The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    We do not know if the peer already canceled the swap. In this state there was no transaction broadcasted to an onchain mempool. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap.

* ### State_SwapInSender_BroadcastOpeningTx
    #### Strategy
    Check the chain and the mempool if the commitment transaction is already broadcasted.
    * __Opening transaction not broadcasted:__
  
        The swap shall not be continued and must get dropped. Send cancel and quit swap.
        
    * __Opening transaction in mempool:__
  
        Transition to State_SwapInSender_SendTxBroadcastedMessage.
    * __Opening transaction is confirmed:__
  
        Check how much time delta is left before the refunding path by CSV is possible. If delta is too low, send cancel and transition to State_WaitCSV. If delta is tolerable, continue to State_SwapInSender_SendTxBroadcastedMessage.

    #### Rational
    If the commitment transaction is not broadcasted, the situation is the same as in the states before. If the transaction is already broadcasted or confirmed, we at least want to try to continue the swap as fees for refunding are already due.

* ### State_SwapInSender_SendTxBroadcastedMessage
    #### Strategy
    Check the chain and the mempool if the commitment transaction is confirmed.      
    * __Opening transaction in mempool:__
    
        Resend broadcasted message and continue.
    * __Opening transaction is confirmed:__
  
        Check how much time delta is left before the refunding path by CSV is possible. If delta is too low, send cancel and transition to State_WaitCSV. If delta is tolerable, send transaction broadcasted message and continue.

    #### Rational
    We dont know if we already sent the transaction broadcasted message to the peer. It is in out interest to continue the swap as fees for refunding are already due. We at least want to try to continue the swap.

* ### State_SwapInSender_AwaitClaimPayment
    #### Strategy
    Check if the swap invoice already was payed.
    * __Swap invoice payed:__
  
        Continue to State_ClaimedPreimage.
    * __Swap invoice not payed:__
  
        Wait for payment until CSV passes, then transition to State_SwapInSender_ClaimSwapCsv.
    
    #### Rational
    We dont know if the invoice was payed in the meantime so we need to check. Also we dont know if the peer tried to cancel the swap cooperatively. If the peer tried but did not reach us, the peer might have stopped trying. Still we can just wait until CSV passes as fees are due anyways on refunding. 
    
    [_Maybe we want to reach out to the peer and try to claim coop here if csv delta is low and if invoice is not payed. This way we would not have to wait for a refund until the CSV passes._]

* ### State_SwapInSender_ClaimSwapCoop
    #### Strategy
    Check if transaction output is already claimed.
    * __Claimed:__
  
        Continue to State_ClaimedCoop.
    * __Not claimed:__
  
        Try to claim and proceed accordingly.
    
    #### Rational
    The swap was already canceled cooperatively by the peer. We just dont know if we already claimed the output.

* ### State_SwapInSender_ClaimSwapCsv
    #### Strategy
    Claim swap by csv path.

    #### Rational
    We already want to claim by csv path.

## Swap In - Receiver
* ### State_SwapInReceiver_CreateSwap
    #### Strategy
    The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    Peerswap stopped execution before sendig the agreement message to the peer. In this state there was no transaction broadcasted to an onchain mempool. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap. The peer might have canceled the swap already.

* ### State_SwapInReceiver_SendAgreement
    
    #### Strategy
    The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    We dont know if the agreement was sent and succeded. We dont know if the commitment transaction is already broadcasted, so we can not pay the invoice, or proceed. We can not cancel cooperatively as we do not have the payment request yet.

* ### State_SwapInReceiver_AwaitTxBroadcastedMessage
    
    #### Strategy
    The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    We dont know the commitment transaction id at this point. We can not get any information about the commitment transaction on chain.

    [_We should reach out to our peer and ask for the opening tx again. This way we could continue the swap instead of canceling it. At the current version the protocol does not allow this._]

* ### State_SwapInReceiver_AwaitTxConfirmation
    
    #### Strategy
    Check if the commitment transaction is already confimed.

    * __Confirmed (below CLV delta):__
    
        Transition to State_SwapInReceiver_ValidateTxAndPayClaimInvoice.
    * __Confirmed (above CLV delta):__
    
        Cancel swap cooperatively.
    * __Not confirmed__:

        Wait for confirmation and continue.
    * __Commitment not found__:

        Send cancel and quit swap.

    #### Rational
    We need to know wether the commitment transaction was already confirmed or not. The swap could have been canceled in the meantime. The commitment could already be claimable by the peer or be too close to CSV path to ensure a safe swap.

* ### State_SwapInReceiver_ValidateTxAndPayClaimInvoice
    
    #### Strategy
    Check if the transaction output is already claimed. Check if the invoice was already payed.

    * __Invoice payed:__
    
    Transition to State_SwapInReceiver_ClaimSwap.

    * __Not payed (below CLV delta)__:

    Pay invoice and continue.
    * __Not payed (above CLV delta)__:

    Cancel swap cooperatively.

    #### Rational
    We do not want to pay the invoice and continue the swap if the CLV claiming path is too close. The commitment could already be claimed.

* ### State_SwapInReceiver_ClaimSwap
    
    #### Strategy
    Check if the claim transaction is broadcasted to the mempool or if the claim transaction is already confirmed.

    * __Is in mempool or confirmed:__
    
    Transition to State_ClaimedPreimage.
    * __Is not in mempool__:

    Broadcast claim transaction and continue.
    * __Is claimed by peer (CSV path)__:

    Hate peer and quit swap.

    #### Rational
    We do not know if we already claimed the commitment output. If the peer claimed the commitment output while we were offline, hate him.

* ### State_SwapInReceiver_BuildSigHash
  
    #### Strategy
    Execute state and continue.

    #### Rational
    We are already in a state where we want a cooperative close. We are nice and continue to be cooperative.

* ### State_SwapInReceiver_SendCoopClose
  
    #### Strategy
    Execute state and continue.

    #### Rational
    We dont know if we alrady sent out the coop close message so we just send it again.

## Swap Out - Sender

* ### State_SwapOutSender_CreateSwap
    
     #### Strategy
    The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    Peerswap stopped execution before sending a request to the peer. In this state there was no transaction broadcasted to an onchain mempool. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap.

* ### State_SwapOutSender_SendRequest
    
    #### Strategy
    The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    We don't know wether we successfully sent a request to the peer or not. In this state there was no transaction broadcasted to an onchain mempool. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap. The peer might have canceled the swap already.

* ### State_SwapOutSender_AwaitFeeResponse
    
    #### Strategy
    The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    At this state we did not pay for the fee invoice yet. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap.
    The peer might have canceled the swap already.

* ### State_SwapOutSender_PayFeeInvoice
    
    #### Strategy
    Check if invoice is already payed.

    * __Invoice payed__:
  
        Continue to next state State_SwapOutSender_AwaitTxBroadcastedMessage.
    * __Invoice not payed__:
     
        The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    We want to continue our swap if the fee is already payed. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap. The peer might have canceled the swap already.

* ### State_SwapOutSender_AwaitTxBroadcastedMessage
  
    #### Strategy
    Check if we know the commitment transaction id.

    * __Commitment tx id known__:
        
        Continue to next state State_SwapOutSender_ValidateTxAndPayClaimInvoice. 
    * __Commitment tx id unknown__:

        Continue with cooperative cancel to state State_SwapOutSender_BuildSigHash.
    #### Rational
    We want to continue our swap as we already payed the fees for the transaction. We might not know the commitment transaction id.

    [_The current protocol does not specify asking for a message that was already sent. If we could ask the peer for the commitment tx again we could continue the swap. Also we might send the raw transaction in the agreement message to be investigated and not only after the fee invoice was payed._]

* ### State_SwapOutSender_AwaitTxConfirmation
  
    #### Strategy
    Check the chain and the mempool if the commitment transaction is confirmed. Check for distance to CSV path possibility.

    * __Transaction is not confirmed__:
    
        Continue state.
    * __Transaction is confirmed and below CSV delta__:
        
        Transition to State_SwapOutSender_ValidateTxAndPayClaimInvoice.
    * __Transaction is confirmed and above CSV delta__:

        Try for cooperative close, transition to State_SwapOutSender_BuildSigHash.

    #### Rational
    We dont want to pay the invoice if the claim by CSV path is too close. We Want to continue the swap as we already payed for the fees.

* ### State_SwapOutSender_ClaimSwap

    #### Strategy
    Claim swap.

    #### Rational
    We payed the swap and want to claim the swap.

* ### State_SwapOutSender_BuildSigHash

    #### Strategy
    Continue cooperative cancel.

    #### Rational
    We already want to cancel the swap, try to be cooperative.

* ### State_SwapOutSender_SendCoopClose
  
    #### Strategy
    Continue cooperative cancel.

    #### Rational
    We already want to cancel the swap, try to be cooperative.

## Swap Out - Receiver

* ### State_SwapOutReceiver_CreateSwap
    #### Strategy
    The swap shall not be continued and must get dropped. Send cancel and quit swap.

    #### Rational
    Peerswap stopped execution before sending a fee invoice to the peer. In this state there was no transaction broadcasted to an onchain mempool. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap.
  
* ### State_SwapOutReceiver_SendFeeInvoice
    #### Strategy
    Check if fee invoice is created.
    * __Invoice is not created__:

        The swap shall not be continued and must get dropped. Send cancel and quit swap.
    * __Invoice is created__:
     
        Send invoice and wait transition to State_SwapOutReceiver_AwaitFeeInvoicePayment.

    #### Rational
    We dont know if a fee invoice is already sent to the peer. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap.

* ### State_SwapOutReciever_AwaitFeeInvoicePayment
    #### Strategy
    Check if invoice is payed.

    * __Invoice payed__:
        
        Continue to State_SwapOutReceiver_BroadcastOpeningTx.
    * __Invoice not payed__:
        
        Continue state.

    #### Rational
    The fee invoice is already sent to the peer. We want to be cooperative and continue the swap.
    [_Or dont we want to be coopertive as we will lock our funds in the next steps, and the peer might have canceled in the meantime?_]
* ### State_SwapOutReceiver_BroadcastOpeningTx
    #### Strategy
    Check if the commitment transaction is already broadcasted.

    * __Not broadcasted__:
    
        Broadcast commitment transaction.

    * __Broadcasted__:

        Continue to State_SwapOutReceiver_SendTxBroadcastedMessage. 

    #### Rational
    The fee invoice is already payed. We want to continue the swap.
    [_Or dont we want to be coopertive as we will lock our funds and the peer might have canceled in the meantime?_]

* ### State_SwapOutReciever_SendTxBroadcastedMessage
    #### Strategy
    Send commitment transaction broadcasted message and continue.

    #### Rational
    The commitment transaction is already broadcasted. We want to continue the swap.

* ### State_SwapOutReceiver_AwaitClaimInvoicePayment
    #### Strategy
    Continue state.

    #### Rational
    The commitment transaction is already broadcasted and the peer got the claim invoice. We want to continue the swap.


* ### State_SwapOutReciever_ClaimSwapCoop
    #### Strategy
    Check if transaction output is already claimed.
    * __Claimed:__
  
        Continue to State_ClaimedCoop.
    * __Not claimed:__
  
        Try to claim and proceed accordingly. 

    #### Rational
    The swap was already canceled cooperatively by the peer. We just dont know if we already claimed the output.

* ### State_SwapOutReceiver_ClaimSwapCsv
    #### Strategy
    Claim swap by csv path.

    #### Rational
    We already want to claim by csv path.

* ### State_SwapOutReceiver_Aborted
    #### Strategy
    Transition to State_SwapOutReceiver_ClaimSwapCsv

    #### Rational
    Coop claim failed, we want to claim by CSV path.

## Agnostic States
These states share the same strategy for all swap types and roles.
* ### State_WaitCsv
    #### Strategy
    Check if transaction output is already claimed.
    * __Claimed:__
    
        Continue to State_SwapInSender_ClaimedCsv.
    * __Not claimed:__
  
        Wait for CSV to pass and claim output.
    
    #### Rational
    We dont know if the output was claimed manually.

* ### State_SendCancel
    #### Strategy
    Resend cancel message and continue.
    
    #### Rational
    We dont know if we already sent a cancel message to the peer. It would be nice of us to send it.7

* ### State_SwapCanceled
    Final state.

* ### State_ClaimedPreimage
    Final state.

* ### State_ClaimedCoop
    Final state.

* ### State_ClaimedCsv
    Final state.
