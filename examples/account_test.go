package examples

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/types"
	"testing"
)

func TestGetAccountBalance(t *testing.T) {
	oklinkClient, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	accountItem := []string{"0xD79053a14BC465d9C1434d4A4fAbdeA7b6a2A94b"}
	symbol := []string{"ETH"}
	contractAddress := []string{"0x00"}
	protocolType := []string{""}
	page := []string{"1"}
	limit := []string{"10"}
	acbr := &types.AccountBalanceRequest{
		ChainShortName:  "ETH",
		ExplorerName:    "etherescan",
		Account:         accountItem,
		Symbol:          symbol,
		ContractAddress: contractAddress,
		ProtocolType:    protocolType,
		Page:            page,
		Limit:           limit,
	}
	etherscanResp, err := etherscanClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========etherscanResp============")
	fmt.Println(etherscanResp.BalanceStr)
	fmt.Println(etherscanResp.Balance.Int())
	fmt.Println(etherscanResp.Account)
	fmt.Println(etherscanResp.Symbol)
	fmt.Println("===========etherscanResp===========")

	oklinkResp, err := oklinkClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========oklinkResp============")
	fmt.Println(oklinkResp.BalanceStr)
	fmt.Println(oklinkResp.Account)
	fmt.Println(oklinkResp.Symbol)
	fmt.Println("==========oklinkResp============")
}
