package main

import (
	"loadTester/handlers"
	"loadTester/types"
	"log"
	"net/http"

	"github.com/Depado/ginprom"

	"github.com/gin-gonic/gin"
)

type tester interface {
	TestEndpoint() types.Result
}

func main() {
	r := gin.Default()
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	r.Use(p.Instrument())
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
			"status":  http.StatusOK,
		})
	})
	r.POST("/http_test", handlers.PostHandlers)
	log.Fatalln(r.Run())
}
