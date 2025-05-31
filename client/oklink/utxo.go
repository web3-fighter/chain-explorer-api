package oklink

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/types"
)

func (c *OKLinkExplorerClient) GetAccountUtxo(request *types.AccountUtxoRequest) ([]*types.AccountUtxoResponse, error) {
	var aurList []*types.AccountUtxoResponse
	apiUrl := fmt.Sprintf("api/v5/explorer/address/utxo?chainShortName=%s&address=%s", request.ChainShortName, request.Address)
	if request.Page != "" {
		apiUrl += fmt.Sprintf("&page=%s", request.Page)
	}
	if request.Limit != "" {
		apiUrl += fmt.Sprintf("&limit=%s", request.Limit)
	}

	var responseData []AddressUtxoData
	err := c.baseClient.Call(ChainExplorerName, "", "", apiUrl, nil, &responseData)
	if err != nil {
		return nil, err
	}

	for _, responseValue := range responseData {
		for _, utxoItem := range responseValue.UtxoList {
			aurList = append(aurList, &types.AccountUtxoResponse{
				TxId:          utxoItem.TxId,
				Height:        utxoItem.Height,
				BlockTime:     utxoItem.BlockTime,
				Address:       utxoItem.Address,
				UnspentAmount: utxoItem.UnspentAmount,
				Index:         utxoItem.Index,
			})
		}
	}
	return aurList, nil
}
