package parsiq_solana_hclient

import (
	"fmt"
	"testing"
)

const (
	testApiRpcAddr = "http://10.20.30.100:8899"
)

func TestGetConfirmedBlock(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetConfirmedBlock(59_211_174)
	if err != nil {
		panic(err)
	}
	fmt.Printf("+%v\n", resp.Result)
}

func TestGetConfirmedBlockWithLimit(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetConfirmedBlocksWithLimit(59_212_174, 100)
	if err != nil {
		panic(err)
	}
	fmt.Printf("+%v\n", resp.Result)
}

func TestGetEpochInfo(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetEpochInfo()
	if err != nil {
		panic(err)
	}
	fmt.Printf("+%v\n", resp.Result)
}

func TestGetAccountInfo(t *testing.T) {
	type Parameters struct {
		Encoding  string `json:"encoding"`
		DataSlice struct {
			Offset uint `json:"offset"`
			Length uint `json:"length"`
		} `json:"dataSlice"`
	}
	params := Parameters{
		Encoding: "base58",
	}
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetAccountInfo("A39v6FsgGKMQcEQjHnb88xPxpGKxx1FkhA6CkDzL8b3e", params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("+%v\n", resp.Result)
}

func TestGetFirstAvailableBlock(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetFirstAvailableBlock()
	if err != nil {
		panic(err)
	}
	fmt.Printf("+%v\n", resp.Result)
}

func TestGetGenesisHash(t *testing.T) {
	client := NewSolanaRpcClient(testApiRpcAddr)
	resp, err := client.GetGenesisHash()
	if err != nil {
		panic(err)
	}
	fmt.Printf("+%v\n", resp.Result)
}
