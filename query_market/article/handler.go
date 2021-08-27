package article

import (
	"fmt"

	"github.com/bperezgo/go-project/config"
	"github.com/bperezgo/go-project/lib/db"
	"github.com/bperezgo/go-project/query_market/app"
	"github.com/gin-gonic/gin"
)

func QueryMarket(c *gin.Context) {
	env := config.GetConfiguration()
	vantageProvider := app.VantageProvider{
		Url:    env.MarketDataProviderUrl,
		ApiKey: env.AlphaVantageApiKey,
	}
	options := app.MarketProviderOptions{Symbol: "IBM"}
	marketData, err := vantageProvider.GetData(options)
	conn := db.GetConnection()
	fmt.Printf("%v\n", marketData)
	fmt.Printf("%v\n", conn)
	// save data in the database
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
