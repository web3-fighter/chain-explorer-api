package solscan

import (
	"errors"
	"github.com/web3-fighter/chain-explorer-api/pkg/common"
	"github.com/web3-fighter/chain-explorer-api/types"
	"math/big"
	"strconv"
)

func (c *ChainExplorerClient) GetTxByAddress(request *types.AccountTxRequest) (*types.TransactionResponse[types.AccountTxResponse], error) {
	if request.Limit > 50 {
		return nil, errors.New("limit must be less than or equal to 50")
	}
	params := common.M{
		"offset":  request.Page,
		"limit":   request.Limit,
		"account": request.Address,
	}

	if request.Action == types.SolScanActionSol {
		resp := &AddressSolTransactionResp{}
		err := c.baseClient.Call(ChainExplorerName, "", "", "/v1.0/account/solTransfers", params, &resp)
		if err != nil {
			return nil, err
		}
		transactionList := make([]types.AccountTxResponse, 0, len(resp.Data))
		for _, tx := range resp.Data {
			operation := "inc"
			if tx.Src == request.Address {
				operation = "dec"
			}
			transactionList = append(transactionList, types.AccountTxResponse{
				TxId:         tx.TxHash,
				From:         tx.Src,
				To:           tx.Dst,
				Height:       strconv.FormatInt(tx.Slot, 10),
				Amount:       strconv.FormatInt(tx.Lamport, 10),
				TokenDecimal: strconv.FormatInt(tx.Decimals, 10),
				State:        tx.Status,
				TxFee:        strconv.FormatInt(tx.Fee, 10),
				Operation:    operation,
			})
		}
		return &types.TransactionResponse[types.AccountTxResponse]{
			PageResponse: types.PageResponse{
				Page:      request.Page,
				Limit:     request.Limit,
				TotalPage: resp.Total,
			},
			TransactionList: transactionList,
		}, nil
	}
	resp := &AddressSplTransactionResp{}
	err := c.baseClient.Call(ChainExplorerName, "", "", "/v1.0/account/splTransfers", params, &resp)
	if err != nil {
		return nil, err
	}
	transactionList := make([]types.AccountTxResponse, 0, len(resp.Data))
	for _, tx := range resp.Data {
		transactionList = append(transactionList, types.AccountTxResponse{
			Amount:       tx.ChangeAmount,
			Height:       strconv.FormatInt(tx.Slot, 10),
			TokenDecimal: strconv.FormatInt(tx.Decimals, 10),
			TxFee:        strconv.FormatInt(tx.Fee, 10),
			// TODO 非 nft 类型 	TokenId 就应该为 ""
			/*
				/v1.0/account/splTransfers 接口返回的格式，确实包含了 Token 的标识信息，
				但不是严格意义上的 “Token ID”，而是 SPL Token 的 Mint Address，即：
				TokenAddress 就是 Token 的唯一标识（Mint Address）
			*/
			TokenId: tx.TokenAddress,
			/*
				TokenContractAddress ≈ tokenAddress ≈ Token ID
				它们本质上 都是指 SPL Token 的 Mint Address，可以认为是：
				SPL Token 的唯一标识符

				tokenContractAddress	表示 Token 所在的合约地址（常用于以太坊 ERC20/721）	在 Solana 中，它指 Mint 地址
				tokenAddress	SPL Token 的 Mint 地址	✅ 标准叫法
				Token ID	通常也是指 Mint 地址，或 NFT 的 token_account 地址	在 Solana SPL Token 体系中就是 Mint 地址
			*/
			TokenContractAddress: tx.TokenAddress,
			TokenSymbol:          tx.Symbol,
			TokenName:            tx.TokenName,
			Operation:            tx.ChangeType,
		})
	}
	return &types.TransactionResponse[types.AccountTxResponse]{
		PageResponse: types.PageResponse{
			Page:      request.Page,
			Limit:     request.Limit,
			TotalPage: resp.Total,
		},
		TransactionList: transactionList,
	}, nil
}

// GetMultiAccountBalance 获取多个账户余额
func (c *ChainExplorerClient) GetMultiAccountBalance(req *types.AccountBalanceRequest) ([]*types.AccountBalanceResponse, error) {
	var responseData []AddressSummaryData
	err := c.baseClient.Call(ChainExplorerName, "", "", "/v1.0/account/tokens", common.M{
		"account": req.Account[0],
	}, &responseData)
	if err != nil {
		return nil, err
	}
	balanceList := make([]*types.AccountBalanceResponse, 0, len(req.Account))
	for _, balance := range responseData {
		balanceList = append(balanceList, &types.AccountBalanceResponse{
			Account:         req.Account[0],                                             // 账户地址
			Balance:         (*common.BigInt)(big.NewInt(balance.TokenAmount.UiAmount)), // 余额
			BalanceStr:      balance.TokenAmount.UiAmountString,                         // 余额字符串
			Symbol:          balance.TokenSymbol,                                        // 代币符号
			ContractAddress: balance.TokenAddress,                                       // 代币合约地址
			TokenId:         balance.TokenAccount,                                       // 代币ID
		})
	}

	return balanceList, nil
}

func (c *ChainExplorerClient) GetAccountBalance(req *types.AccountBalanceRequest) (*types.AccountBalanceResponse, error) {
	var responseData []AddressSummaryData
	err := c.baseClient.Call(ChainExplorerName, "", "", "/v1.0/account/tokens", common.M{
		"account": req.Account[0],
	}, &responseData)
	if err != nil {
		return nil, err
	}
	return &types.AccountBalanceResponse{
		Account:         req.Account[0],                                                     // 账户地址
		Balance:         (*common.BigInt)(big.NewInt(responseData[0].TokenAmount.UiAmount)), // 余额
		BalanceStr:      responseData[0].TokenAmount.UiAmountString,                         // 余额字符串
		Symbol:          responseData[0].TokenSymbol,                                        // 代币符号
		ContractAddress: responseData[0].TokenAddress,                                       // 代币合约地址
		TokenId:         responseData[0].TokenAccount,                                       // 代币ID
	}, nil
}
