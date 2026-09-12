package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() error {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	v1 := r.Group("api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}

	return r.Run(":5005")
}
