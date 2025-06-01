package examples

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/client/etherscan"
	"github.com/web3-fighter/chain-explorer-api/client/oklink"
	"time"
)

var (
	//EtherscanBaseUrl = "https://api.etherscan.io/api?"
	//EtherscanApiKey  = "HZEZGEPJJDA633N421AYW9NE8JFNZZC7JT"
	//EtherscanTimeout = time.Second * 20

	OklinkBaseUrl = "https://www.oklink.com"
	OklinkApiKey  = "5181d535-b68f-41cf-bbc6-25905e46b6a6"
	OkTimeout     = time.Second * 20

	// test success
	EtherscanBaseUrl = "https://api-holesky.etherscan.io/api?"
	EtherscanApiKey  = "HZEZGEPJJDA633N421AYW9NE8JFNZZC7JT"
	EtherscanTimeout = time.Second * 20
)

func NewMockClient() (*oklink.OKLinkExplorerClient, *etherscan.ChainExplorerClient, error) {
	var err error
	oklinkCli, err := oklink.NewChainExplorerClient(OklinkApiKey, OklinkBaseUrl, false, time.Duration(OkTimeout))
	if err != nil {
		fmt.Println("Mock oklink client fail", "err", err)
		return nil, nil, err
	}

	etherscanCli, err := etherscan.NewChainExplorerClient(EtherscanApiKey, EtherscanBaseUrl, false, time.Duration(EtherscanTimeout))
	if err != nil {
		fmt.Println("Mock etherscan client fail", "err", err)
		return nil, nil, err
	}
	return oklinkCli, etherscanCli, nil
}
