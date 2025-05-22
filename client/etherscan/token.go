package etherscan

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/pkg/helper"
	"github.com/web3-fighter/chain-explorer-api/types"
)

func (c *ChainExplorerClient) GetTokenList(req *types.TokenRequest) ([]*types.TokenResponse, error) {
	var responseData []TokensResp
	param := helper.M{
		"contractaddress": req.ContractAddress,
	}
	err := c.baseClient.Call(ChainExplorerName, "token", "tokeninfo", "", param, &responseData)
	if err != nil {
		fmt.Println("call token list for etherscan fail", "err", err)
		return nil, err
	}
	var tokenList []*types.TokenResponse
	for _, tokenValue := range responseData {
		tokenItem := &types.TokenResponse{
			Symbol:               tokenValue.Symbol,
			TokenContractAddress: tokenValue.ContractAddress,
			TokenId:              tokenValue.TokenId,
			TotalSupply:          tokenValue.TotalSupply,
			Decimal:              tokenValue.Divisor,
		}
		tokenList = append(tokenList, tokenItem)
	}
	return tokenList, nil
}
