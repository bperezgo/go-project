package app

import (
	"github.com/bperezgo/go-project/query_market/domain"
)

type MarketService struct {
	domain.MarketRepository
}

func NewMarketService() *MarketService {
	// This is a begin of dependency injection
	return &MarketService{
		MarketRepository: domain.MarketRepository{},
	}
}

func (s *MarketService) SaveData(marketData []domain.MarketData) error {
	return s.MarketRepository.SaveMarketData(marketData)
}
