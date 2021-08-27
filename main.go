package main

import (
	"fmt"
	"github.com/bperezgo/go-project/config"
	"github.com/bperezgo/go-project/query_market/article"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.GetConfiguration()
	router := gin.Default()
	router.GET("/query-market", article.QueryMarket)
	router.Run(fmt.Sprintf(":%d", config.Port))
}
