package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-pismo-code-challenge/internal/config"
	"github.com/ppcamp/go-pismo-code-challenge/internal/http/handlers"
	"github.com/ppcamp/go-pismo-code-challenge/internal/http/middlewares"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Routes define the http entry point (routes and middlewares).
//
// This function is also responsible to create http handlers/controllers.
func Routes(h *handlers.Handler) http.Handler {
	r := gin.New()

	registerMiddlewares(r)

	registerAccountRoutes(r, h)
	registerTransactionRoutes(r, h)

	r.GET("/health", handlers.HealthCheckHandler)

	return r.Handler()
}

func registerMiddlewares(r *gin.Engine) {
	r.Use(
		middlewares.RequestId(),
		middlewares.Cors(),
		gin.Recovery(),
	)

	if viper.GetBool(config.LoggingHttpEnabled) {
		logrus.Debug("running gin in debug mode")
		r.Use(gin.Logger())
	}
}

func registerAccountRoutes(r *gin.Engine, handler *handlers.Handler) {
	h := handlers.NewAccountHandler(handler)

	group := r.Group("/accounts")

	group.GET("/:id", h.Get)
	group.POST("", h.Create)
}

func registerTransactionRoutes(r *gin.Engine, handler *handlers.Handler) {
	h := handlers.NewTransactionHandler(handler)

	group := r.Group("/transactions")

	group.POST("", h.Create)
}
