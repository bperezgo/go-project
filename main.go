package main

import (
	"fmt"

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
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(fmt.Sprintf(":%d", configuration.Port))
}
