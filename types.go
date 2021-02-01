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

type GetConfirmedBlockResp struct {
	SolanaBaseRpcResponse
	Result *SolanaBlock `json:"result"`
}

type GetTokenLargestAccountsResp struct {
	SolanaBaseRpcResponse
	Result *TokenLargestAccount `json:"result"`
}

type GetTokenSupply struct {
	SolanaBaseRpcResponse
	Result *TokenAccountBalance `json:"result"`
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

type GetGenesisHashResp struct {
	SolanaBaseRpcResponse
	Result string `json:"result"`
}

type GetBalanceResp struct {
	SolanaBaseRpcResponse
	Result *Balance `json:"result"`
}

type GetClusterNodesResp struct {
	SolanaBaseRpcResponse
	Result []*ClusterNodes `json:"result"`
}

type GetLargestAccountsResp struct {
	SolanaBaseRpcResponse
	Result *LargestAccounts `json:"result"`
}

type GetFeesResp struct {
	SolanaBaseRpcResponse
	Result *Fees `json:"result"`
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

type BlockCommitment struct {
	Commitment []uint64 `json:"commitment"`
	TotalStake uint64   `json:"totalStake"`
}

type StakeActivation struct {
	State    string `json:"state"`
	Active   uint64 `json:"active"`
	Inactive uint64 `json:"inactive"`
}

type StakeActivationParam struct {
	Commitment string `json:"commitment"`
	Epoch      uint64 `json:"epoch"`
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
	Filters []Filter `json:"filters"`
}

type Filter struct {
	Memcmp struct {
		Offset uint64 `json:"offset"`
		Bytes  string `json:"bytes"`
	} `json:"memcmp"`
	DataSize uint64 `json:"dataSize"`
}

type ClusterNodes struct {
	Gossip  string `json:"gossip"`
	Pubkey  string `json:"pubkey"`
	Rpc     string `json:"rpc"`
	Tpu     string `json:"tpu"`
	Version string `json:"version"`
}

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

type EpochInfo struct {
	AbsoluteSlot uint64 `json:"absoluteSlot"`
	BlockHeight  uint64 `json:"blockHeight"`
	Epoch        uint64 `json:"epoch"`
	SlotIndex    uint64 `json:"slotIndex"`
	SlotsInEpoch uint64 `json:"slotsInEpoch"`
}

type SolanaBlock struct {
	BlockTime         uint64 `json:"blockTime"`
	Blockhash         string `json:"blockhash"`
	ParentSlot        uint64 `json:"parentSlot"`
	PreviousBlockhash string `json:"previousBlockhash"`
	Rewards           []struct {
		Lamports    uint64 `json:"lamports"`
		PostBalance uint64 `json:"postBalance"`
		Pubkey      string `json:"pubkey"`
		RewardType  string `json:"rewardType"`
	} `json:"rewards"`
	Transactions []struct {
		Meta struct {
			Err               interface{}   `json:"err"`
			Fee               uint64        `json:"fee"`
			InnerInstructions []interface{} `json:"innerInstructions"`
			LogMessages       []string      `json:"logMessages"`
			PostBalances      []uint64      `json:"postBalances"`
			PreBalances       []uint64      `json:"preBalances"`
			Status            struct {
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
	} `json:"transactions"`
}
