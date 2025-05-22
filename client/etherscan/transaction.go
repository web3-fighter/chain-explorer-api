package etherscan

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/types"
)

func (c *ChainExplorerClient) GetTxByHash(request *types.TxRequest) (*types.TxResponse, error) {
	// 参数校验
	if request == nil || request.Txid == "" {
		return nil, fmt.Errorf("invalid TxRequest: txid is empty")
	}

	// 构造请求参数
	param := map[string]interface{}{
		"txhash": request.Txid,
	}

	// Etherscan 返回结构
	var rawResp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  struct {
			BlockNumber      string `json:"blockNumber"`
			TimeStamp        string `json:"timeStamp"`
			Hash             string `json:"hash"`
			From             string `json:"from"`
			To               string `json:"to"`
			Value            string `json:"value"`
			Gas              string `json:"gas"`
			GasPrice         string `json:"gasPrice"`
			GasUsed          string `json:"gasUsed"`
			Input            string `json:"input"`
			Nonce            string `json:"nonce"`
			ContractAddress  string `json:"contractAddress"`
			TransactionIndex string `json:"transactionIndex"`
			IsError          string `json:"isError"`
			TxreceiptStatus  string `json:"txreceipt_status"`
		} `json:"result"`
	}

	// 发起请求
	err := c.baseClient.Call(
		ChainExplorerName,
		"proxy",
		"eth_getTransactionByHash",
		"",
		param,
		&rawResp,
	)
	if err != nil {
		return nil, err
	}
	if rawResp.Status != "1" || rawResp.Result.Hash == "" {
		return nil, fmt.Errorf("no transaction found for hash %s", request.Txid)
	}

	// 封装为 TxResponse
	txResp := &types.TxResponse{
		ChainShortName:  request.ChainShortName,
		ChainFullName:   "Ethereum",
		Txid:            rawResp.Result.Hash,
		Height:          rawResp.Result.BlockNumber,
		TransactionTime: rawResp.Result.TimeStamp,
		Amount:          rawResp.Result.Value,
		MethodId:        extractMethodID(rawResp.Result.Input),
		InputData:       rawResp.Result.Input,
		Index:           rawResp.Result.TransactionIndex,
		Confirm:         "", // 可选: 二次查询最新区块高度对比计算确认数
		State:           txStatus(rawResp.Result.IsError, rawResp.Result.TxreceiptStatus),
		GasLimit:        rawResp.Result.Gas,
		GasUsed:         rawResp.Result.GasUsed,
		GasPrice:        rawResp.Result.GasPrice,
		Nonce:           rawResp.Result.Nonce,
	}
	return txResp, nil
}

func extractMethodID(input string) string {
	if len(input) >= 10 {
		return input[:10] // 0x + 8字节 method id
	}
	return ""
}

func txStatus(isError, receiptStatus string) string {
	if isError == "1" || receiptStatus == "0" {
		return "failed"
	}
	return "success"
}
