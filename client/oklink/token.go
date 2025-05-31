package oklink

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/types"
)

// GetTokenList GET /api/v5/explorer/token/token-list
func (c *OKLinkExplorerClient) GetTokenList(request *types.TokenRequest) ([]*types.TokenResponse, error) {
	var responseList []*types.TokenResponse
	_protocolType := request.ProtocolType
	if types.IsBTCEcosystemProtocol(_protocolType) {
		apiUrl := fmt.Sprintf("api/v5/explorer/inscription/token-list?chainShortName=%s&protocolType=%s&tokenInscriptionId=%s&symbol=%s&projectId=%s&page=%s&limit=%s",
			request.ChainShortName, _protocolType, request.TokenInscriptionId, request.Symbol, request.ProjectId, request.Page, request.Limit)
		var responseData []TokenListData
		err := c.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			return nil, err
		}
		for _, tokenValue := range responseData[0].TokenList {
			responseList = append(responseList, &types.TokenResponse{
				Symbol:               tokenValue.Symbol,
				TokenContractAddress: tokenValue.TokenContractAddress,
				TokenId:              tokenValue.TokenInscriptionId,
				TotalSupply:          tokenValue.TotalSupply,
				Decimal:              tokenValue.Precision,
			})
		}
	} else {
		apiUrl := fmt.Sprintf("api/v5/explorer/token/token-list?chainShortName=%s&protocolType=%s&tokenContractAddress=%s&page=%s&limit=%s",
			request.ChainShortName, _protocolType, request.ContractAddress, request.Page, request.Limit)
		var responseData []TokenListData
		err := c.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			return nil, err
		}
		for _, tokenValue := range responseData[0].TokenList {
			responseList = append(responseList, &types.TokenResponse{
				Symbol:               tokenValue.Token,
				TokenContractAddress: tokenValue.TokenContractAddress,
				TokenId:              tokenValue.TokenId,
				TotalSupply:          tokenValue.TotalSupply,
				Decimal:              tokenValue.Precision,
			})
		}
	}
	return responseList, nil
}
