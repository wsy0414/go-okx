package okx

import (
	"fmt"
	"net/url"
	"time"

	"github.com/wsy0414/go-okx/internal/client"
	"github.com/wsy0414/go-okx/internal/model"
)

func GetBalanceFromExchangeAccount() {
	m := make(map[string]any, 0)
	if err := client.Get(ACCOUNT_CONFIG, nil, &m); err != nil {
		// 應該會是log
		panic(err)
	}
	fmt.Println(m)
}

func GetCandles(coin string, bar string, t time.Time, limit int) ([][]string, error) {
	var result model.CandleResponse
	params := url.Values{}
	params.Set("instId", coin)
	params.Set("bar", bar)
	params.Set("after", fmt.Sprint(t.UnixMilli()))
	params.Set("limit", fmt.Sprint(limit))

	if err := client.Get(ORDER_BOOK_MARKET_DATA_CANDLES_HISTORY, &params, &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result.Data, nil
}
