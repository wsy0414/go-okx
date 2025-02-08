package okx

import (
	"time"
)

type OkxResponse[T any] struct {
	Code string `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
}

type InstumentResponse struct {
	InstType       string `json:"instType"`
	InstId         string `json:"instId"`
	Uly            string `json:"uly"`
	InstFamily     string `json:"instFamily"`
	BaseCcy        string `json:"baseCcy"`
	QuoteCcy       string `json:"quoteCcy"`
	SettleCcy      string `json:"settleCcy"`
	CtVal          string `json:"ctVal"`
	CtMult         string `json:"ctMult"`
	CtValCcy       string `json:"ctValCcy"`
	OptType        string `json:"optType"`
	Stk            string `json:"stk"`
	ListTime       string `json:"listTime"`
	AuctionEndTime string `json:"auctionEndTime"`
	ExpTime        string `json:"expTime"`
	Lever          string `json:"lever"`
	TickSz         string `json:"tickSz"`
	LotSz          string `json:"lotSz"`
	MinSz          string `json:"minSz"`
	CtType         string `json:"ctType"`
	State          string `json:"state"`
	RuleType       string `json:"ruleType"`
	MaxLmtSz       string `json:"maxLmtSz"`
	MaxMktSz       string `json:"maxMktSz"`
	MaxLmtAmt      string `json:"maxLmtAmt"`
	MaxMktAmt      string `json:"maxMktAmt"`
	MaxTwapSz      string `json:"maxTwapSz"`
	MaxIcebergSz   string `json:"maxIcebergSz"`
	MaxTriggerSz   string `json:"maxTriggerSz"`
	MaxStopSz      string `json:"maxStopSz"`
}

type BalanceResponse struct {
	UTime   string `json:"uTime"`   // Update time of account information, millisecond format of Unix timestamp, e.g. 1597026383085
	TotalEq string `json:"totalEq"` // The total amount of equity in USD
	IsoEq   string `json:"isoEq"`   // Isolated margin equity in USD
	AdjEq   string `json:"adjEq"`   // Adjusted / Effective equity in USD
	// The net fiat value of the assets in the account that can provide margins for spot, expiry futures, perpetual futures and options under the cross-margin mode.
	// In multi-ccy or PM mode, the asset and margin requirement will all be converted to USD value to process the order check or liquidation.
	// Due to the volatility of each currency market, our platform calculates the actual USD value of each currency based on discount rates to balance market risks.
	// Applicable to Spot mode/Multi-currency margin and Portfolio margin
	OrdFroz string `json:"ordFroz"` // Cross margin frozen for pending orders in USD
	// Only applicable to Spot mode/Multi-currency margin/Portfolio margin
	Imr string `json:"imr"` // Initial margin requirement in USD
	// The sum of initial margins of all open positions and pending orders under cross-margin mode in USD.
	// Applicable to Spot mode/Multi-currency margin/Portfolio margin
	Mmr string `json:"mmr"` // Maintenance margin requirement in USD
	// The sum of maintenance margins of all open positions and pending orders under cross-margin mode in USD.
	// Applicable to Spot mode/Multi-currency margin/Portfolio margin
	BorrowFroz string `json:"borrowFroz"` // Potential borrowing IMR of the account in USD
	// Only applicable to Spot mode/Multi-currency margin/Portfolio margin. It is "" for other margin modes.
	MgnRatio string `json:"mgnRatio"` // Margin ratio in USD
	// Applicable to Spot mode/Multi-currency margin/Portfolio margin
	NotionalUsd string `json:"notionalUsd"` // Notional value of positions in USD
	// Applicable to Spot mode/Multi-currency margin/Portfolio margin
	Upl string // Cross-margin info of unrealized profit and loss at the account level in USD
	// Applicable to Multi-currency margin/Portfolio margin
	Details []BalancdDetail `json:"details"`
}
type BalancdDetail struct {
	Ccy               string `json:"ccy"`               // Currency
	Eq                string `json:"eq"`                // Equity of currency
	CashBal           string `json:"cashBal"`           // Cash balance
	UTime             string `json:"uTime"`             // Update time of currency balance information, Unix timestamp format in milliseconds, e.g. 1597026383085
	IsoEq             string `json:"isoEq"`             // Isolated margin equity of currency
	AvailEq           string `json:"availEq"`           // Available equity of currency
	DisEq             string `json:"disEq"`             // Discount equity of currency in USD.
	FixedBal          string `json:"fixedBal"`          // Frozen balance for Dip Sniper and Peak Sniper
	AvailBal          string `json:"availBal"`          // Available balance of currency
	FrozenBal         string `json:"frozenBal"`         // Frozen balance of currency
	OrdFrozen         string `json:"ordFrozen"`         // Margin frozen for open orders
	Liab              string `json:"liab"`              // Liabilities of currency
	Upl               string `json:"upl"`               // The sum of the unrealized profit & loss of all margin and derivatives positions of currency.
	UplLiab           string `json:"uplLiab"`           // Liabilities due to Unrealized loss of currency
	CrossLiab         string `json:"crossLiab"`         // Cross liabilities of currency
	RewardBal         string `json:"rewardBal"`         // Trial fund balance
	IsoLiab           string `json:"isoLiab"`           // Isolated liabilities of currency
	MgnRatio          string `json:"mgnRatio"`          // Cross margin ratio of currency
	Imr               string `json:"imr"`               // Cross initial margin requirement at the currency level
	Mmr               string `json:"mmr"`               // Cross maintenance margin requirement at the currency level
	Interest          string `json:"interest"`          // Accrued interest of currency
	Twap              string `json:"twap"`              // Risk indicator of auto liability repayment
	MaxLoan           string `json:"maxLoan"`           // Max loan of currency
	EqUsd             string `json:"eqUsd"`             // Equity in USD of currency
	BorrowFroz        string `json:"borrowFroz"`        // Potential borrowing IMR of currency in USD
	NotionalLever     string `json:"notionalLever"`     // Leverage of currency
	StgyEq            string `json:"stgyEq"`            // Strategy equity
	IsoUpl            string `json:"isoUpl"`            // Isolated unrealized profit and loss of currency
	SpotInUseAmt      string `json:"spotInUseAmt"`      // Spot in use amount
	ClSpotInUseAmt    string `json:"clSpotInUseAmt"`    // User-defined spot risk offset amount
	MaxSpotInUse      string `json:"maxSpotInUse"`      // Max possible spot risk offset amount
	SpotIsoBal        string `json:"spotIsoBal"`        // Spot isolated balance
	SmtSyncEq         string `json:"smtSyncEq"`         // Smart sync equity
	SpotCopyTradingEq string `json:"spotCopyTradingEq"` // Spot smart sync equity.
	SpotBal           string `json:"spotBal"`           // Spot balance. The unit is currency, e.g. BTC. More details
	OpenAvgPx         string `json:"openAvgPx"`         // Spot average cost price. The unit is USD. More details
	AccAvgPx          string `json:"accAvgPx"`          // Spot accumulated cost price. The unit is USD. More details
	SpotUpl           string `json:"spotUpl"`           // Spot unrealized profit and loss. The unit is USD. More details
	SpotUplRatio      string `json:"spotUplRatio"`      // Spot unrealized profit and loss ratio. More details
	TotalPnl          string `json:"totalPnl"`          // Spot accumulated profit and loss. The unit is USD. More details
	TotalPnlRatio     string `json:"totalPnlRatio"`     // Spot accumulated profit and loss ratio. More details
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
