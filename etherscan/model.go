package etherscan

import (
	"github.com/shopspring/decimal"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type GasTrackerEstimate struct {
	decimal.Decimal
}

type GasTrackerOracle struct {
	LastBlock       int64           `json:"LastBlock"`
	SafeGasPrice    decimal.Decimal `json:"SafeGasPrice"`
	ProposeGasPrice decimal.Decimal `json:"ProposeGasPrice"`
	FastGasPrice    decimal.Decimal `json:"FastGasPrice"`
	SuggestBaseFee  decimal.Decimal `json:"suggestBaseFee"`
	GasUsedRatio    string          `json:"gasUsedRatio"`
}

type StatsETHSupply decimal.Decimal

type StatsETH2Supply struct {
	EthSupply   decimal.Decimal `json:"EthSupply"`
	Eth2Staking decimal.Decimal `json:"Eth2Staking"`
	BurntFees   decimal.Decimal `json:"BurntFees"`
}

type StatsETHPrice struct {
	ETHToBTC          decimal.Decimal `json:"ethbtc"`
	ETHToBTCTimestamp int64           `json:"ethbtc_timestamp"`
	ETHToUSD          decimal.Decimal `json:"ethusd"`
	ETHToUSDTimestamp int64           `json:"ethusd_timestamp"`
}
