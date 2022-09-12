package log

import (
	"github.com/sirupsen/logrus"
)

var GlobalLogger *logrus.Logger
var Level = logrus.InfoLevel

func Log() *logrus.Logger {
	if GlobalLogger == nil {
		GlobalLogger = NewLogger()
	}

	return GlobalLogger
}
