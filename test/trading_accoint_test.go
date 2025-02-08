package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/wsy0414/go-okx/okx"
)

func TestGetInstruments(t *testing.T) {
	okxClient := okx.NewOkxClient("{YOUR_ACCESS_KEY}", "{YOUR_SECRET_KEY}", "{YOUR_PASSPHRASE}")
	result, err := okxClient.GetInstruments(okx.Spot, nil, nil, nil)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, len(result), 0)
}

func TestGetBalance(t *testing.T) {
	okxClient := okx.NewOkxClient("{YOUR_ACCESS_KEY}", "{YOUR_SECRET_KEY}", "{YOUR_PASSPHRASE}")
	ccy := "USDT"
	result, err := okxClient.GetBalance(&ccy)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, len(*result), 0)
}
