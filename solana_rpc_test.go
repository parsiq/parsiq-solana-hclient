package parsiq_solana_hclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const (
	testApiRpcAddr = "http://10.20.30.100:8899"
)

func TestGetConfirmedBlock(t *testing.T) {

	client := NewSolanaRpcClient(testApiRpcAddr)
	epoch, _ := client.GetEpochInfo()
	resp, err := client.GetConfirmedBlock(epoch.Result.AbsoluteSlot)
	if err != nil {
		panic(err)
	}
	assert.NotEqual(t, resp.Result.BlockTime, nil, "block time is nil")
	assert.NotEqual(t, resp.Result.Blockhash, nil, "blockhash is nil")
	assert.NotEqual(t, resp.Result.Rewards, nil, "block time is nil")
	fmt.Printf("%+v\n", resp.Result)
}

func TestGetConfirmedBlockWithLimit(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	epoch, _ := client.GetEpochInfo()
	resp, err := client.GetConfirmedBlocksWithLimit(epoch.Result.AbsoluteSlot, 100)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}

func TestGetEpochInfo(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetEpochInfo()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
	assert.NotEqual(t, resp.Result.Epoch, nil, "epoch is nil")
	assert.NotEqual(t, resp.Result.AbsoluteSlot, nil, "absolute slot is nil")
	assert.NotEqual(t, resp.Result.SlotsInEpoch, nil, "slots in epoch is nil")
}

func TestGetAccountInfo(t *testing.T) {
	params := &AccountInfoParams{
		Encoding: "base64",
	}
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetAccountInfo("4fYNw3dojWmQ4dXtSGE9epjRGy9pFSx62YypT7avPYvA", params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}

func TestGetFirstAvailableBlock(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetFirstAvailableBlock()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}

func TestGetGenesisHash(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetGenesisHash()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}

func TestGetBalance(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	commitment := &Commitment{}
	commitment.Commitment = "max"
	resp, err := client.GetBalance("GK2zqSsXLA2rwVZk347RYhh6jJpRsCA69FjLW93ZGi3B", commitment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}

func TestGetClusterNodes(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetClusterNodes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result[0])
}

func TestGetLargestAccounts(t *testing.T) {
	params := &LargestAccountsParams{
		Commitment: "max",
		Filter:     "nonCirculating",
	}
	client := &http.Client{
		Transport: &http.Transport{
			IdleConnTimeout:       5 * time.Minute,
			MaxIdleConnsPerHost:   100,
			ResponseHeaderTimeout: time.Second * 30,
		},
	}
	custom := NewCustomSolanaRpcClient(testApiRpcAddr, client)
	resp, err := custom.GetLargestAccounts(params)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v\n", resp.Result)
	params = &LargestAccountsParams{
		Commitment: "max",
		Filter:     "circulating",
	}

	resp2, err := custom.GetLargestAccounts(params)
	if err != nil {
		panic(err)
	}
	assert.NotEqual(t, resp, resp2, "objects are identical")
}

func TestGetFees(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetFees(&Commitment{Commitment: "max"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}

func TestSimulateTransaction(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.SimulateTransaction("4hXTCkRzt9WyecNzV1XPgCDfGAZzQKNxLXgynz5QDuWWPSAZBZSHptvWRL3BjCvzUXRdKvHL2b7yGrRQcWyaqsaBCncVG7BFggS8w9snUts67BSh3EqKpXLUm5UMHfD7ZBe9GhARjbNQMLJ1QD3Spr6oMTBU6EhdB4RD8CP2xUxr2u3d6fos36PD98XS6oX8TQjLpsMwncs5DAMiD4nNnR8NBfyghGCWvCVifVwvA8B8TJxE1aiyiv2L429BCWfyzAme5sZW8rDb14NeCQHhZbtNqfXhcp2tAnaAT")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}

//TODO for testing required v2 token
func TestGetTokenAccountBalance(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetTokenAccountBalance("DY3613oY9RcPhs4jeZPevpueZGZcWRXNysicT8i3DwFZ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}

func TestGetTokenAccountByDelegate(t *testing.T) {
	client := &http.Client{
		Transport: &http.Transport{
			IdleConnTimeout:       5 * time.Minute,
			MaxIdleConnsPerHost:   100,
			ResponseHeaderTimeout: time.Second * 320,
		},
	}
	custom := NewCustomSolanaRpcClient(testApiRpcAddr, client)
	resp, err := custom.GetTokenAccountByDelegate("63hfbwj4LMkL45t1hhbVp4ajBsdwMTp1Jg6kjGAJq1SU", "B79Rux3VRvZWgTqbgFp8vq4ezyMzrApvvfDLuNmuLeen")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}

func TestGetTokenLargestAccounts(t *testing.T) {
	client := &http.Client{
		Transport: &http.Transport{
			IdleConnTimeout:       5 * time.Minute,
			MaxIdleConnsPerHost:   100,
			ResponseHeaderTimeout: time.Second * 60,
		},
	}
	custom := NewCustomSolanaRpcClient(testApiRpcAddr, client)
	resp, err := custom.GetTokenLargestAccounts("5jqymuoXXVcUuJKrf1MWiHSqHyg2osMaJGVy69NsJWyP")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}

func TestGetTokenSupply(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetTokenSupply("5jqymuoXXVcUuJKrf1MWiHSqHyg2osMaJGVy69NsJWyP")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result)
}
