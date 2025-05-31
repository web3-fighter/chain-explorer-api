package unimplement

import (
	"github.com/web3-fighter/chain-explorer-api/client"
	"github.com/web3-fighter/chain-explorer-api/client/base"
	"github.com/web3-fighter/chain-explorer-api/types"
)

var _ client.ChainExplorer = (*UnimplementedChainExplorerClient)(nil)

const ChainExplorerName = "unimplemented"

type UnimplementedChainExplorerClient struct {
	baseClient *base.BaseClient
}

func (c *UnimplementedChainExplorerClient) GetAccountBalance(req *types.AccountBalanceRequest) (*types.AccountBalanceResponse, error) {
	return &types.AccountBalanceResponse{}, nil
}

func (c *UnimplementedChainExplorerClient) GetMultiAccountBalance(req *types.AccountBalanceRequest) ([]*types.AccountBalanceResponse, error) {
	return []*types.AccountBalanceResponse{}, nil
}

func (c *UnimplementedChainExplorerClient) GetAccountUtxo(req *types.AccountUtxoRequest) ([]*types.AccountUtxoResponse, error) {
	return []*types.AccountUtxoResponse{}, nil
}

func (c *UnimplementedChainExplorerClient) GetTxByAddress(request *types.AccountTxRequest) (*types.TransactionResponse[types.AccountTxResponse], error) {
	return &types.TransactionResponse[types.AccountTxResponse]{}, nil
}

func (c *UnimplementedChainExplorerClient) GetTokenList(request *types.TokenRequest) ([]*types.TokenResponse, error) {
	return []*types.TokenResponse{}, nil
}

func (c *UnimplementedChainExplorerClient) GetEstimateGasFee(req *types.GasEstimateFeeRequest) (*types.GasEstimateFeeResponse, error) {
	return &types.GasEstimateFeeResponse{}, nil
}

func (c *UnimplementedChainExplorerClient) GetTxByHash(request *types.TxRequest) (*types.TxResponse, error) {
	return &types.TxResponse{}, nil
}
