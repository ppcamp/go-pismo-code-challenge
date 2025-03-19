package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIdKey = "request_id"

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqid := uuid.NewString()

		ctx.Set(RequestIdKey, reqid)
	}
}
