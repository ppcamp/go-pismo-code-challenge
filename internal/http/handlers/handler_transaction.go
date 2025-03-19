package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-pismo-code-challenge/internal/services"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/dtos"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/utils/helpers"
	"github.com/sirupsen/logrus"
)

type transactionHandler struct{ svc services.Transaction }

func NewTransactionHandler(h *Handler) *transactionHandler {
	return &transactionHandler{h.Transaction}
}

func (s *transactionHandler) Create(c *gin.Context) {
	log := logrus.WithContext(c)

	var input dtos.CreateTransaction

	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.WithError(err).Error("fail to bind json")
		helpers.GinError(c, http.StatusBadRequest, fmt.Sprintf("fail to load bind json: %v", err))
		return
	}

	err = s.svc.Create(c, &input)
	if err == nil {
		c.JSON(http.StatusCreated, "")
		return
	}

	log.WithError(err).Error("some unexpected error occurred")
	helpers.GinError(c, http.StatusInternalServerError, "some unexpected error")
}
