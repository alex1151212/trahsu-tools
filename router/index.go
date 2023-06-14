package router

import (
	"net/http"
	"trahsu-tools/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(CorsConfig()))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	r.GET("/crawler", service.Crawler)
	r.GET("/punch_in", service.PunchIn)

	// r.GET("/crawler-ws", service.Crawler)

	return r
}
