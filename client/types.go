package client

import (
	"github.com/web3-fighter/chain-explorer-api/types"
)

type ChainExplorer interface {
	// GetChainExplorer 获取链信息
	GetChainExplorer(req *types.SupportChainExplorerRequest) (*types.SupportChainExplorerResponse, error)
	// GetAccountBalance 获取账户余额
	GetAccountBalance(req *types.AccountBalanceRequest) (*types.AccountBalanceResponse, error)
	// GetMultiAccountBalance 获取多个账户余额
	GetMultiAccountBalance(req *types.AccountBalanceRequest) ([]*types.AccountBalanceResponse, error)
	// GetAccountUtxo 获取账户utxo
	GetAccountUtxo(req *types.AccountUtxoRequest) ([]*types.AccountUtxoResponse, error)
	// GetTxByAddress 根据地址获取交易信息
	GetTxByAddress(request *types.AccountTxRequest) (*types.TransactionResponse[types.AccountTxResponse], error)
	// GetTokenList 获取代币列表
	GetTokenList(request *types.TokenRequest) ([]*types.TokenResponse, error)
	// GetEstimateGasFee 获取预估手续费
	GetEstimateGasFee(req *types.GasEstimateFeeRequest) (*types.GasEstimateFeeResponse, error)
	// GetTxByHash 根据交易hash获取交易信息
	GetTxByHash(request *types.TxRequest) (*types.TxResponse, error)
}
