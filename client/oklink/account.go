package oklink

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/client/etherscan"
	"github.com/web3-fighter/chain-explorer-api/pkg/common"
	"github.com/web3-fighter/chain-explorer-api/types"
	"math/big"
	"strconv"
	"strings"
)

func (c *OKLinkExplorerClient) GetTxByAddress(request *types.AccountTxRequest) (*types.TransactionResponse[types.AccountTxResponse], error) {
	type TransactionType = types.AccountTxResponse
	type TransactionResponseType = etherscan.TransactionResponse[TransactionType]
	var resp []TransactionResponseType
	// utxo ??TODO 为什么oklink没有支持account？
	if request.Action == types.OkLinkActionUtxo {
		baseURL := "/api/v5/explorer/address/transaction-list"
		fullURL := fmt.Sprintf("%s?%s", baseURL, request.ToQueryUrl())
		err := c.baseClient.Call(ChainExplorerName, "", "", fullURL, nil, &resp)
		if err != nil {
			return nil, err
		}
		// utxo is transactionLists, not transactionList
		if len(resp) > 0 {
			resp[0].TransactionList = resp[0].TransactionLists
		}
	}

	// normal transaction
	if request.Action == types.OkLinkActionNormal {
		baseURL := "/api/v5/explorer/address/normal-transaction-list"
		fullURL := fmt.Sprintf("%s?%s", baseURL, request.ToQueryUrl())
		//resp := &ApiResponse[[]account.TransactionResponse[account.NormalTransaction]]{}
		err := c.baseClient.Call(ChainExplorerName, "", "", fullURL, nil, &resp)
		if err != nil {
			return nil, err
		}
	}

	// internal transaction
	if request.Action == types.OkLinkActionInternal {
		baseURL := "/api/v5/explorer/address/internal-transaction-list"
		fullURL := fmt.Sprintf("%s?%s", baseURL, request.ToQueryUrl())
		err := c.baseClient.Call(ChainExplorerName, "", "", fullURL, nil, &resp)
		if err != nil {
			return nil, err
		}
	}

	// token transaction
	if request.Action == types.OkLinkActionToken {
		baseURL := "/api/v5/explorer/address/token-transaction-list"
		fullURL := fmt.Sprintf("%s?%s", baseURL, request.ToQueryUrl())
		err := c.baseClient.Call(ChainExplorerName, "", "", fullURL, nil, &resp)
		if err != nil {
			return nil, err
		}
	}

	tempResp := resp[0]
	return &types.TransactionResponse[types.AccountTxResponse]{
		PageResponse: types.PageResponse{
			Page:      StringToUint64(tempResp.Page, request.Page),
			Limit:     StringToUint64(tempResp.Limit, request.Limit),
			TotalPage: StringToUint64(tempResp.TotalPage, request.Page+1),
		},
		TransactionList: tempResp.TransactionList,
	}, nil
}

func (c *OKLinkExplorerClient) GetMultiAccountBalance(request *types.AccountBalanceRequest) ([]*types.AccountBalanceResponse, error) {
	var abrpsList []*types.AccountBalanceResponse
	addressStr := make([]string, len(request.Account))
	for i, v := range request.Account {
		addressStr[i] = v
	}
	result := strings.Join(addressStr, ",")
	if request.ContractAddress[0] == "0x00" {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/balance-multi?chainShortName=%s&address=%s", request.ChainShortName, result)
		var responseData []AddressBalanceMultiData
		err := c.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			return nil, err
		}
		for _, dataValue := range responseData {
			for _, value := range dataValue.BalanceList {
				balance, _ := new(big.Int).SetString(value.Balance, 10)
				abrps := &types.AccountBalanceResponse{
					Account:         value.Address,
					Balance:         (*common.BigInt)(balance),
					BalanceStr:      value.Balance,
					Symbol:          dataValue.Symbol,
					ContractAddress: "0x00",
					//TokenId:         "0x00",
				}
				abrpsList = append(abrpsList, abrps)
			}
		}
	} else {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/token-balance-multi?chainShortName=%s&address=%s&protocolType=%s&page=%s&limit=%s", request.ChainShortName, result, request.ProtocolType[0], request.Page[0], request.Limit[0])
		var responseData []AddressTokenBalanceMultiData
		err := c.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			return nil, err
		}
		for _, dataValue := range responseData {
			for _, token := range dataValue.BalanceList {
				balance, _ := new(big.Int).SetString(token.HoldingAmount, 10)
				abrps := &types.AccountBalanceResponse{
					Account:         token.Address,
					Balance:         (*common.BigInt)(balance),
					BalanceStr:      token.HoldingAmount,
					Symbol:          token.Symbol,
					ContractAddress: token.TokenContractAddress,
					TokenId:         token.TokenId,
				}
				abrpsList = append(abrpsList, abrps)
			}
		}
	}
	return abrpsList, nil
}

// GetAccountBalance GET /api/v5/explorer/address/address-summary?chainShortName=eth&address=0x85c6627c4ed773cb7c32644b041f58a058b00d30
func (c *OKLinkExplorerClient) GetAccountBalance(request *types.AccountBalanceRequest) (*types.AccountBalanceResponse, error) {
	if request.ContractAddress[0] == "0x00" {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/address-summary?chainShortName=%s&address=%s", request.ChainShortName, request.Account[0])
		var responseData []AddressSummaryData
		err := c.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			return nil, err
		}
		balance, _ := new(big.Int).SetString(responseData[0].Balance, 10)
		return &types.AccountBalanceResponse{
			Account:         responseData[0].Address,
			Balance:         (*common.BigInt)(balance),
			BalanceStr:      responseData[0].Balance,
			Symbol:          responseData[0].BalanceSymbol,
			ContractAddress: responseData[0].CreateContractAddress,
			//TokenId:         "0x00",
		}, nil
	}
	apiUrl := fmt.Sprintf("api/v5/explorer/address/token-balance?chainShortName=%s&address=%s&tokenContractAddress=%s&protocolType=%s&limit=%d", request.ChainShortName, request.Account[0], request.ContractAddress[0], request.ProtocolType[0], request.Limit[0])
	var responseData []AddressTokenBalanceData
	err := c.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
	if err != nil {
		return nil, err
	}
	balance, _ := new(big.Int).SetString(responseData[0].TokenList[0].HoldingAmount, 10)
	return &types.AccountBalanceResponse{
		Account:         request.Account[0],
		Balance:         (*common.BigInt)(balance),
		BalanceStr:      responseData[0].TokenList[0].HoldingAmount,
		Symbol:          responseData[0].TokenList[0].Symbol,
		ContractAddress: responseData[0].TokenList[0].TokenContractAddress,
		TokenId:         responseData[0].TokenList[0].TokenId,
	}, nil
}

func StringToUint64(s string, defaultValue uint64) uint64 {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Println("err", err)
		return defaultValue
	}
	return value
}
