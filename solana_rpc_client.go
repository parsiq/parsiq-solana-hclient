package parsiq_solana_hclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type SolanaRpcClient struct {
	rawClient *http.Client
	host      string
}

func NewSolanaRpcClient(address string) *SolanaRpcClient {
	client := &http.Client{
		Timeout: time.Second * 3,
		Transport: &http.Transport{
			IdleConnTimeout:       5 * time.Minute,
			MaxIdleConnsPerHost:   100,
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
	return NewCustomSolanaRpcClient(address, client)
}
func NewCustomSolanaRpcClient(address string, httpClient *http.Client) *SolanaRpcClient {
	return &SolanaRpcClient{
		rawClient: httpClient,
		host:      address,
	}
}

// https://docs.solana.com/developing/clients/jsonrpc-api#getconfirmedblock
func (client *SolanaRpcClient) GetConfirmedBlock(slotNumber uint64) (*GetConfirmedBlockResp, error) {
	request := client.buildRequest("getConfirmedBlock", slotNumber, "json")
	responseObj := &GetConfirmedBlockResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	responseObj.Result.SlotNumber = slotNumber
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#getconfirmedblockswithlimit
func (client *SolanaRpcClient) GetConfirmedBlocksWithLimit(startSlot, limit uint64) (*GetConfirmedBlocksWithLimitResp, error) {
	request := client.buildRequest("getConfirmedBlocksWithLimit", startSlot, limit)
	responseObj := &GetConfirmedBlocksWithLimitResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#getepochinfo
func (client *SolanaRpcClient) GetEpochInfo(commitment ...*Commitment) (*GetEpochInfoResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getEpochInfo")
	} else {
		request = client.buildRequest("getEpochInfo", commitment[0])
	}
	responseObj := &GetEpochInfoResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#getaccountinfo
//TODO doesnt return full data if optional parameters are passed, possibly bug on solana side
func (client *SolanaRpcClient) GetAccountInfo(pubKey string, params ...*AccountInfoParams) (*GetAccountInfoResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("getAccountInfo", pubKey)
	} else {
		request = client.buildRequest("getAccountInfo", pubKey, params[0])
	}
	responseObj := &GetAccountInfoResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#getfirstavailableblock
func (client *SolanaRpcClient) GetFirstAvailableBlock() (*GetFirstAvailableBlockResp, error) {
	request := client.buildRequest("getFirstAvailableBlock")
	responseObj := &GetFirstAvailableBlockResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#getgenesishash
func (client *SolanaRpcClient) GetGenesisHash() (*GetGenesisHashResp, error) {
	request := client.buildRequest("getGenesisHash")
	responseObj := &GetGenesisHashResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#getbalance
func (client *SolanaRpcClient) GetBalance(pubKey string, commitment ...*Commitment) (*GetBalanceResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getBalance", pubKey)
	} else {
		request = client.buildRequest("getBalance", pubKey, commitment[0])
	}
	responseObj := &GetBalanceResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getblockcommitment
//TODO make a proper test for this one
func (client *SolanaRpcClient) GetBlockCommitment(block uint64) (*GetBlockCommitmentResp, error) {
	request := client.buildRequest("getBlockCommitment", block)
	responseObj := &GetBlockCommitmentResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#getclusternodes
func (client *SolanaRpcClient) GetClusterNodes() (*GetClusterNodesResp, error) {
	request := client.buildRequest("getClusterNodes")
	responseObj := &GetClusterNodesResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#getlargestaccounts requires more time to process, around 15 seconds
func (client *SolanaRpcClient) GetLargestAccounts(params ...*LargestAccountsParams) (*GetLargestAccountsResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("getLargestAccounts")
	} else {
		request = client.buildRequest("getLargestAccounts", params[0])
	}
	responseObj := &GetLargestAccountsResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#getfees
func (client *SolanaRpcClient) GetFees(commitment ...*Commitment) (*GetFeesResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getFees")
	} else {
		request = client.buildRequest("getFees", commitment[0])
	}
	responseObj := &GetFeesResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#simulatetransaction
func (client *SolanaRpcClient) SimulateTransaction(blockHash string, params ...*SimulateTransactionParam) (*SimulateTransactionResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("simulateTransaction", blockHash)
	} else {
		request = client.buildRequest("simulateTransaction", blockHash, params[0])
	}
	responseObj := &SimulateTransactionResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#sendtransaction
func (client *SolanaRpcClient) SendTransaction(transaction string, params ...*SendTransactionParams) (*SendTransactionResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("simulateTransaction", transaction)
	} else {
		request = client.buildRequest("simulateTransaction", transaction, params[0])
	}
	responseObj := &SendTransactionResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#gettokenaccountbalance UNSTABLE USE AT YOUR OWN RISK
func (client *SolanaRpcClient) GetTokenAccountBalance(pubKey string, commitment ...*Commitment) (*GetTokenAccountBalanceResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getTokenAccountBalance", pubKey)
	} else {
		request = client.buildRequest("getTokenAccountBalance", pubKey, commitment[0])
	}
	responseObj := &GetTokenAccountBalanceResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#gettokenaccountsbydelegate mint is used instead of programId. UNSTABLE USE AT YOUR OWN RISK
func (client *SolanaRpcClient) GetTokenAccountByDelegate(pubKey, mint string, params ...*AccountInfoParams) (*GetTokenAccountsResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("getTokenAccountsByDelegate", pubKey, &Mint{Mint: mint})
	} else {
		request = client.buildRequest("getTokenAccountsByDelegate", pubKey, &Mint{Mint: mint}, params[0])
	}
	responseObj := &GetTokenAccountsResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#gettokenaccountsbydelegate programId instead of mint. UNSTABLE USE AT YOUR OWN RISK
func (client *SolanaRpcClient) GetTokenAccountByDelegateByProgramID(pubKey, programId string, params ...*AccountInfoParams) (*GetTokenAccountsResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("getTokenAccountsByDelegate", pubKey, &ProgramID{ProgramID: programId})
	} else {
		request = client.buildRequest("getTokenAccountsByDelegate", pubKey, &ProgramID{ProgramID: programId}, params[0])
	}
	responseObj := &GetTokenAccountsResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#gettokenaccountsbyowner mint is used instead of programId. UNSTABLE USE AT YOUR OWN RISK
func (client *SolanaRpcClient) GetTokenAccountByOwner(pubKey, mint string, params ...*AccountInfoParams) (*GetTokenAccountsResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("getTokenAccountsByOwner", pubKey, &Mint{Mint: mint})
	} else {
		request = client.buildRequest("getTokenAccountsByOwner", pubKey, &Mint{Mint: mint}, params[0])
	}
	responseObj := &GetTokenAccountsResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#gettokenaccountsbyowner programId is used instead of mint. UNSTABLE USE AT YOUR OWN RISK
func (client *SolanaRpcClient) GetTokenAccountByOwnerByProgramID(pubKey, programId string, params ...*AccountInfoParams) (*GetTokenAccountsResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("getTokenAccountsByOwner", pubKey, &ProgramID{ProgramID: programId})
	} else {
		request = client.buildRequest("getTokenAccountsByOwner", pubKey, &ProgramID{ProgramID: programId}, params[0])
	}
	responseObj := &GetTokenAccountsResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#gettokenlargestaccounts UNSTABLE USE AT YOUR OWN RISK
func (client *SolanaRpcClient) GetTokenLargestAccounts(pubKey string, commitment ...*Commitment) (*GetTokenLargestAccountsResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getTokenLargestAccounts", pubKey)
	} else {
		request = client.buildRequest("getTokenLargestAccounts", pubKey, commitment[0])
	}
	responseObj := &GetTokenLargestAccountsResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

// https://docs.solana.com/developing/clients/jsonrpc-api#gettokensupply UNSTABLE USE AT YOUR OWN RISK
func (client *SolanaRpcClient) GetTokenSupply(pubKey string, commitment ...*Commitment) (*GetTokenSupplyResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getTokenSupply", pubKey)
	} else {
		request = client.buildRequest("getTokenSupply", pubKey, commitment[0])
	}
	responseObj := &GetTokenSupplyResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getslot
func (client *SolanaRpcClient) GetSlot(commitment ...*Commitment) (*GetSlotResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getSlot")
	} else {
		request = client.buildRequest("getSlot", commitment[0])
	}
	responseObj := &GetSlotResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getblocktime
func (client *SolanaRpcClient) GetBlockTime(block uint64) (*GetBlockTimeResp, error) {
	request := client.buildRequest("getBlockTime", block)
	responseObj := &GetBlockTimeResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getrecentblockhash
func (client *SolanaRpcClient) GetRecentBlockhash(commitment ...*Commitment) (*GetRecentBlockHashResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getRecentBlockhash")
	} else {
		request = client.buildRequest("getRecentBlockhash", commitment[0])
	}
	responseObj := &GetRecentBlockHashResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getidentity
func (client *SolanaRpcClient) GetIdentity() (*GetIdentityResp, error) {
	request := client.buildRequest("getIdentity")
	responseObj := &GetIdentityResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getinflationgovernor
func (client *SolanaRpcClient) GetInflationGovernor(commitment ...*Commitment) (*GetInflationGovernorResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getInflationGovernor")
	} else {
		request = client.buildRequest("getInflationGovernor", commitment[0])
	}
	responseObj := &GetInflationGovernorResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getinflationrate
func (client *SolanaRpcClient) GetInflationRate() (*GetInflationRateResp, error) {
	request := client.buildRequest("getInflationRate")
	responseObj := &GetInflationRateResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getleaderschedule
func (client *SolanaRpcClient) GetLeadersSchedule(params ...*LeadersSchedule) (*GetLeaderScheduleResp, error) {
	request := &SolanaRpcRequest{}
	if params != nil {
		if params[0].Commitment == "" {
			if params[0].Slot != 0 {
				request = client.buildRequest("getLeaderSchedule", params[0].Slot)
			} else {
				request = client.buildRequest("getLeaderSchedule")
			}
		} else {
			if params[0].Slot != 0 {
				request = client.buildRequest("getLeaderSchedule", params[0].Slot, Commitment{Commitment: params[0].Commitment})
			} else {
				request = client.buildRequest("getLeaderSchedule", 0, Commitment{Commitment: params[0].Commitment})
			}
		}
	} else {
		request = client.buildRequest("getLeaderSchedule")
	}
	responseObj := &GetLeaderScheduleResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getfeecalculatorforblockhash
func (client *SolanaRpcClient) GetFeeCalculatorForBlockhash(blockhash string, commitment ...*Commitment) (*GetFeeCalculatorForBlockhashResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getFeeCalculatorForBlockhash", blockhash)
	} else {
		request = client.buildRequest("getFeeCalculatorForBlockhash", blockhash, commitment[0])
	}
	responseObj := &GetFeeCalculatorForBlockhashResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getfeerategovernor
func (client *SolanaRpcClient) GetFeeRateGovernor() (*GetFeeRateGovernorResp, error) {
	request := client.buildRequest("getFeeRateGovernor")
	responseObj := &GetFeeRateGovernorResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getmultipleaccounts
func (client *SolanaRpcClient) GetMultipleAccounts(pubkeys []string, params ...*AccountInfoParams) (*GetMultipleAccountsResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("getMultipleAccounts", pubkeys)
	} else {
		request = client.buildRequest("getMultipleAccounts", pubkeys, params[0])
	}
	responseObj := &GetMultipleAccountsResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getprogramaccounts
//TODO "code":-32602,"message":"Invalid params: invalid value: map, expected map with a single key."
func (client *SolanaRpcClient) GetProgramAccounts(pubKey string, params ...*ProgramAccountParams) (*GetProgramAccountsResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("getProgramAccounts", pubKey)
	} else {
		request = client.buildRequest("getProgramAccounts", pubKey, params[0])
	}
	responseObj := &GetProgramAccountsResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getslotleader
func (client *SolanaRpcClient) GetSlotLeader(commitment ...*Commitment) (*GetSlotLeaderResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getSlotLeader")
	} else {
		request = client.buildRequest("getSlotLeader", commitment[0])
	}
	responseObj := &GetSlotLeaderResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getstakeactivation
func (client *SolanaRpcClient) GetStakeActivation(pubKey string, param ...*StakeActivationParam) (*GetStakeActivationResp, error) {
	request := &SolanaRpcRequest{}
	if param == nil {
		request = client.buildRequest("getStakeActivation", pubKey)
	} else {
		request = client.buildRequest("getStakeActivation", pubKey, param[0])
	}
	responseObj := &GetStakeActivationResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getsupply
func (client *SolanaRpcClient) GetSupply(commitment ...*Commitment) (*GetSupplyResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getSupply")
	} else {
		request = client.buildRequest("getSupply", commitment[0])
	}
	responseObj := &GetSupplyResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getconfirmedsignaturesforaddress2
func (client *SolanaRpcClient) GetConfirmedSignaturesForAddress2(address string, params ...*ConfirmedSignaturesParams) (*GetConfirmedSignaturesForAddressResp, error) {
	request := &SolanaRpcRequest{}
	if params == nil {
		request = client.buildRequest("getConfirmedSignaturesForAddress2", address)
	} else {
		if params[0].Limit <= 0 {
			//1000 is default parameter for limit
			params[0].Limit = 1000
		}
		if params[0].Before != "" && params[0].Until != "" && params[0].Limit > 0 {
			request = client.buildRequest("getConfirmedSignaturesForAddress2", address,
				Limit{Limit: params[0].Limit},
				Before{Before: params[0].Before},
				Until{Until: params[0].Until})
		}

		if params[0].Until != "" {
			request = client.buildRequest("getConfirmedSignaturesForAddress2", address,
				Limit{Limit: params[0].Limit},
				Until{Until: params[0].Until})
		}
		if params[0].Before != "" {
			request = client.buildRequest("getConfirmedSignaturesForAddress2", address,
				Limit{Limit: params[0].Limit},
				Before{Before: params[0].Before})
		}
		request = client.buildRequest("getConfirmedSignaturesForAddress2", address, Limit{Limit: params[0].Limit})
	}
	responseObj := &GetConfirmedSignaturesForAddressResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

//https://docs.solana.com/developing/clients/jsonrpc-api#getminimumbalanceforrentexemption
func (client *SolanaRpcClient) GetMinimumBalanceForRentExemption(accountDataSize uint, commitment ...*Commitment) (*GetMinimumBalanceForRentExemptionResp, error) {
	request := &SolanaRpcRequest{}
	if commitment == nil {
		request = client.buildRequest("getMinimumBalanceForRentExemption", accountDataSize)
	} else {
		request = client.buildRequest("getMinimumBalanceForRentExemption", accountDataSize, commitment[0])
	}
	responseObj := &GetMinimumBalanceForRentExemptionResp{}
	if err := client.doRequest(request, responseObj); err != nil {
		return nil, err
	}
	return responseObj, nil
}

func (client *SolanaRpcClient) doRequest(request *SolanaRpcRequest, responseObj interface{}) error {
	buffer := &bytes.Buffer{}
	data, _ := json.Marshal(request)
	buffer.Write(data)

	response, err := client.rawClient.Post(client.host, "application/json", buffer)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(bodyBytes))
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, responseObj)
}

func (request *SolanaRpcClient) buildRequest(method string, paramsIn ...interface{}) *SolanaRpcRequest {
	return &SolanaRpcRequest{
		Version: "2.0",
		Id:      1,
		Method:  method,
		Params:  paramsIn,
	}
}
