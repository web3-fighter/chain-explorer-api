package etherscan

import (
	"github.com/web3-fighter/chain-explorer-api/client"
	"github.com/web3-fighter/chain-explorer-api/client/base"
	"github.com/web3-fighter/chain-explorer-api/client/unimplement"
	"time"
)

var _ client.ChainExplorer = (*ChainExplorerClient)(nil)

const ChainExplorerName = "etherscan"

type ChainExplorerClient struct {
	unimplement.UnimplementedChainExplorerClient
	// https://api.etherscan.io/api?'
	baseClient *base.BaseClient
}

func NewChainExplorerClient(key string, baseURL string, verbose bool, timeout time.Duration) (*ChainExplorerClient, error) {
	baseClient := base.NewBaseClient(key, baseURL, verbose, timeout)
	return &ChainExplorerClient{
		baseClient: baseClient,
	}, nil
}
