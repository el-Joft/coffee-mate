package exception_test

import (
	"coffee-mate/src/middleware/exception"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalServerError(t *testing.T) {
	var errors []map[string]interface{}
	defer func() {
		if err := recover(); err != nil {
			assert.NotEmpty(t, err)
		}
	}()
	exception.InternalServerError("ISE Testing", errors)
}
