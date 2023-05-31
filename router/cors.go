package router

import (
	"time"

	"github.com/gin-contrib/cors"
)

func CorsConfig() cors.Config {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT"}
	corsConfig.AllowHeaders = []string{"Authorization", "Origin"}
	corsConfig.AllowCredentials = true
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.MaxAge = 12 * time.Hour

	return corsConfig
}
