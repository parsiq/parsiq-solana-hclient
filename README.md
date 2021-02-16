# Solana Blockchain Golang JSON RPC client

Client initialization example:

```golang
client := NewSolanaRpcClient(testApiRpcAddr)
resp, err := client.GetConfirmedBlock(59_211_174)
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", resp.Result)
```

Supported API methods:

* getConfirmedBlock
* getConfirmedBlocksWithLimit
* getEpochInfo
* getFirstAvailableBlock
* getGenesisHash
* getAccountInfo
* getBalance
* getClusterNodes
* getLargestAccounts
* getFees
* getFeeCalculatorForBlockhash
* getFeeRateGovernor
* simulateTransaction
* sendTransaction
* getTokenAccountBalance
* getTokenAccountsByDelegate
* getTokenAccountsByOwner
* getTokenLargestAccounts
* getTokenSupply
* getSlot
* getSlotLeader
* getStakeActivation
* getSupply
* getBlockCommitment
* getBlockTime
* getRecentBlockhash
* getIdentity
* getInflationGovernor
* getInflationRate
* getLeadersSchedule
* getMultipleAccounts