package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-pismo-code-challenge/internal/config"
	"github.com/spf13/viper"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     viper.GetStringSlice(config.AppCorsAllowedOrigins),  // Allow specific origins
		AllowMethods:     viper.GetStringSlice(config.AppCorsAllowedOrigins),  // Allow specific methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allow specific headers
		AllowCredentials: true,                                                // Allow cookies or credentials
	})

}
