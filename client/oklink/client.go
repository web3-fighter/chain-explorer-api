package oklink

import (
	"github.com/web3-fighter/chain-explorer-api/client"
	"github.com/web3-fighter/chain-explorer-api/client/base"
	"time"
)

var _ client.ChainExplorer = (*OKLinkExplorerClient)(nil)

const ChainExplorerName = "oklink"

type OKLinkExplorerClient struct {
	// https://www.oklink.com
	baseClient *base.BaseClient
}

func NewChainExplorerClient(key string, baseURL string, verbose bool, timeout time.Duration) (*OKLinkExplorerClient, error) {
	baseClient := base.NewBaseClient(key, baseURL, verbose, timeout)
	return &OKLinkExplorerClient{
		baseClient: baseClient,
	}, nil
}
