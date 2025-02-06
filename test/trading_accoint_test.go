package test

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/wsy0414/go-okx/okx"
)

func TestGetInstruments(t *testing.T) {
	okxClient := okx.NewOkxClient("{YOUR_ACCESS_KEY}", "{YOUR_SECRET_KEY}", "{YOUR_PASSPHRASE}")
	result, err := okxClient.GetInstruments(okx.Swap, nil, nil, nil)
	assert.Equal(t, err, nil)
	fmt.Println(result)
}

func TestGetBalance(t *testing.T) {
	okxClient := okx.NewOkxClient("{YOUR_ACCESS_KEY}", "{YOUR_SECRET_KEY}", "{YOUR_PASSPHRASE}")
	ccy := "USDT"
	result, err := okxClient.GetBalance(&ccy)
	assert.Equal(t, err, nil)
	fmt.Println(result)
}
