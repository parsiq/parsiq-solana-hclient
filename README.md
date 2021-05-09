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

* getAccountInfo
* getBalance
* getBlockCommitment
* getBlockTime
* getClusterNodes
* getConfirmedBlock
* getConfirmedBlocks
* getConfirmedBlocksWithLimit
* getConfirmedSignaturesForAddress2
* getConfirmedTransaction
* getEpochInfo
* getEpochSchedule
* getFees
* getFeeCalculatorForBlockhash
* getFeeRateGovernor
* getFirstAvailableBlock
* getGenesisHash
* getHealth
* getIdentity
* getInflationGovernor
* getInflationRate
* getLargestAccounts
* getLeadersSchedule
* getMaxRetransmitSlot
* getMaxShredInsertSlot
* getMinimumBalanceForRentExemption
* getMultipleAccounts
* getProgramAccounts
* getRecentBlockhash
* getRecentPerformanceSamples
* getSignatureStatuses
* getSlot
* getSlotLeader
* getStakeActivation
* getSupply
* getTokenAccountBalance
* getTokenAccountsByDelegate
* getTokenAccountsByOwner
* getTokenLargestAccounts
* getTokenSupply
* getTransactionCount
* getVersion
* getVoteAccounts
* minimumLedgerSlot
* requestAirdrop
* sendTransaction
* simulateTransaction
* getSnapshotSlot
