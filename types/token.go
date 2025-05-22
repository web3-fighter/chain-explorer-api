package types

// type TokenConfig struct {
// 	ChainShortName string `json:"chainShortName"`
// 	ExplorerName   string `json:"explorerName"`
// }

// StandardTokenProtocols 20 tokens: token_20; 721 tokens: token_721; 1155 tokens: token_1155;
// Inscription token protocol type
// Runes: runes
// BRC-20 tokens: brc20
// SRC-20 tokens: src20
// ARC-20 tokens: arc20
// Ordinals NFT: ordinals_nft
type StandardTokenProtocols struct {
	Token20   string
	Token721  string
	Token1155 string
}

type BTCEcosystemProtocols struct {
	Runes       string
	BRC20       string
	SRC20       string
	ARC20       string
	OrdinalsNFT string
}

type Protocols struct {
	StandardTokenProtocols
	BTCEcosystemProtocols
}

var Protocol = Protocols{
	StandardTokenProtocols{
		Token20:   "token_20",
		Token721:  "token_721",
		Token1155: "token_1155",
	},
	BTCEcosystemProtocols{
		Runes:       "runes",
		BRC20:       "brc20",
		SRC20:       "src20",
		ARC20:       "arc20",
		OrdinalsNFT: "ordinals_nft",
	},
}

type TokenRequest struct {
	ChainShortName     string `json:"chainShortName"`
	ExplorerName       string `json:"explorerName"`
	ContractAddress    string `json:"contractAddress"`
	ProtocolType       string `json:"protocolType"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	ProjectId          string `json:"projectId"`
	Page               string `json:"page"`
	Limit              string `json:"limit"`
}

type TokenResponse struct {
	Symbol               string `json:"symbol"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenId              string `json:"tokenId"`
	TotalSupply          string `json:"totalSupply"` // TotalSupply
	Decimal              string `json:"decimal"`     //  default 1
}

// IsStandardProtocol Determine whether the given protocol type is a standard token protocol
/*
判断是否是「标准 Token 协议」
传入的 protocolType 是一个字符串，比如 "ERC20"、"ERC721" 或 "ERC1155"。
函数判断该协议是否是「标准 Token 协议」（如以太坊的 ERC 系列）。
如果是以下三种类型中的任意一种，就返回 true：
Token20（类似 ERC20）
Token721（类似 ERC721）
Token1155（类似 ERC1155）
否则返回 false。
*/
func IsStandardProtocol(protocolType string) bool {
	switch protocolType {
	case Protocol.StandardTokenProtocols.Token20, Protocol.StandardTokenProtocols.Token721, Protocol.StandardTokenProtocols.Token1155:
		return true
	default:
		return false
	}
}

// IsBTCEcosystemProtocol Determine whether the given protocol type is a special token protocol for the Bitcoin ecosystem
/*
判断是否是「比特币生态协议」
传入的 protocolType 是一个字符串，比如 "BRC20"、"Runes" 等。
函数判断该协议是否属于比特币生态中常见的一些协议：
Runes（Ordinals 团队的新代币协议）
BRC20（比特币上的代币标准）
SRC20、ARC20（是基于 Ordinals 的其他变种）
OrdinalsNFT（比特币链上的 NFT 协议）
如果匹配这些之一，就返回 true，否则返回 false。
*/
func IsBTCEcosystemProtocol(protocolType string) bool {
	switch protocolType {
	case Protocol.BTCEcosystemProtocols.Runes, Protocol.BTCEcosystemProtocols.BRC20, Protocol.BTCEcosystemProtocols.SRC20, Protocol.BTCEcosystemProtocols.ARC20, Protocol.BTCEcosystemProtocols.OrdinalsNFT:
		return true
	default:
		return false
	}
}
