package article

import (
	"log"

	"github.com/bperezgo/go-project/query_market/app"
	"github.com/gin-gonic/gin"
)

func QueryMarket(c *gin.Context) {
	vantageProvider := app.VantageProvider{
		Url:    "https://www.alphavantage.co",
		ApiKey: "Q91SW473ZUOZ9FZE",
	}
	options := app.MarketProviderOptions{Symbol: "IBM"}
	_, err := vantageProvider.GetData(options)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
