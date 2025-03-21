package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-pismo-code-challenge/internal/services"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/dtos"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/utils/helpers"
	"github.com/sirupsen/logrus"
)

type accountHandler struct{ svc services.Account }

func NewAccountHandler(h *Handler) *accountHandler {
	return &accountHandler{h.Account}
}

func (s *accountHandler) Create(c *gin.Context) {
	log := logrus.WithContext(c)

	var input dtos.CreateAccount

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

func (s *accountHandler) Get(c *gin.Context) {
	log := logrus.WithContext(c)

	id, shouldReturn := getIdFromParam(c, log)
	if shouldReturn {
		return
	}

	v, err := s.svc.Get(c, id)
	if err == nil {
		c.JSON(http.StatusOK, v)
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		helpers.GinError(c, http.StatusNotFound, "not found any item for this id")
	} else {
		helpers.GinError(c, http.StatusInternalServerError, "some unexpected error")
	}
}

func (s *accountHandler) GetLimits(c *gin.Context) {
	log := logrus.WithContext(c)

	id, shouldReturn := getIdFromParam(c, log)
	if shouldReturn {
		return
	}

	v, err := s.svc.GetAccountLimits(c, id)
	if err == nil {
		c.JSON(http.StatusOK, v)
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		helpers.GinError(c, http.StatusNotFound, "not found any item for this id")
	} else {
		helpers.GinError(c, http.StatusInternalServerError, "some unexpected error")
	}
}

func (s *accountHandler) SetLimit(c *gin.Context) {
	log := logrus.WithContext(c)

	var input dtos.ChangeAccountLimit

	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.WithError(err).Error("fail to bind json")
		helpers.GinError(c, http.StatusBadRequest, fmt.Sprintf("fail to load bind json: %v", err))
		return
	}

	id, shouldReturn := getIdFromParam(c, log)
	if shouldReturn {
		return
	}

	err = s.svc.SetLimit(c, id, input.NewLimit)
	if err == nil {
		c.JSON(http.StatusCreated, "")
		return
	}

	log.WithError(err).Error("some unexpected error occurred")
	helpers.GinError(c, http.StatusInternalServerError, "some unexpected error")
}
