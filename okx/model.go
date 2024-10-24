package okx

import "time"

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
