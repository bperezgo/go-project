package article

import (
	"fmt"

	"github.com/bperezgo/go-project/config"
	"github.com/bperezgo/go-project/query_market/app"
	"github.com/bperezgo/go-project/query_market/article/validators"
	"github.com/gin-gonic/gin"
)

func QueryMarket(c *gin.Context) {
	var body validators.MarketValidator
	if errv := c.Bind(&body); errv != nil {
		fmt.Println(errv)
	}
	env := config.GetConfiguration()
	vantageProvider := app.VantageProvider{
		Url:    env.MarketDataProviderUrl,
		ApiKey: env.AlphaVantageApiKey,
	}
	options := app.MarketProviderOptions{Symbol: body.Symbol}
	marketData, err := vantageProvider.GetData(options)
	if err != nil {
		fmt.Println(err)
	}
	marketService := app.NewMarketService()
	err1 := marketService.SaveMarketData(marketData)
	if err1 != nil {
		fmt.Println(err)
	}
	c.JSON(201, gin.H{
		"message": "pong",
	})
}
