package domain

import "github.com/bperezgo/go-project/lib/db"

type MarketRepository struct{}

func (r *MarketRepository) SaveMarketData(marketData []MarketData) error {
	conn := db.GetConnection()
	conn.Create(marketData)
	return nil
}
