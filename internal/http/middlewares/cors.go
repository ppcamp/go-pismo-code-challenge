package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-pismo-code-challenge/internal/config"
	"github.com/spf13/viper"
)

func Cors() gin.HandlerFunc {
	c := cors.Config{
		AllowOrigins:     viper.GetStringSlice(config.CorsAllowedOrigins), // Allow specific origins
		AllowMethods:     viper.GetStringSlice(config.CorsAllowedOrigins), // Allow specific methods
		AllowHeaders:     viper.GetStringSlice(config.CorsAllowedHeaders), // Allow specific headers
		AllowCredentials: true,                                            // Allow cookies or credentials
	}

	return cors.New(c)

}
