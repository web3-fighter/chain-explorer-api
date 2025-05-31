package oklink

import (
	"errors"
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/types"
)

func (c *OKLinkExplorerClient) GetEstimateGasFee(req *types.GasEstimateFeeRequest) (*types.GasEstimateFeeResponse, error) {
	apiUrl := fmt.Sprintf("api/v5/explorer/blockchain/fee?chainShortName=%s", req.ChainShortName)
	var resp []*types.GasEstimateFeeResponse
	err := c.baseClient.Call(ChainExplorerName, "", "", apiUrl, nil, &resp)
	if err != nil {
		return nil, err
	}
	// 检查切片的长度
	if len(resp) == 0 {
		return nil, errors.New("no data returned from API")
	}

	return resp[0], nil
}
