package oklink

type ApiResponse[T any] struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// AddressSummaryData The Data field within the Response structure of
// Fundamental blockchain data -> Address Data -> Get basic address details
type AddressSummaryData struct {
	ChainFullName                 string `json:"chainFullName"`
	ChainShortName                string `json:"chainShortName"`
	Address                       string `json:"address"`
	ContractAddress               string `json:"contractAddress"`
	Balance                       string `json:"balance"`
	BalanceSymbol                 string `json:"balanceSymbol"`
	TransactionCount              string `json:"transactionCount"`
	Verifying                     string `json:"verifying"`
	SendAmount                    string `json:"sendAmount"`
	ReceiveAmount                 string `json:"receiveAmount"`
	TokenAmount                   string `json:"tokenAmount"`
	TotalTokenValue               string `json:"totalTokenValue"`
	CreateContractAddress         string `json:"createContractAddress"`
	CreateContractTransactionHash string `json:"createContractTransactionHash"`
	FirstTransactionTime          string `json:"firstTransactionTime"`
	LastTransactionTime           string `json:"lastTransactionTime"`
	Token                         string `json:"token"`
	Bandwidth                     string `json:"bandwidth"`
	Energy                        string `json:"energy"`
	VotingRights                  string `json:"votingRights"`
	UnclaimedVotingRewards        string `json:"unclaimedVotingRewards"`
	IsAaAddress                   bool   `json:"isAaAddress"`
}

// AddressTokenBalanceToken The TokenList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance details by address
type AddressTokenBalanceToken struct {
	Symbol               string `json:"symbol"`
	TokenContractAddress string `json:"tokenContractAddress"`
	HoldingAmount        string `json:"holdingAmount"`
	PriceUsd             string `json:"priceUsd"`
	ValueUsd             string `json:"valueUsd"`
	TokenId              string `json:"tokenId"`
}

// AddressTokenBalanceData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance details by address
type AddressTokenBalanceData struct {
	Limit     string                     `json:"limit"`
	Page      string                     `json:"page"`
	TotalPage string                     `json:"totalPage"`
	TokenList []AddressTokenBalanceToken `json:"tokenList"`
}

// AddressBalanceMultiBalance The BalanceList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get native token balance in batches
type AddressBalanceMultiBalance struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

// AddressBalanceMultiData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get native token balance in batches
type AddressBalanceMultiData struct {
	Symbol      string                       `json:"symbol"`
	BalanceList []AddressBalanceMultiBalance `json:"balanceList"`
}

// AddressTokenBalanceMultiBalance The BalanceList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance in batches
type AddressTokenBalanceMultiBalance struct {
	Address              string `json:"address"`
	HoldingAmount        string `json:"holdingAmount"`
	TokenContractAddress string `json:"tokenContractAddress"`

	Symbol   string `json:"symbol"`
	PriceUsd string `json:"priceUsd"`
	ValueUsd string `json:"valueUsd"`
	TokenId  string `json:"tokenId"`
}

// AddressTokenBalanceMultiData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance in batches
type AddressTokenBalanceMultiData struct {
	Page        string                            `json:"page"`
	Limit       string                            `json:"limit"`
	TotalPage   string                            `json:"totalPage"`
	BalanceList []AddressTokenBalanceMultiBalance `json:"balanceList"`
}

type TokenListInfo struct {
	Symbol               string `json:"symbol"`
	Token                string `json:"token"`         // 代币名字简称：USDC
	TokenId              string `json:"tokenId"`       // 默认为0
	TotalSupply          string `json:"totalSupply"`   // 最大供应量
	TokenFullName        string `json:"tokenFullName"` // 代币名字全称：USDCoin
	Precision            string `json:"precision"`     // 精度 默认为1
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenInscriptionId   string `json:"tokenInscriptionId"` // 铭文代币的铭文ID
}

type TokenListData struct {
	Limit     string          `json:"limit"`
	Page      string          `json:"page"`
	TotalPage string          `json:"totalPage"`
	TokenList []TokenListInfo `json:"tokenList"`
}

// AddressUtxoData The Data field within the Response structure of
// UTXO-specific data -> Get remaining UTXO addresses
type AddressUtxoData struct {
	Page      string            `json:"page"`
	Limit     string            `json:"limit"`
	TotalPage string            `json:"totalPage"`
	UtxoList  []AddressUtxoUtxo `json:"utxoList"`
}

// AddressUtxoUtxo The UtxoList field within the Data field within the Response structure of
// UTXO-specific data -> Get remaining UTXO addresses
type AddressUtxoUtxo struct {
	TxId          string `json:"txid"`
	Height        string `json:"height"`
	BlockTime     string `json:"blockTime"`
	Address       string `json:"address"`
	UnspentAmount string `json:"unspentAmount"`
	Index         string `json:"index"`
}
