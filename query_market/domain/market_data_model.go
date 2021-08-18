package domain

type MarketDataModel struct {
	Id         string
	Time       string
	OpenValue  float32
	HighValue  float32
	LowValue   float32
	CloseValue float32
	Symbol     string
}
