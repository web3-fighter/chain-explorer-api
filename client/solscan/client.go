package solscan

import (
	"github.com/web3-fighter/chain-explorer-api/client"
	"github.com/web3-fighter/chain-explorer-api/client/base"
	"github.com/web3-fighter/chain-explorer-api/client/unimplement"
)

var _ client.ChainExplorer = (*ChainExplorerClient)(nil)

const ChainExplorerName = "solscan"

type ChainExplorerClient struct {
	unimplement.UnimplementedChainExplorerClient
	baseClient *base.BaseClient
}
