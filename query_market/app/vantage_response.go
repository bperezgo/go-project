package app

type VantageValuesData struct {
	OpenValue        string `json:"1. open"`
	HighValue        string `json:"2. high"`
	LowValue         string `json:"3. low"`
	CloseValue       string `json:"4. close"`
	AdjustedClose    string `json:"5. adjusted close"`
	Volume           string `json:"6. volume"`
	DividedAmount    string `json:"7. dividend amount"`
	SplitCoefficient string `json:"8. split coefficient"`
}

type VantageResponse struct {
	TimeSeries map[string]*VantageValuesData `json:"Time Series (Daily)"`
}
