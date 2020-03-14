package zaplogger

import (
	"go.uber.org/zap"
)

// GetInstance return a logger instance with options
func GetInstance(args ...interface{}) *Logger {
	l := &Logger{
		zap: zap.NewExample(),
	}
	return l
}

// Logger implements logger by uber zap
type Logger struct {
	zap *zap.Logger
}

// INFO create a info level log
func (l *Logger) INFO(args ...interface{}) {
	defer l.zap.Sync() // flushes buffer, if any
	sugar := l.zap.Sugar()
	sugar.Infow(args[0].(string), args[1:])
}

// DEBUG create a debug level log
func (l *Logger) DEBUG(args ...interface{}) {
	defer l.zap.Sync() // flushes buffer, if any
	sugar := l.zap.Sugar()
	sugar.Debugw(args[0].(string), args[1:])
}

// WARN create a warn level log
func (l *Logger) WARN(args ...interface{}) {
	defer l.zap.Sync() // flushes buffer, if any
	sugar := l.zap.Sugar()
	sugar.Warnw(args[0].(string), args[1:])
}

// ERROR create a error level log
func (l *Logger) ERROR(args ...interface{}) {
	defer l.zap.Sync() // flushes buffer, if any
	sugar := l.zap.Sugar()
	sugar.Errorw(args[0].(string), args[1:])
}

// PANIC create a panic level log
func (l *Logger) PANIC(args ...interface{}) {
	defer l.zap.Sync() // flushes buffer, if any
	sugar := l.zap.Sugar()
	sugar.Panicw(args[0].(string), args[1:])
}

// FATAL create a fatal level log
func (l *Logger) FATAL(args ...interface{}) {
	defer l.zap.Sync() // flushes buffer, if any
	sugar := l.zap.Sugar()
	sugar.Fatalw(args[0].(string), args[1:])
}
