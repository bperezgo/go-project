package domain

import (
	"github.com/satori/go.uuid"
	"time"
)

type MarketDataModel struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Time       time.Time
	OpenValue  float32
	HighValue  float32
	LowValue   float32
	CloseValue float32
	Symbol     string
}
