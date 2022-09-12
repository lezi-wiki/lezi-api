package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogger(t *testing.T) {
	asserts := assert.New(t)

	{
		logger := Log()
		asserts.NotNil(logger)
		asserts.Equal(logger, GlobalLogger)
	}

	{
		GlobalLogger = nil
		logger := Log()
		asserts.NotNil(logger)
		asserts.Equal(logger, GlobalLogger)
	}
}

func BenchmarkLogger(b *testing.B) {
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Log().Printf("%s", "tests")
	}

	b.StopTimer()
}

func BenchmarkGormLogger_Printf(b *testing.B) {
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		new(GormLogger).Printf("%s", "tests")
	}

	b.StopTimer()
}
