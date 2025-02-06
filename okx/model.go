package okx

import (
	"time"
)

type InstumentResponse struct {
	InstType       string
	InstId         string
	Uly            string
	InstFamily     string
	BaseCcy        string
	QuoteCcy       string
	SettleCcy      string
	CtVal          string
	CtMult         string
	CtValCcy       string
	OptType        string
	Stk            string
	ListTime       string
	AuctionEndTime string
	ExpTime        string
	Lever          string
	TickSz         string
	LotSz          string
	MinSz          string
	CtType         string
	State          string
	RuleType       string
	MaxLmtSz       string
	MaxMktSz       string
	MaxLmtAmt      string
	MaxMktAmt      string
	MaxTwapSz      string
	MaxIcebergSz   string
	MaxTriggerSz   string
	MaxStopSz      string
}

type BalanceResponse struct {
	UTime   string // Update time of account information, millisecond format of Unix timestamp, e.g. 1597026383085
	TotalEq string // The total amount of equity in USD
	IsoEq   string // Isolated margin equity in USD
	AdjEq   string // Adjusted / Effective equity in USD
	// The net fiat value of the assets in the account that can provide margins for spot, expiry futures, perpetual futures and options under the cross-margin mode.
	// In multi-ccy or PM mode, the asset and margin requirement will all be converted to USD value to process the order check or liquidation.
	// Due to the volatility of each currency market, our platform calculates the actual USD value of each currency based on discount rates to balance market risks.
	// Applicable to Spot mode/Multi-currency margin and Portfolio margin
	OrdFroz string // Cross margin frozen for pending orders in USD
	// Only applicable to Spot mode/Multi-currency margin/Portfolio margin
	Imr string // Initial margin requirement in USD
	// The sum of initial margins of all open positions and pending orders under cross-margin mode in USD.
	// Applicable to Spot mode/Multi-currency margin/Portfolio margin
	Mmr string // Maintenance margin requirement in USD
	// The sum of maintenance margins of all open positions and pending orders under cross-margin mode in USD.
	// Applicable to Spot mode/Multi-currency margin/Portfolio margin
	BorrowFroz string // Potential borrowing IMR of the account in USD
	// Only applicable to Spot mode/Multi-currency margin/Portfolio margin. It is "" for other margin modes.
	MgnRatio string // Margin ratio in USD
	// Applicable to Spot mode/Multi-currency margin/Portfolio margin
	NotionalUsd string // Notional value of positions in USD
	// Applicable to Spot mode/Multi-currency margin/Portfolio margin
	Upl string // Cross-margin info of unrealized profit and loss at the account level in USD
	// Applicable to Multi-currency margin/Portfolio margin
	Details []BalancdDetail
}
type BalancdDetail struct {
	Ccy               string // Currency
	Eq                string // Equity of currency
	CashBal           string // Cash balance
	UTime             string // Update time of currency balance information, Unix timestamp format in milliseconds, e.g. 1597026383085
	IsoEq             string // Isolated margin equity of currency
	AvailEq           string // Available equity of currency
	DisEq             string // Discount equity of currency in USD.
	FixedBal          string // Frozen balance for Dip Sniper and Peak Sniper
	AvailBal          string // Available balance of currency
	FrozenBal         string // Frozen balance of currency
	OrdFrozen         string // Margin frozen for open orders
	Liab              string // Liabilities of currency
	Upl               string // The sum of the unrealized profit & loss of all margin and derivatives positions of currency.
	UplLiab           string // Liabilities due to Unrealized loss of currency
	CrossLiab         string // Cross liabilities of currency
	RewardBal         string // Trial fund balance
	IsoLiab           string // Isolated liabilities of currency
	MgnRatio          string // Cross margin ratio of currency
	Imr               string // Cross initial margin requirement at the currency level
	Mmr               string // Cross maintenance margin requirement at the currency level
	Interest          string // Accrued interest of currency
	Twap              string // Risk indicator of auto liability repayment
	MaxLoan           string // Max loan of currency
	EqUsd             string // Equity in USD of currency
	BorrowFroz        string // Potential borrowing IMR of currency in USD
	NotionalLever     string // Leverage of currency
	StgyEq            string // Strategy equity
	IsoUpl            string // Isolated unrealized profit and loss of currency
	SpotInUseAmt      string // Spot in use amount
	ClSpotInUseAmt    string // User-defined spot risk offset amount
	MaxSpotInUse      string // Max possible spot risk offset amount
	SpotIsoBal        string // Spot isolated balance
	SmtSyncEq         string // Smart sync equity
	SpotCopyTradingEq string // Spot smart sync equity.
	SpotBal           string // Spot balance. The unit is currency, e.g. BTC. More details
	OpenAvgPx         string // Spot average cost price. The unit is USD. More details
	AccAvgPx          string // Spot accumulated cost price. The unit is USD. More details
	SpotUpl           string // Spot unrealized profit and loss. The unit is USD. More details
	SpotUplRatio      string // Spot unrealized profit and loss ratio. More details
	TotalPnl          string // Spot accumulated profit and loss. The unit is USD. More details
	TotalPnlRatio     string // Spot accumulated profit and loss ratio. More details
}
type CandeleRequest struct {
	InstrumentID string
	Bar          string
	After        time.Time
	Before       time.Time
	Limit        int
}

type CandleRespone struct {
	BarTime             time.Time
	OpenPrice           string
	ClosePrice          string
	LowestPrice         string
	HighestPrice        string
	Volume              string
	VolumeCurrency      string
	VolumeCurrencyQuote string
	Confirm             string
}
