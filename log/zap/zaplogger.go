package zaplogger

import (
	"Edwardz43/tgbot/log"

	"go.uber.org/zap"
)

// GetInstance return a logger instance with options
func GetInstance(args ...interface{}) *Logger {
	l := &Logger{
		Zap: zap.NewExample(),
	}
	return l
}

// Logger implements logger by uber zap
type Logger struct {
	Zap *zap.Logger
}

func (l *Logger) hook() error {

	return nil
}

// INFO create a info level log
func (l *Logger) INFO(log *log.Content) {
	defer l.Zap.Sync() // flushes buffer, if any
	sugar := l.Zap.Sugar()
	sugar.Infow(log.Message, "date", &log.Date, "line", log.Line)
}

// DEBUG create a debug level log
func (l *Logger) DEBUG(log *log.Content) {
	defer l.Zap.Sync() // flushes buffer, if any
	sugar := l.Zap.Sugar()
	sugar.Debugw(log.Message, "date", &log.Date, "line", log.Line)
}

// WARN create a warn level log
func (l *Logger) WARN(log *log.Content) {
	defer l.Zap.Sync() // flushes buffer, if any
	sugar := l.Zap.Sugar()
	sugar.Warnw(log.Message, "date", &log.Date, "line", log.Line)
}

// ERROR create a error level log
func (l *Logger) ERROR(log *log.Content) {
	defer l.Zap.Sync() // flushes buffer, if any
	sugar := l.Zap.Sugar()
	sugar.Errorw(log.Message, "date", &log.Date, "line", log.Line)
}

// PANIC create a panic level log
func (l *Logger) PANIC(log *log.Content) {
	defer l.Zap.Sync() // flushes buffer, if any
	sugar := l.Zap.Sugar()
	sugar.Panicw(log.Message, "date", &log.Date, "line", log.Line)
}

// FATAL create a fatal level log
func (l *Logger) FATAL(log *log.Content) {
	defer l.Zap.Sync() // flushes buffer, if any
	sugar := l.Zap.Sugar()
	sugar.Fatalw(log.Message, "date", &log.Date, "line", log.Line)
}
