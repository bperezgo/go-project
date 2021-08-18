package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/bperezgo/go-project/query_market/domain"
)

type VantageProvider struct {
	Url    string
	ApiKey string
}

func (v *VantageProvider) GetData(options MarketProviderOptions) ([]domain.MarketDataModel, error) {
	var vantageResponse VantageResponse
	url := fmt.Sprintf("%s/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=%s&apikey=%s", v.Url, options.Symbol, v.ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseData, &vantageResponse); err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", vantageResponse)
	return []domain.MarketDataModel{}, nil
}
