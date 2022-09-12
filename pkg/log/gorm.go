package log

type GormLogger struct{}

func (g *GormLogger) Printf(format string, args ...interface{}) {
	Log().WithField("level", "Database").Debugf(format, args...)
}
