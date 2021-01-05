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
