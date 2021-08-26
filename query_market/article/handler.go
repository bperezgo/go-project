package article

import (
	"fmt"

	"github.com/bperezgo/go-project/query_market/app"
	"github.com/gin-gonic/gin"
)

func QueryMarket(c *gin.Context) {
	vantageProvider := app.VantageProvider{
		Url:    "https://www.alphavantage.co",
		ApiKey: "Q91SW473ZUOZ9FZE",
	}
	options := app.MarketProviderOptions{Symbol: "IBM"}
	marketData, err := vantageProvider.GetData(options)
	fmt.Printf("%v\n", marketData)
	// save data in the database
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
