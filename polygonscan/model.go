package polygonscan

import (
	"github.com/shopspring/decimal"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type LastPrice struct {
	Maticbtc          decimal.Decimal `json:"maticbtc"`
	MaticbtcTimestamp string          `json:"maticbtc_timestamp"`
	Maticusd          decimal.Decimal `json:"maticusd"`
	MaticusdTimestamp string          `json:"maticusd_timestamp"`
}
