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
    Peerswap stopped execution before sending a request to the peer. In this state there was no transaction broadcasted to an onchain mempool. In the time peerswap was shut down, the channel balance might have changed, we should favor a manual invocation of a new swap. The peer might have canceled the swap already.

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

## Agnostic States
These states share the same strategy for all swap types and roles
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
    We dont know if we already sent a cancel message to the peer. It would be nice of us to send it.

* ### State_SwapCanceled
    Final state.

* ### State_ClaimedPreimage
    Final state.

* ### State_ClaimedCoop
    Final state.

* ### State_ClaimedCsv
    Final state.
