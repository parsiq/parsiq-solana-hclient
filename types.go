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

type GetConfirmedBlocksWithLimitResp struct {
	SolanaBaseRpcResponse
	Result []uint64 `json:"result"`
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
