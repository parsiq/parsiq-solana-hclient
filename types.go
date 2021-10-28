package parsiq_solana_hclient

type SolanaRpcRequest struct {
	Version string        `json:"jsonrpc"`
	Id      uint64        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type SolanaBaseRpcResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Id int `json:"id"`
}

type RequestAirdropResp struct {
	SolanaBaseRpcResponse
	Result string `json:"result"`
}

type GetConfirmedBlocksResp struct {
	SolanaBaseRpcResponse
	Result []uint64 `json:"result"`
}

type GetSnapshotSlotResp struct {
	SolanaBaseRpcResponse
	Result uint64 `json:"result"`
}

type GetBlockResp struct {
	SolanaBaseRpcResponse
	Result *SolanaBlock `json:"result"`
}

type GetTokenLargestAccountsResp struct {
	SolanaBaseRpcResponse
	Result *TokenLargestAccount `json:"result"`
}

type GetTokenSupplyResp struct {
	SolanaBaseRpcResponse
	Result *TokenAccountBalance `json:"result"`
}

type GetSignatureStatusesResp struct {
	SolanaBaseRpcResponse
	Result *SignatureStatuses `json:"result"`
}

type GetConfirmedTransactionResp struct {
	SolanaBaseRpcResponse
	Result *ConfirmedTransaction `json:"result"`
}

type GetLeaderScheduleResp struct {
	SolanaBaseRpcResponse
	Result interface{} `json:"result"`
}

type GetIdentityResp struct {
	SolanaBaseRpcResponse
	Result struct {
		Identity string `json:"identity"`
	} `json:"result"`
}

type GetConfirmedBlocksWithLimitResp struct {
	SolanaBaseRpcResponse
	Result []uint64 `json:"result"`
}

type GetFirstAvailableBlockResp struct {
	SolanaBaseRpcResponse
	Result uint64 `json:"result"`
}

type GetEpochInfoResp struct {
	SolanaBaseRpcResponse
	Result *EpochInfo `json:"result"`
}

type GetAccountInfoResp struct {
	SolanaBaseRpcResponse
	Result *AccountInfo `json:"result"`
}

type GetEpochScheduleResp struct {
	SolanaBaseRpcResponse
	Result *EpochSchedule `json:"result"`
}

type GetFeeCalculatorForBlockhashResp struct {
	Result interface{} `json:"result"`
}

type MinimumLedgerSlotResp struct {
	SolanaBaseRpcResponse
	Result uint64 `json:"result"`
}

type GetGenesisHashResp struct {
	SolanaBaseRpcResponse
	Result string `json:"result"`
}

type GetBalanceResp struct {
	SolanaBaseRpcResponse
	Result *Balance `json:"result"`
}

type GetVersionResp struct {
	SolanaBaseRpcResponse
	Result *SolanaVersion `json:"result"`
}

type GetRecentPerformanceSamplesResp struct {
	SolanaBaseRpcResponse
	Result []RpcPerfSample `json:"result"`
}

type GetMultipleAccountsResp struct {
	SolanaBaseRpcResponse
	Result *MultipleAccounts `json:"result"`
}

type GetMinimumBalanceForRentExemptionResp struct {
	SolanaBaseRpcResponse
	Result *uint64 `json:"result"`
}

type GetInflationGovernorResp struct {
	SolanaBaseRpcResponse
	Result *InflationGovernor `json:"result"`
}

type GetClusterNodesResp struct {
	SolanaBaseRpcResponse
	Result []*ClusterNodes `json:"result"`
}

type GetInflationRateResp struct {
	SolanaBaseRpcResponse
	Result *InflationRate `json:"result"`
}

type GetLargestAccountsResp struct {
	SolanaBaseRpcResponse
	Result *LargestAccounts `json:"result"`
}

type GetFeesResp struct {
	SolanaBaseRpcResponse
	Result *Fees `json:"result"`
}

type GetMaxRetransmitSlotResp struct {
	SolanaBaseRpcResponse
	Result uint64 `json:"result"`
}

type GetMaxShredInsertSlotResp struct {
	SolanaBaseRpcResponse
	Result uint64 `json:"result"`
}
type GetHealhtResp struct {
	SolanaBaseRpcResponse
	Result interface{} `json:"result"`
}

type SendTransactionResp struct {
	SolanaBaseRpcResponse
	Result string `json:"result"`
}

type SimulateTransactionResp struct {
	SolanaBaseRpcResponse
	Result *SimulateTransaction `json:"result"`
}

type GetTokenAccountBalanceResp struct {
	SolanaBaseRpcResponse
	Result *TokenAccountBalance `json:"result"`
}

type GetTokenAccountsResp struct {
	SolanaBaseRpcResponse
	Result *TokenAccounts `json:"result"`
}

type GetConfirmedSignaturesForAddressResp struct {
	SolanaBaseRpcResponse
	Result *[]ConfirmedSignaturesForAddress `json:"result"`
}

type GetBlockTimeResp struct {
	SolanaBaseRpcResponse
	Result int64 `json:"result"`
}

type GetStakeActivationResp struct {
	SolanaBaseRpcResponse
	Result *StakeActivation `json:"result"`
}

type GetSlotResp struct {
	SolanaBaseRpcResponse
	Result uint64 `json:"result"`
}

type GetSlotLeaderResp struct {
	SolanaBaseRpcResponse
	Result string `json:"result"`
}

type GetVoteAccountsResp struct {
	SolanaBaseRpcResponse
	Result *VoteAccounts `json:"result"`
}

type GetFeeRateGovernorResp struct {
	SolanaBaseRpcResponse
	Result *FeeRateGovernor `json:"result"`
}

type GetSupplyResp struct {
	SolanaBaseRpcResponse
	Result *Supply `json:"result"`
}

type GetProgramAccountsResp struct {
	SolanaBaseRpcResponse
	Result []*ProgramAccounts `json:"result"`
}

type GetRecentBlockHashResp struct {
	SolanaRpcClient
	Result *RecentBlockHash `json:"result"`
}

type GetBlockCommitmentResp struct {
	SolanaBaseRpcResponse
	Result *BlockCommitment `json:"result"`
}

type GetTransactionCountResp struct {
	SolanaBaseRpcResponse
	Result uint64 `json:"result"`
}

type BlockCommitment struct {
	Commitment []uint64 `json:"commitment"`
	TotalStake uint64   `json:"totalStake"`
}

type StakeActivation struct {
	State    string `json:"state"`
	Active   uint64 `json:"active"`
	Inactive uint64 `json:"inactive"`
}

type ConfirmedSignaturesForAddress struct {
	Signature string      `json:"signature"`
	Slot      uint64      `json:"slot"`
	Err       interface{} `json:"err"`
	Memo      string      `json:"memo"`
}

type ConfirmedSignaturesParams struct {
	Limit  uint16 `json:"limit"` //1-1000, default 1000
	Before string `json:"before"`
	Until  string `json:"until"`
}

type Limit struct {
	Limit uint16 `json:"limit"`
}

type Before struct {
	Before string `json:"before"`
}

type Until struct {
	Until string `json:"until"`
}

type SearchTransactionHistory struct {
	SearchTransactionHistory bool `json:"searchTransactionHistory"`
}

type RpcPerfSample struct {
	Slot              uint64 `json:"slot"`
	NumTransactions   uint64 `json:"numTransactions"`
	NumSlots          uint64 `json:"numSlots"`
	SamplePeriodSpecs uint16 `json:"samplePeriodSpecs"`
}

type InflationGovernor struct {
	Initial        float64 `json:"initial"`
	Terminal       float64 `json:"terminal"`
	Taper          float64 `json:"taper"`
	Foundation     float64 `json:"foundation"`
	FoundationTerm float64 `json:"foundationTerm"`
}

type InflationRate struct {
	Total      float64 `json:"total"`
	Validator  float64 `json:"validator"`
	Foundation float64 `json:"foundation"`
	Epoch      float64 `json:"epoch"`
}

//NOT USING DEPRECATED FIELDS
type SignatureStatuses struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value []struct {
		Slot               uint64      `json:"slot"`
		Confirmations      uint        `json:"confirmations"`
		Err                interface{} `json:"err"`
		ConfirmationStatus string      `json:"confirmationStatus"`
	} `json:"value"`
}

type EpochSchedule struct {
	SlotsPerEpoch            uint64 `json:"slotsPerEpoch"`
	LeaderScheduleSlotOffset uint64 `json:"leaderScheduleSlotOffset"`
	Warmup                   bool   `json:"warmup"`
	FirstNormalEpoch         uint64 `json:"firstNormalEpoch"`
	FirstNormalSlot          uint64 `json:"firstNormalSlot"`
}

type SolanaVersion struct {
	SolanaCore string `json:"solana-core"`
	FeatureSet uint64 `json:"feature-set"`
}

type VoteAccounts struct {
	Current    []VoteAccount `json:"current"`
	Delinquent []VoteAccount `json:"delinquent"`
}

type VoteAccount struct {
	VotePubkey       string     `json:"votePubkey"`
	NodePubkey       string     `json:"nodePubkey"`
	ActivatedStake   uint64     `json:"activatedStake"`
	EpochVoteAccount bool       `json:"epochVoteAccount"`
	Commission       uint8      `json:"commission"`
	LastVote         uint64     `json:"lastVote"`
	EpochCredits     [][]uint64 `json:"epochCredits"`
}

type ConfirmedTransaction struct {
	Slot        uint64      `json:"slot"`
	Transaction interface{} `json:"transaction"`
	BlockTime   int64       `json:"blockTime"`
	Meta        struct {
		Err               interface{} `json:"err"`
		Fee               uint64      `json:"fee"`
		PreBalances       []uint64    `json:"preBalances"`
		PostBalances      []uint64    `json:"postBalances"`
		InnerInstructions interface{} `json:"innerInstructions"`
		PreTokenBalances  interface{} `json:"preTokenBalances"`
		PostTokenBalances interface{} `json:"postTokenBalances"`
		LogMessages       []string    `json:"logMessages"`
	} `json:"meta"`
}

type StakeActivationParam struct {
	Commitment string `json:"commitment"`
	Epoch      uint64 `json:"epoch"`
}

type FeeRateGovernor struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value struct {
		FeeRateGovernor struct {
			BurnPercent                uint8  `json:"burnPercent"`
			MaxLamportsPerSignature    uint64 `json:"maxLamportsPerSignature"`
			MinLamportsPerSignature    uint64 `json:"minLamportsPerSignature"`
			TargetLamportsPerSignature uint64 `json:"targetLamportsPerSignature"`
			TargetSignaturesPerSlot    uint64 `json:"targetSignaturesPerSlot"`
		} `json:"feeRateGovernor"`
	}
}

type Supply struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value struct {
		Total                  uint64   `json:"total"`
		Circulating            uint64   `json:"circulating"`
		NonCirculating         uint64   `json:"nonCirculating"`
		NonCirculatingAccounts []string `json:"nonCirculatingAccounts"`
	}
}

type SimulateTransaction struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value struct {
		Err  interface{} `json:"err"`
		Logs []string    `json:"logs"`
	} `json:"value"`
}

type LeadersSchedule struct {
	Slot       uint64
	Commitment string
}

type TokenAccountBalance struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value struct {
		UiAmount float64 `json:"uiAmount"`
		Amount   string  `json:"amount"`
		Decimals uint8   `json:"decimals"`
	} `json:"value"`
}

type TokenLargestAccount struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value []struct {
		Address  string  `json:"address"`
		UiAmount float64 `json:"uiAmount"`
		Amount   string  `json:"amount"`
		Decimals uint8   `json:"decimals"`
	} `json:"value"`
}

type Mint struct {
	Mint string `json:"mint"`
}

type ProgramID struct {
	ProgramID string `json:"programId"`
}

type TokenAccounts struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value []struct {
		PubKey  string `json:"pubKey"`
		Account struct {
			Lamports   uint64      `json:"lamports"`
			Owner      string      `json:"owner"`
			Data       interface{} `json:"data"`
			Executable bool        `json:"executable"`
			RentEpoch  uint64      `json:"rentEpoch"`
		} `json:"account"`
	} `json:"value"`
}

type Fees struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value struct {
		Blockhash     string `json:"blockhash"`
		FeeCalculator struct {
			LamportsPerSignature uint64 `json:"lamportsPerSignature"`
		} `json:"feeCalculator"`
		LastValidSlot uint64 `json:"lastValidSlot"`
	}
}

type Balance struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value uint64 `json:"value"`
}

type ProgramAccounts struct {
	PubKey  string `json:"pubKey"`
	Account struct {
		Lamports   uint64 `json:"lamports"`
		Owner      string `json:"owner"`
		Executable bool   `json:"executable"`
		RentEpoch  uint64 `json:"rentEpoch"`
	} `json:"account"`
}

type RecentBlockHash struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value struct {
		RpcResponse   interface{} `json:"RpcResponse"`
		BlockHash     string      `json:"blockhash"`
		FeeCalculator interface{} `json:"feeCalculator"`
	} `json:"value"`
}

type LargestAccounts struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value []struct {
		Lamports uint64 `json:"lamports"`
		Address  string `json:"address"`
	} `json:"value"`
}

type LargestAccountsParams struct {
	Commitment string `json:"commitment"`
	Filter     string `json:"filter"`
}

type AccountInfoParams struct {
	Commitment string `json:"commitment"`
	Encoding   string `json:"encoding"`
	DataSlice  struct {
		Offset uint `json:"offset"`
		Length uint `json:"length"`
	} `json:"dataSlice"`
}

type ProgramAccountParams struct {
	Commitment string `json:"commitment"`
	Encoding   string `json:"encoding"`
	DataSlice  struct {
		Offset uint `json:"offset"`
		Length uint `json:"length"`
	} `json:"dataSlice"`
	Filters []interface{} `json:"filters"`
}

//We can't pass dataSize within Filter since it MUST have other position within the array in ProgramAccountParams
type DataSize struct {
	DataSize uint64 `json:"dataSize"`
}

type Filter struct {
	Memcmp struct {
		Offset uint64 `json:"offset"`
		Bytes  string `json:"bytes"`
	} `json:"memcmp"`
}

type ClusterNodes struct {
	Gossip  string `json:"gossip"`
	Pubkey  string `json:"pubkey"`
	Rpc     string `json:"rpc"`
	Tpu     string `json:"tpu"`
	Version string `json:"version"`
}

//https://docs.solana.com/developing/clients/jsonrpc-api#configuring-state-commitment
type Commitment struct {
	Commitment string `json:"commitment"`
}

type SimulateTransactionParam struct {
	SigVerify  bool   `json:"sigVerify"`
	Commitment string `json:"commitment"`
	Encoding   string `json:"encoding"`
}

type SendTransactionParams struct {
	SkipPreflight       bool   `json:"skipPreflight"`
	PreflightCommitment string `json:"preflightCommitment"`
	Encoding            string `json:"encoding"`
}

type AccountInfo struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value struct {
		Data       interface{} `json:"data"`
		Executable bool        `json:"executable"`
		Lamports   uint64      `json:"lamports"`
		Owner      string      `json:"owner"`
		RentEpoch  uint64      `json:"rentEpoch"`
	} `json:"value"`
}

type MultipleAccounts struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value []struct {
		Data       interface{} `json:"data"`
		Executable bool        `json:"executable"`
		Lamports   uint64      `json:"lamports"`
		Owner      string      `json:"owner"`
		RentEpoch  uint64      `json:"rentEpoch"`
	} `json:"value"`
}

type EpochInfo struct {
	AbsoluteSlot uint64 `json:"absoluteSlot"`
	BlockHeight  uint64 `json:"blockHeight"`
	Epoch        uint64 `json:"epoch"`
	SlotIndex    uint64 `json:"slotIndex"`
	SlotsInEpoch uint64 `json:"slotsInEpoch"`
}

type SolanaBlock struct {
	SlotNumber        uint64 `json:"slotNumber"`
	BlockTime         uint64 `json:"blockTime"`
	Blockhash         string `json:"blockhash"`
	ParentSlot        uint64 `json:"parentSlot"`
	PreviousBlockhash string `json:"previousBlockhash"`
	Rewards           []struct {
		Lamports    int64  `json:"lamports"`
		PostBalance uint64 `json:"postBalance"`
		Pubkey      string `json:"pubkey"`
		RewardType  string `json:"rewardType"`
	} `json:"rewards"`
	Transactions []*SolanaTransaction `json:"transactions"`
}

type SolanaTransaction struct {
	Meta struct {
		Err               interface{}   `json:"err"`
		Fee               uint64        `json:"fee"`
		InnerInstructions []interface{} `json:"innerInstructions"`
		LogMessages       []string      `json:"logMessages"`
		PreBalances       []uint64      `json:"preBalances"`
		PostBalances      []uint64      `json:"postBalances"`
		PreTokenBalances  []*struct {
			AccountIndex  uint16 `json:"accountIndex"`
			Mint          string `json:"mint"`
			UiTokenAmount struct {
				Amount   string  `json:"amount"`
				Decimals uint8   `json:"decimals"`
				UiAmount float64 `json:"uiAmount"`
			} `json:"uiTokenAmount"`
		} `json:"preTokenBalances"`
		PostTokenBalances []*struct {
			AccountIndex  uint16 `json:"accountIndex"`
			Mint          string `json:"mint"`
			UiTokenAmount struct {
				Amount   string  `json:"amount"`
				Decimals uint8   `json:"decimals"`
				UiAmount float64 `json:"uiAmount"`
			} `json:"uiTokenAmount"`
		} `json:"postTokenBalances"`
		Status struct {
			Ok interface{} `json:"Ok"`
		} `json:"status"`
	} `json:"meta"`
	Transaction struct {
		Message struct {
			AccountKeys []string `json:"accountKeys"`
			Header      struct {
				NumReadonlySignedAccounts   uint16 `json:"numReadonlySignedAccounts"`
				NumReadonlyUnsignedAccounts uint16 `json:"numReadonlyUnsignedAccounts"`
				NumRequiredSignatures       uint16 `json:"numRequiredSignatures"`
			} `json:"header"`
			Instructions []struct {
				Accounts       []uint16 `json:"accounts"`
				Data           string   `json:"data"`
				ProgramIDIndex uint16   `json:"programIdIndex"`
			} `json:"instructions"`
			RecentBlockhash string `json:"recentBlockhash"`
		} `json:"message"`
		Signatures []string `json:"signatures"`
	} `json:"transaction"`
}
