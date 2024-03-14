package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.ForwardedByClientIP = false
	r.GET("../resources", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "test-crypt.serv",
		})
	})

	r.Run("localhost:8080")
}
