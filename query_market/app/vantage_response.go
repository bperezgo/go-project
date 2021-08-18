package app

type VantageValuesData struct {
	OpenValue  float32 `json:"1. open"`
	HighValue  float32 `json:"2. high"`
	LowValue   float32 `json:"3. low"`
	CloseValue float32 `json:"4. close"`
}

type TimeSeriesDailyStruct struct {
	Time map[string]VantageValuesData
}

type VantageResponse struct {
	TimeSeriesDaily TimeSeriesDailyStruct `json:"Time Series (Daily)"`
}
