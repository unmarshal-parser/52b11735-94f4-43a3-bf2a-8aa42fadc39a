// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"math/big"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = datatypes.JSON{}
	_ = time.Time{}
)

func GetEthWithdrawnEventHash() string {
	return "0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b"
}

type EthWithdrawnEvent struct {
	Target string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:30b9d940-793a-4d80-8f27-ec468dfaaca0,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:30b9d940-793a-4d80-8f27-ec468dfaaca0,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:30b9d940-793a-4d80-8f27-ec468dfaaca0,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSwapTargetAddedEventHash() string {
	return "0xb907822409611d127ab6a64611591b98e03a6a85ade4f258bae26b7c1efdfeaf"
}

type SwapTargetAddedEvent struct {
	Target string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:ffd0b7cc-9bca-4bc8-a39d-8bfa26fc7893,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:ffd0b7cc-9bca-4bc8-a39d-8bfa26fc7893,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:ffd0b7cc-9bca-4bc8-a39d-8bfa26fc7893,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSwapTargetRemovedEventHash() string {
	return "0x393b8be3e26787f19285ecd039dfd80bc6507828750f4d50367e6efe2524695c"
}

type SwapTargetRemovedEvent struct {
	Target string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:543c468c-635f-4a8e-9a52-50d1e0a18359,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:543c468c-635f-4a8e-9a52-50d1e0a18359,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:543c468c-635f-4a8e-9a52-50d1e0a18359,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTokenWithdrawnEventHash() string {
	return "0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620"
}

type TokenWithdrawnEvent struct {
	Token  string
	Target string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:d83f5814-d43a-4051-8fe8-5333fb5000a6,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:d83f5814-d43a-4051-8fe8-5333fb5000a6,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:d83f5814-d43a-4051-8fe8-5333fb5000a6,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToTokenWithPermitMethodHash() string {
	return "b0480bbd"
}

type FillQuoteTokenToTokenWithPermitMethod struct {
	SellTokenAddress string
	BuyTokenAddress  string
	Target           string
	SwapCallData     []byte
	SellAmount       decimal.Decimal `gorm:"type:numeric"`
	FeeAmount        decimal.Decimal `gorm:"type:numeric"`
	PermitData       datatypes.JSON

	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`
	TokenPriceSellAmount      float64 `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:2452b6c8-fe03-412e-98db-c7872f0806e9,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:2452b6c8-fe03-412e-98db-c7872f0806e9,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTransferOwnershipMethodHash() string {
	return "f2fde38b"
}

type TransferOwnershipMethod struct {
	NewOwner string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:83954f9a-9042-4ca1-985f-470ca8f23464,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:83954f9a-9042-4ca1-985f-470ca8f23464,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetUpdateSwapTargetsMethodHash() string {
	return "97bbda0e"
}

type UpdateSwapTargetsMethod struct {
	Target string
	Add    bool

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:bfdd3be7-8ec7-41bf-aa05-80190c6796a2,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:bfdd3be7-8ec7-41bf-aa05-80190c6796a2,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToEthMethodHash() string {
	return "999b6464"
}

type FillQuoteTokenToEthMethod struct {
	SellTokenAddress         string
	Target                   string
	SwapCallData             []byte
	SellAmount               decimal.Decimal `gorm:"type:numeric"`
	FeePercentageBasisPoints decimal.Decimal `gorm:"type:numeric"`

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeePercentageBasisPoints      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeePercentageBasisPoints float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:a39ccd50-bf75-4023-8d69-227f34da0680,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:a39ccd50-bf75-4023-8d69-227f34da0680,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToTokenMethodHash() string {
	return "55e4b7be"
}

type FillQuoteTokenToTokenMethod struct {
	SellTokenAddress string
	BuyTokenAddress  string
	Target           string
	SwapCallData     []byte
	SellAmount       decimal.Decimal `gorm:"type:numeric"`
	FeeAmount        decimal.Decimal `gorm:"type:numeric"`

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:52c58cae-a9d3-4299-8cf2-3047ad0c15be,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:52c58cae-a9d3-4299-8cf2-3047ad0c15be,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetWithdrawEthMethodHash() string {
	return "1b9a91a4"
}

type WithdrawEthMethod struct {
	To     string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:327c5cae-c828-4e00-a485-a66eb65f62b8,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:327c5cae-c828-4e00-a485-a66eb65f62b8,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetWithdrawTokenMethodHash() string {
	return "01e33667"
}

type WithdrawTokenMethod struct {
	Token  string
	To     string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:b483685a-4f5a-4eb0-a049-df975040f04c,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:b483685a-4f5a-4eb0-a049-df975040f04c,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteEthToTokenMethodHash() string {
	return "3c2b9a7d"
}

type FillQuoteEthToTokenMethod struct {
	BuyTokenAddress string
	Target          string
	SwapCallData    []byte
	FeeAmount       decimal.Decimal `gorm:"type:numeric"`

	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`
	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:ac636c10-3ceb-4890-bf4e-5ee43eb6a037,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:ac636c10-3ceb-4890-bf4e-5ee43eb6a037,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToEthWithPermitMethodHash() string {
	return "b3093838"
}

type FillQuoteTokenToEthWithPermitMethod struct {
	SellTokenAddress         string
	Target                   string
	SwapCallData             []byte
	SellAmount               decimal.Decimal `gorm:"type:numeric"`
	FeePercentageBasisPoints decimal.Decimal `gorm:"type:numeric"`
	PermitData               datatypes.JSON

	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`
	TokenPriceSellAmount      float64 `gorm:"type:numeric"`

	DecimalAdjustedFeePercentageBasisPoints float64 `gorm:"type:numeric"`
	TokenPriceFeePercentageBasisPoints      float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:ceba8432-c41c-41b3-a945-78682122e757,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:ceba8432-c41c-41b3-a945-78682122e757,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

type LastSyncedBlock struct {
	Contract    string `gorm:"primaryKey"`
	ChainID     string `gorm:"primaryKey"`
	SyncType    string `gorm:"primaryKey"`
	BlockNumber uint64
}

// Plugin Models

type TokenDetails struct {
	ID      int
	Address string `gorm:"uniqueIndex:address_and_chain"`
	Symbol  string
	ChainID string `gorm:"uniqueIndex:address_and_chain"`
	Decimal int
}

// Config

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
	TablePrefix      string `mapstructure:"table_prefix"`
	CreateBatchSize  int    `mapstructure:"create_batch_size"`
}

type IndexerConfig struct {
	EthEndpoint       string `mapstructure:"eth_endpoint"`
	ContractAddress   string `mapstructure:"contract_address"`
	StartBlock        int    `mapstructure:"start_block"`
	ApiKey            string `mapstructure:"api_key"`
	PostgresConfig    `mapstructure:"postgres_config"`
	LagToHighestBlock int `mapstructure:"lag_to_highest_block"`
	StepSize          int `mapstructure:"step_size"`
	ParallelCalls     int `mapstructure:"parallel_calls_for_logs"`
}

func (i *IndexerConfig) AssignDefaults() {
	if i.PostgresConfig.CreateBatchSize == 0 {
		i.PostgresConfig.CreateBatchSize = 100
	}
	if i.StepSize == 0 {
		i.StepSize = 50
	}
	if i.LagToHighestBlock == 0 {
		i.LagToHighestBlock = 10
	}
}
