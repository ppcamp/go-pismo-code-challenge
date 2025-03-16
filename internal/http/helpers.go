package http

import (
	"context"

	"github.com/gin-gonic/gin"
)

type FuncErr[Arg any] func(ctx context.Context, arg Arg) error
type Func[Arg any, Res any] func(ctx context.Context, arg Arg) (Res, error)

type GinFunc func(c *gin.Context)
