package db

import (
	"fmt"

	errutils "github.com/ppcamp/go-pismo-code-challenge/pkg/utils/errors"
)

var (
	ErrDriverError = fmt.Errorf("driver error")
	ErrNotFound    = fmt.Errorf("%w: %w", ErrDriverError, errutils.ErrNotFound)
)
