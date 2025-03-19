package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HealthCheckHandler(c *gin.Context) {
	logrus.WithContext(c).Info("healthcheck")
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
