package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/utils"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/utils/helpers"
	"github.com/sirupsen/logrus"
)

func getIdFromParam(c *gin.Context, log *logrus.Entry) (int64, bool) {
	rawId, ok := c.Params.Get("id")
	if !ok {
		log.Error("Missing Id")
		helpers.GinError(c, http.StatusBadRequest, "missing id")
		return 0, true
	}

	id, err := utils.ParseInt64(rawId)
	if err != nil {
		log.WithError(err).WithField("id", id).Error("Invalid id")
		helpers.GinError(c, http.StatusBadRequest, "id must be a valid integer")
		return 0, true
	}
	return id, false
}
