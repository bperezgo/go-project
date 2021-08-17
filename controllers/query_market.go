package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func QueryMarket(c *gin.Context) {
	req, err := http.NewRequest("GET", configuration.MarketDataProviderUrl)
}
