package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-pismo-code-challenge/internal/config"
	"github.com/spf13/viper"
)

func Handlers() http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	registerMiddlewares(r)
	registerAccountingRoutes(r)

	r.GET("/health", healthCheckHandler)

	return r.Handler()
}

func registerMiddlewares(r *gin.Engine) {
	r.Use(gin.Recovery())

	if viper.GetBool(config.LoggingHttpEnabled) {
		r.Use(gin.Logger())
	}
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func registerAccountingRoutes(r *gin.Engine) {
	group := r.Group("/accounts")

	group.GET("{:id}", func(c *gin.Context) {})

	group.POST("", func(c *gin.Context) {})
}
