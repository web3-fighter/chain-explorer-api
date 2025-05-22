package etherscan

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/pkg/helper"
	"github.com/web3-fighter/chain-explorer-api/types"
)

func (c *ChainExplorerClient) GetAccountBalance(req *types.AccountBalanceRequest) (*types.AccountBalanceResponse, error) {
	balance := new(helper.BigInt)
	// 主币
	if req.ContractAddress[0] == "0x00" {
		param := helper.M{
			"tag":     "latest",
			"address": req.Account[0],
		}
		err := c.baseClient.Call(ChainExplorerName, "account", "balance", "", param, balance)
		if err != nil {
			fmt.Println("err", err)
			return &types.AccountBalanceResponse{}, nil
		}
	} else {
		// 代币
		param := helper.M{
			"contractaddress": req.ContractAddress[0],
			"address":         req.Account[0],
			"tag":             "latest",
		}
		err := c.baseClient.Call(ChainExplorerName, "account", "tokenbalance", "", param, &balance)
		if err != nil {
			fmt.Println("err", err)
		}
	}
	return &types.AccountBalanceResponse{
		Account:         req.Account[0],
		Balance:         balance,
		BalanceStr:      balance.Int().String(),
		Symbol:          req.Symbol[0],
		ContractAddress: req.ContractAddress[0],
		TokenId:         "0x0",
	}, nil
}

func (c *ChainExplorerClient) GetMultiAccountBalance(req *types.AccountBalanceRequest) ([]*types.AccountBalanceResponse, error) {
	var abrList []*types.AccountBalanceResponse
	// 批量接口不支持代币
	if req.ContractAddress[0] != "0x00" {
		return []*types.AccountBalanceResponse{}, nil
	}
	param := helper.M{
		"tag":     "latest",
		"address": req.Account,
	}
	balances := make([]AccountBalance, 0, len(req.Account))
	err := c.baseClient.Call(ChainExplorerName, "account", "balancemulti", "", param, &balances)
	if err != nil {
		fmt.Println("err", err)
		return []*types.AccountBalanceResponse{}, nil
	}
	for _, balance := range balances {
		abr := &types.AccountBalanceResponse{
			Account:         balance.Account,
			Balance:         balance.Balance,
			Symbol:          "ETH",
			ContractAddress: "0x0",
			TokenId:         "0x0",
		}
		abrList = append(abrList, abr)
	}
	return abrList, nil
}

func (c *ChainExplorerClient) GetTxByAddress(request *types.AccountTxRequest) (*types.TransactionResponse[types.AccountTxResponse], error) {
	resp := &[]AddressTransactionResp{}

	tempRequest := AddressTransactionRequest{
		Address:    request.Address,
		StartBlock: request.StartBlockHeight,
		EndBlock:   request.EndBlockHeight,
		Page:       request.Page,
		Offset:     request.Limit,
		Sort:       request.Sort.ToString(),
	}

	// normal transaction
	if request.Action == types.EtherscanActionTxList {
		err := c.baseClient.Call(ChainExplorerName, "account", "txlist", "", tempRequest.ToQueryParamMap(), &resp)
		if err != nil {
			fmt.Println("err", err)
			return &types.TransactionResponse[types.AccountTxResponse]{}, nil
		}
	}

	// internal transaction
	if request.Action == types.EtherscanActionTxListInternal {
		err := c.baseClient.Call(ChainExplorerName, "account", "txlistinternal", "", tempRequest.ToQueryParamMap(), &resp)
		if err != nil {
			fmt.Println("err", err)
			return &types.TransactionResponse[types.AccountTxResponse]{}, nil
		}
	}

	// token ERC20 transaction
	if request.Action == types.EtherscanActionTokenTx {
		err := c.baseClient.Call(ChainExplorerName, "account", "tokentx", "", tempRequest.ToQueryParamMap(), &resp)
		if err != nil {
			fmt.Println("err", err)
			return &types.TransactionResponse[types.AccountTxResponse]{}, nil
		}
	}

	// nft ERC721 transaction
	if request.Action == types.EtherscanActionTokenNftTx {
		err := c.baseClient.Call(ChainExplorerName, "account", "tokennfttx", "", tempRequest.ToQueryParamMap(), &resp)
		if err != nil {
			fmt.Println("err", err)
			return &types.TransactionResponse[types.AccountTxResponse]{}, nil
		}
	}

	// nft ERC1155 transaction
	if request.Action == types.EtherscanActionToken1155Tx {
		err := c.baseClient.Call(ChainExplorerName, "account", "token1155tx", "", tempRequest.ToQueryParamMap(), &resp)
		if err != nil {
			fmt.Println("err", err)
			return &types.TransactionResponse[types.AccountTxResponse]{}, nil
		}
	}

	var transactionList []types.AccountTxResponse
	for _, tx := range *resp {
		tempOkTx := types.AccountTxResponse{
			TxId:                 tx.Hash,
			BlockHash:            tx.BlockHash,
			Height:               tx.BlockNumber,
			TransactionTime:      tx.TimeStamp,
			TransactionIndex:     tx.TransactionIndex,
			From:                 tx.From,
			To:                   tx.To,
			Nonce:                tx.Nonce,
			Amount:               tx.Value,
			Symbol:               tx.TokenSymbol,
			Operation:            tx.Type,
			GasPrice:             tx.GasPrice,
			GasLimit:             tx.Gas,
			GasUsed:              tx.GasUsed,
			TxFee:                tx.CumulativeGasUsed,
			State:                tx.TxReceiptStatus,
			TransactionType:      "",
			Confirmations:        tx.Confirmations,
			IsError:              tx.IsError,
			TraceId:              tx.TraceId,
			Input:                tx.Input,
			MethodId:             tx.MethodId,
			FunctionName:         tx.FunctionName,
			TokenContractAddress: tx.ContractAddress,
			IsFromContract:       false,
			IsToContract:         false,
			TokenId:              tx.TokenID,
			TokenName:            tx.TokenName,
			TokenSymbol:          tx.TokenSymbol,
			TokenDecimal:         tx.TokenDecimal,
			TokenValue:           tx.TokenValue,
		}
		transactionList = append(transactionList, tempOkTx)
	}

	var tempTotal = request.Page
	if len(transactionList) >= int(request.Limit) {
		tempTotal++
	}
	pageResponse := types.PageResponse{
		Page:      request.Page,
		Limit:     request.Limit,
		TotalPage: tempTotal,
	}

	if transactionList == nil || len(transactionList) == 0 {
		transactionList = []types.AccountTxResponse{}
	}

	return &types.TransactionResponse[types.AccountTxResponse]{
		PageResponse:    pageResponse,
		TransactionList: transactionList,
	}, nil
}
