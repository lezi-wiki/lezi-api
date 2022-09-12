package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGormLogger_Printf(t *testing.T) {
	asserts := assert.New(t)

	logger := new(GormLogger)
	asserts.NotNil(logger)
	asserts.NotPanics(func() {
		logger.Printf("test %s", "tests")
	})
}
