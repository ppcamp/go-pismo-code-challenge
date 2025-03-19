package requests

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type reqIdStruct struct{}

var reqIdKey = reqIdStruct{}

func NewContextWithRequestId(ctx context.Context) context.Context {
	reqid := uuid.NewString()
	return context.WithValue(ctx, reqIdKey, reqid)
}

func RequestIdFromContext(ctx context.Context) (string, error) {
	v, ok := ctx.Value(reqIdKey).(string)
	if !ok {
		return "", errors.New("request id not found in context")
	}

	return v, nil
}
