package oklink

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/types"
)

func (c *OKLinkExplorerClient) GetTxByHash(request *types.TxRequest) (*types.TxResponse, error) {
	apiUrl := fmt.Sprintf("api/v5/explorer/transaction/transaction-fills?chainShortName=%s&txid=%s", request.ChainShortName, request.Txid)
	var response []types.TxResponse
	err := c.baseClient.Call(ChainExplorerName, "", "", apiUrl, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response[0], nil
}
