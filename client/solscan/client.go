package solscan

import (
	"github.com/web3-fighter/chain-explorer-api/client"
	"github.com/web3-fighter/chain-explorer-api/client/base"
	"github.com/web3-fighter/chain-explorer-api/client/unimplement"
	"time"
)

var _ client.ChainExplorer = (*ChainExplorerClient)(nil)

const ChainExplorerName = "solscan"

type ChainExplorerClient struct {
	unimplement.UnimplementedChainExplorerClient
	// https://public-api.solscan.io, 没有提供GetEstimateGasFee接口
	baseClient *base.BaseClient
}

func NewChainExplorerClient(key string, baseURL string, verbose bool, timeout time.Duration) (*ChainExplorerClient, error) {
	baseClient := base.NewBaseClient(key, baseURL, verbose, timeout)
	return &ChainExplorerClient{
		baseClient: baseClient,
	}, nil
}
