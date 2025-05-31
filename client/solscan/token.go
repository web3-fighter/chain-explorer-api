package solscan

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/pkg/common"
	"github.com/web3-fighter/chain-explorer-api/types"
	"strconv"
)

var pageSizes = []string{"10", "20", "30", "40", "60", "100"}

// 判断字符串是否在数组中
func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func (c *ChainExplorerClient) GetTokenList(req *types.TokenRequest) ([]*types.TokenResponse, error) {
	var responseList []*types.TokenResponse
	var responseData TokenListResp
	// 判断req.Limit是否在pageSizes内
	if !contains(pageSizes, req.Limit) {
		return responseList, fmt.Errorf("limit must be one of %v", pageSizes)
	}
	params := common.M{
		"page":      req.Page,
		"page_size": req.Limit,
	}
	err := c.baseClient.Call(ChainExplorerName, "", "", "/v1.0/account/splTransfers", params, &responseData)
	if err != nil {
		return responseList, err
	}
	for _, t := range responseData.Data {
		responseList = append(responseList, &types.TokenResponse{
			Symbol:               t.TokenSymbol,
			TokenContractAddress: t.MintAddress,
			TokenId:              t.Address,
			TotalSupply:          t.Supply.UiAmountString,
			Decimal:              strconv.Itoa(t.Decimals),
		})
	}
	return responseList, nil
}
