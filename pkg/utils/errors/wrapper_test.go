package errors_test

import (
	"database/sql"
	"errors"
	"testing"

	errwrapper "github.com/ppcamp/go-pismo-code-challenge/pkg/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestErrorWrapper(t *testing.T) {
	assert := assert.New(t)

	err := errwrapper.Error{
		Base:    errors.New("example err"),
		Wrapped: sql.ErrNoRows,
	}

	assert.ErrorContains(err, "example")
}
