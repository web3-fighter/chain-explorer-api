package etherscan

import (
	"fmt"
	"github.com/web3-fighter/chain-explorer-api/pkg/common"
	"github.com/web3-fighter/chain-explorer-api/types"
)

func (c *ChainExplorerClient) GetTokenList(req *types.TokenRequest) ([]*types.TokenResponse, error) {
	var responseData []TokensResp
	param := common.M{
		"contractaddress": req.ContractAddress,
	}
	err := c.baseClient.Call(ChainExplorerName, "token", "tokeninfo", "", param, &responseData)
	if err != nil {
		fmt.Println("call token list for etherscan fail", "err", err)
		return nil, err
	}
	var tokenList []*types.TokenResponse
	for _, tokenValue := range responseData {
		tokenItem := &types.TokenResponse{
			Symbol: tokenValue.Symbol,
			/*
				✅ 对于 ERC-20：
				Token Contract Address = Token ID
				因为每个 ERC-20 Token 是一个合约。
				所以合约地址就是该 Token 的唯一标识符。
				举例：
				USDT 合约地址：0xdAC17F958D2ee523a2206206994597C13D831ec7
				👉 就是 USDT 的「Token ID」
				✅ 对于 ERC-721（NFT）或 ERC-1155：
				Token Contract Address ≠ Token ID
				合约地址只是标识 NFT 属于哪个合约系列（例如某个 NFT 项目）。
				但每一个 NFT 还有一个唯一的 tokenId（通常是数字 ID）。

				举例：
					Bored Ape 合约地址：0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d
					Ape #123 的 tokenId: 123
					所以完整唯一标识是：

				{ contract: 0xbc4..., tokenId: 123 }
				🔍 总结：
				Token 类型	Token Contract Address	Token ID	是否一样
				ERC-20	✅ 就是唯一标识	✅ 就是合约地址	✅ 一样
				ERC-721	✅ 合约系列地址	❌ NFT编号	❌ 不一样
				ERC-1155	✅ 合约地址	❌ tokenId	❌ 不一样
			*/
			TokenContractAddress: tokenValue.ContractAddress,
			TokenId:              tokenValue.TokenId,
			TotalSupply:          tokenValue.TotalSupply,
			Decimal:              tokenValue.Divisor,
		}
		tokenList = append(tokenList, tokenItem)
	}
	return tokenList, nil
}
