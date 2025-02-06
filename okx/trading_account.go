package okx

import (
	"errors"
	"net/url"
)

// Retrieve available instruments info of current account.
// instType: Instrument type
func (client *OkxClient) GetInstruments(instType InstrumentType, uly *string, instFamily *string, instId *string) ([]InstumentResponse, error) {
	if instType == "OPTION" {
		if uly == nil && instFamily == nil {
			return nil, errors.New("InvalidParam")
		}
	}
	param := url.Values{}
	param.Set("instType", string(instType))
	if uly != nil {
		param.Set("uly", *uly)
	}
	if instFamily != nil {
		param.Set("instFamily", *instFamily)
	}
	if instId != nil {
		param.Set("instId", *instId)
	}

	result := []InstumentResponse{}
	err := client.get("api/v5/account/instruments", nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *OkxClient) GetBalance(ccy *string) (*BalanceResponse, error) {
	param := url.Values{}
	if ccy != nil {
		param.Set("ccy", *ccy)
	}

	result := &BalanceResponse{}
	err := client.get("api/v5/account/balance", &param, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
