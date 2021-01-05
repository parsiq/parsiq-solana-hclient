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
	transport := &http.Transport{
		MaxIdleConns:        320,
		IdleConnTimeout:     5 * time.Minute,
		MaxIdleConnsPerHost: 32,
	}
	return NewCustomSolanaRpcClient(address, transport)
}

func NewCustomSolanaRpcClient(address string, transport *http.Transport) *SolanaRpcClient {
	client := &http.Client{
		Timeout:   time.Second * 2,
		Transport: transport,
	}
	return &SolanaRpcClient{
		rawClient: client,
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
