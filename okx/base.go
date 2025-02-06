package okx

const OKX_HEADER_ACCESS_KEY = "OK-ACCESS-KEY"
const OKX_HEADER_ACCESS_SIGN = "OK-ACCESS-SIGN"
const OKX_HEADER_ACCESS_TIMESTAMP = "OK-ACCESS-TIMESTAMP"
const OKX_HEADER_ACCESS_PASSPHRASE = "OK-ACCESS-PASSPHRASE"
const domain = "https://www.okx.com"

type InstrumentType string

const (
	Spot    InstrumentType = "SPOT"
	Margin  InstrumentType = "MARGIN"
	Swap    InstrumentType = "SWAP"
	Futures InstrumentType = "FUTURES"
	Option  InstrumentType = "OPTION"
)
