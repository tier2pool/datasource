package ethermine

import (
	"errors"
	"github.com/shopspring/decimal"
)

type Response struct {
	Status string      `json:"status"`
	Error  Error       `json:"error"`
	Data   interface{} `json:"data"`
}

type Error struct {
	error
}

func (e *Error) UnmarshalJSON(data []byte) error {
	e.error = errors.New(string(data))
	return nil
}

type MinerDashboard struct {
	Statistics       []MinerDashboardStatistic      `json:"statistics"`
	Workers          []MinerDashboardWorker         `json:"workers"`
	CurrentStatistic MinerDashboardCurrentStatistic `json:"currentStatistics"`
	Setting          MinerDashboardSetting          `json:"settings"`
}

type MinerDashboardStatistic struct {
	Time             int64           `json:"time"`
	ReportedHashrate decimal.Decimal `json:"reportedHashrate"`
	CurrentHashrate  decimal.Decimal `json:"currentHashrate"`
	ValidShares      int64           `json:"validShares"`
	InvalidShares    int64           `json:"invalidShares"`
	StaleShares      int64           `json:"staleShares"`
}

type MinerDashboardWorker struct {
	Worker           string          `json:"worker"`
	Time             int64           `json:"time"`
	LastSeen         int64           `json:"lastSeen"`
	ReportedHashrate decimal.Decimal `json:"reportedHashrate"`
	CurrentHashrate  decimal.Decimal `json:"currentHashrate"`
	ValidShares      int64           `json:"validShares"`
	InvalidShares    int64           `json:"invalidShares"`
	StaleShares      int64           `json:"staleShares"`
}

type MinerDashboardCurrentStatistic struct {
	Time             int64           `json:"time"`
	LastSeen         int64           `json:"lastSeen"`
	ReportedHashrate decimal.Decimal `json:"reportedHashrate"`
	AverageHashrate  decimal.Decimal `json:"averageHashrate"`
	CurrentHashrate  decimal.Decimal `json:"currentHashrate"`
	ValidShares      int64           `json:"validShares"`
	InvalidShares    int64           `json:"invalidShares"`
	StaleShares      int64           `json:"staleShares"`
	ActiveWorkers    int64           `json:"active_workers"`
	Unpaid           decimal.Decimal `json:"unpaid"`
	Unconfirmed      decimal.Decimal `json:"unconfirmed"`
}

type MinerDashboardSetting struct {
	Email     string          `json:"email"`
	Monitor   int             `json:"monitor"`
	MinPayout decimal.Decimal `json:"min_payout"`
}

type MinerStats struct {
	Time             int64           `json:"time"`
	LastSeen         int64           `json:"lastSeen"`
	ReportedHashrate decimal.Decimal `json:"reportedHashrate"`
	AverageHashrate  decimal.Decimal `json:"averageHashrate"`
	CurrentHashrate  decimal.Decimal `json:"currentHashrate"`
	ValidShares      int64           `json:"validShares"`
	InvalidShares    int64           `json:"invalidShares"`
	StaleShares      int64           `json:"staleShares"`
	ActiveWorkers    int64           `json:"activeWorkers"`
	Unpaid           decimal.Decimal `json:"unpaid"`
	Unconfirmed      decimal.Decimal `json:"unconfirmed"`
	CoinsPerMin      decimal.Decimal `json:"coinsPerMin"`
	USDPerMin        decimal.Decimal `json:"usdPerMin"`
	BTCPerMin        decimal.Decimal `json:"btcPerMin"`
}
