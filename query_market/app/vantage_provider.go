package app

import (
	"encoding/json"
	"fmt"
	"github.com/bperezgo/go-project/query_market/domain"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"time"
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
	marketDataModel := []domain.MarketDataModel{}
	for date, v := range vantageResponse.TimeSeries {
		openValue := ParseToFloat32(*&v.OpenValue)
		closeValue := ParseToFloat32(*&v.CloseValue)
		highValue := ParseToFloat32(*&v.HighValue)
		lowValue := ParseToFloat32(*&v.LowValue)
		symbol := options.Symbol
		timeDate, _ := time.Parse("2006-01-02", date)
		marketDataModel = append(marketDataModel, domain.MarketDataModel{
			Id:         uuid.NewV4(),
			Time:       timeDate,
			OpenValue:  openValue,
			HighValue:  highValue,
			LowValue:   lowValue,
			CloseValue: closeValue,
			Symbol:     symbol,
		})
	}
	return marketDataModel, nil
}
