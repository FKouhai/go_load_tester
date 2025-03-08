package handlers

import (
	"loadTester/http_tester"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostHandlers(c *gin.Context) {
	var j http_tester.EndpointInfo
	if err := c.BindJSON(&j); err != nil {
		return
	}
	res := j.TestEndpoint()
	c.JSON(http.StatusOK, res)
}
