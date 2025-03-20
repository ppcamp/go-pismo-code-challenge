package utils_test

import (
	"testing"

	"github.com/ppcamp/go-pismo-code-challenge/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseInt64(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		TestName    string
		Input       string
		ExpectError bool
	}{
		{"should not work", "124f", true},
		{"should work", "124", false},
	}

	for _, tt := range cases {
		_, err := utils.ParseInt64(tt.Input)
		if tt.ExpectError {
			assert.Error(err)
		} else {
			assert.NoError(err)
		}
	}
}
