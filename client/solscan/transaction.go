package solscan

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/types"
	"strconv"
	"time"
)

// GetTxByHash
/*
✅ SPL Token 的 tokenAddress ≠ tokenId（NFT）
	我们需要明确一点：在不同类型的资产中，tokenId 的含义是不同的。

一、对于 Fungible Token（可替代代币，例如 USDC、SRM 等）：
	概念	含义
		Token 合约地址	一个 mint 地址，比如 USDC 的 mint
		TokenId	❌ 没有 tokenId！
		SPL 中每种代币（如 USDC）由一个 mint 合约地址（tokenAddress）唯一标识。
		每个用户持有的 token 实际上是由 token account 表示。
		这些代币是可替代的（Fungible），所以根本就没有 tokenId 的必要。

二、对于 NFT（Non-Fungible Token）：
	概念	含义
		Token 合约地址	mint 地址，唯一标识该 NFT
		TokenId	✅ 是的，可以认为 tokenId = mint address（本体）

Solana 的 NFT（如 Metaplex 的标准）本身是一个特殊的 SPL Token：
	decimals = 0
	supply = 1
	每一个 NFT 的 mint 地址（即 tokenAddress）就 相当于 tokenId。
	所以在 NFT 场景中，tokenId 和 tokenAddress 是同一个。

✅ 总结归纳：
	资产类型	tokenAddress	tokenId
	Fungible Token（同质化代币）	合约地址（mint）	❌ 无 tokenId
	NFT（非同质化代币）	合约地址（mint）	✅ tokenId = mint 地址
*/
func (c *ChainExplorerClient) GetTxByHash(request *types.TxRequest) (*types.TxResponse, error) {
	apiUrl := fmt.Sprintf("v1.0/transaction/%s", request.Txid)
	var resp Transaction
	err := c.baseClient.Call(ChainExplorerName, "", "", apiUrl, nil, &resp)
	if err != nil {
		return nil, err
	}
	tokenTransferDetails := make([]types.TokenTransferDetail, 0, len(resp.ParsedInstruction))
	// 构造 token 转账详情
	for _, ins := range resp.ParsedInstruction {
		detail := types.TokenTransferDetail{
			Index:                "", // Solana 暂无 index 概念
			Token:                ins.Params.TokenAddress,
			TokenContractAddress: ins.Params.TokenAddress,
			From:                 ins.Params.Source,
			To:                   ins.Params.Destination,
			IsFromContract:       false,
			IsToContract:         false,
			TokenId:              "", // SPL token 没有 tokenId 概念，除非是 NFT mint
			Amount:               strconv.Itoa(ins.Params.Amount),
		}

		// 匹配补全 symbol 和 decimals
		for _, b := range resp.TokenBalances {
			if b.Account == ins.Params.Source || b.Account == ins.Params.Destination {
				detail.Decimals = b.Token.Decimals
				detail.Symbol = b.Token.Symbol
				break
			}
		}
		if isNFT(ins.Params.TokenAddress, resp.TokenBalances) {
			detail.TokenId = ins.Params.TokenAddress // NFT 的 tokenId 就是它的 mint
		}
		tokenTransferDetails = append(tokenTransferDetails, detail)
	}

	// 构造 InputDetails / OutputDetails
	inputDetails := make([]types.InputDetail, 0)
	outputDetails := make([]types.OutputDetail, 0)
	for _, acc := range resp.InputAccount {
		pre := acc.PreBalance
		post := acc.PostBalance
		diff := post - pre
		//谁花了钱（Pre > Post）
		//谁收了钱（Post > Pre）
		/*
			Solana 合约账户的判断需要额外调用：
			getAccountInfo 检查返回结构的 executable 字段：
				executable: true → 是合约（Program Account）
				executable: false → 是普通账户
		*/
		if diff < 0 {
			// 该账户支付了 SOL
			inputDetails = append(inputDetails, types.InputDetail{
				InputHash:  acc.Account,
				IsContract: false,               // 暂时无法判断是否合约
				Amount:     strconv.Itoa(-diff), // 转为正数
			})
		} else if diff > 0 {
			// 该账户收到了 SOL
			outputDetails = append(outputDetails, types.OutputDetail{
				OutputHash: acc.Account,
				IsContract: false,
				Amount:     strconv.Itoa(diff),
			})
		}
	}

	txResponse := &types.TxResponse{
		ChainFullName:        "Solana",
		ChainShortName:       "SOL",
		Txid:                 resp.TxHash,
		Height:               strconv.Itoa(resp.Slot),
		TransactionTime:      time.Unix(int64(resp.BlockTime), 0).Format(time.RFC3339),
		TransactionType:      "transfer",
		Amount:               strconv.Itoa(resp.Lamport),
		TransactionSymbol:    "SOL",
		MethodId:             "", // ❌ Solana 无 ABI / MethodId 概念
		ErrorLog:             "", // ❌ 如果 log 中出现 error 可补全
		InputData:            "", // ❌ 无原始 calldata（Solana 没有 input data 概念）
		Txfee:                strconv.Itoa(resp.Fee),
		Index:                "", // ❌ 无 txIndex 概念
		Confirm:              fmt.Sprintf("%v", resp.Confirmations),
		InputDetails:         inputDetails,
		OutputDetails:        outputDetails,
		State:                resp.Status,
		GasLimit:             "",    // ❌ Solana 非 EVM 架构，无 gasLimit
		GasUsed:              "",    // ❌ 同上
		GasPrice:             "",    // ❌ 同上
		TotalTransactionSize: "",    // ✅ 可补，若接口返回 raw size，可加
		VirtualSize:          "",    // ✅ 可加，Solana 暂未提供
		Weight:               "",    // ❌ 无 weight 概念（Bitcoin 特有）
		Nonce:                "",    // ❌ Solana 无 nonce 概念（由区块顺序/slot 控制）
		IsAaTransaction:      false, // ✅ Solana 无 AA 模式
		TokenTransferDetails: tokenTransferDetails,
		// TODO 如何提取合约地址
		ContractDetails: nil, // ✅ 可从内嵌 Instruction 中 programId 提取合约地址（可选）
	}
	return txResponse, nil
}

func isNFT(tokenAddress string, balances []TokenBalance) bool {
	for _, b := range balances {
		if b.Token.TokenAddress == tokenAddress && b.Token.Decimals == 0 {
			// 补充：还可以判断余额为 1
			preAmt, _ := strconv.Atoi(strconv.Itoa(b.Amount.PreAmount))
			postAmt, _ := strconv.Atoi(b.Amount.PostAmount)
			if preAmt <= 1 && postAmt <= 1 {
				return true
			}
		}
	}
	return false
}
