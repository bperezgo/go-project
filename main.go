package main

import (
	"fmt"
	"github.com/bperezgo/go-project/query_market/article"

	"github.com/gin-gonic/gin"
	"github.com/tkanos/gonfig"
)

func main() {
	configuration := Configuration{}
	err := gonfig.GetConf("configuration.json", &configuration)
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.GET("/ping", article.QueryMarket)
	router.Run(fmt.Sprintf(":%d", configuration.Port))
}
