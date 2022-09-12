package log

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type formatter struct {
	pid string
}

// 日志颜色
var colors = map[logrus.Level]func(format string, a ...interface{}) string{
	logrus.WarnLevel:  color.New(color.FgYellow).Add(color.Bold).SprintfFunc(),
	logrus.PanicLevel: color.New(color.BgHiRed).Add(color.Bold).SprintfFunc(),
	logrus.FatalLevel: color.New(color.BgRed).Add(color.Bold).SprintfFunc(),
	logrus.ErrorLevel: color.New(color.FgRed).Add(color.Bold).SprintfFunc(),
	logrus.InfoLevel:  color.New(color.FgCyan).Add(color.Bold).SprintfFunc(),
	logrus.DebugLevel: color.New(color.FgWhite).Add(color.Bold).SprintfFunc(),
}

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	colorFunc := colors[entry.Level]

	level := entry.Level.String()
	if entry.Data["level"] != nil {
		level = entry.Data["level"].(string)
	}

	switch entry.Logger.Level {
	case logrus.DebugLevel:
		level = colorFunc("%-11s", "["+strings.ToUpper(level)+"]")
	default:
		level = colorFunc("%-7s", "["+strings.ToUpper(level)+"]")
	}

	return []byte(fmt.Sprintf(
		"%s %s | %s | %s\n",
		level,
		f.pid,
		entry.Time.Format("2006-01-02 15:04:05.000"),
		entry.Message,
	)), nil
}

func NewFormatter() *formatter {
	return &formatter{
		pid: color.New(color.FgHiMagenta).Sprint(os.Getpid()),
	}
}
