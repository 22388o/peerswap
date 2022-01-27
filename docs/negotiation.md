# Negotiation

```mermaid
sequenceDiagram
participant A as Alice
participant B as Bob

Note over A: Alice wants to swap
A->>B: swap request (direction, amt, pubkey, chain params)

Note over B: Check policy
alt Alice wants to swap in
B->>A: Send info (pubkey, policy)
Note over A: Create preimage
A->>B: Send Swap Invoice(amt - policy von bob)
else Alice wants to swap out
Note over B: Create preimage
B->>A: Send Swap Invoice(amt + policy von bob)
else Bob does not agree
B->>A: Send cancel
end
```

Beide State:

Swap Invoice (swap amount, pubkeyA, pubkeyB, swap phash)

Peer Lightning pubkey

Taker:
Swap Preimage

Swap listener
- Phashs von contracts listenen und dann opening tx broadcast
# Swap
```mermaid
sequenceDiagram
participant A as Maker
participant B as Taker
A->>A: listen for swap phash payments
Note over A: Check swap payment
Note over A: broadcast opening tx
A->>A: check for claim tx (starting block height)
Note over B: on N confs, broadcast claim tx
alt Taker broadcasted claim tx, quick resolve
B->>A: Send Preimage message
Note over A: Settle lightning payment with preimage from message
else Taker broadcasted claim tx
Note over A: Settle lightning payment with preimage from claim tx
else Bob never claimed, preimage never known to alice, csv passed
Note over A: cancel hodl payment
Note over A: broadcast claim tx (pk alice)
end
```