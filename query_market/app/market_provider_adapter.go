package app

import "github.com/bperezgo/go-project/query_market/domain"

type MarketProviderOptions struct {
	Symbol string
}

type IMarketProviderAdapter interface {
	GetData(options MarketProviderOptions) []domain.MarketData
}
