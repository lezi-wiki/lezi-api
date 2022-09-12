package log

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	asserts := assert.New(t)

	logger := NewLogger()
	asserts.NotNil(logger)
	asserts.Equal(logrus.DebugLevel, logger.Level)

	asserts.NotPanics(func() {
		logger.SetOutput(os.Stdout)
	})

	asserts.NotPanics(func() {
		logger.Debugf("debugf %s", "tests")
		logger.Infof("infof %s", "tests")
		logger.Errorf("errorf %s", "tests")
	})
	asserts.Panics(func() {
		logger.Panicf("panicf %s", "tests")
	})
	asserts.Panics(func() {
		logger.Fatalf("fatalf %s", "tests")
	})
}
