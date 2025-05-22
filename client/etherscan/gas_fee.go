package etherscan

import (
	"github.com/web3-fighter/chain-explorer-api/types"
)

func (c *ChainExplorerClient) GetEstimateGasFee(request *types.GasEstimateFeeRequest) (*types.GasEstimateFeeResponse, error) {
	resp := &GasTrackerGasOracleResp{}
	err := c.baseClient.Call(ChainExplorerName, "gastracker", "gasoracle", "", nil, &resp)

	if err != nil {
		return nil, err
	}
	return &types.GasEstimateFeeResponse{
		ChainFullName:         "Ethereum",
		ChainShortName:        "ETH",
		Symbol:                "ETH",
		BestTransactionFee:    "",
		BestTransactionFeeSat: "",
		RecommendedGasPrice:   resp.ProposeGasPrice,
		RapidGasPrice:         resp.FastGasPrice,
		StandardGasPrice:      resp.ProposeGasPrice,
		SlowGasPrice:          resp.SafeGasPrice,
		BaseFee:               resp.SuggestBaseFee,
		GasUsedRatio:          resp.GasUsedRatio,
	}, nil
}
