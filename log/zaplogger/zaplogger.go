package zaplogger

import (
	"Edwardz43/tgbot/config"
	"Edwardz43/tgbot/log"
	"runtime"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// GetInstance return a logger instance with options
func GetInstance(args ...interface{}) log.Logger {
	l := &Logger{
		Zap: zap.NewExample(),
	}

	isHooked, err := strconv.ParseBool(config.GetLogHook())

	if err != nil {
		errors.Errorf("zaplogger GetInstance failed")
	}

	l.IsHook = isHooked

	return l
}

func trace() (string, int, string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frames.Next()
	frame, _ := frames.Next()
	return frame.File, frame.Line, frame.Function
}

// Logger implements logger by uber zap
type Logger struct {
	Zap    *zap.Logger
	IsHook bool
}

func (l *Logger) hook(msg string, date time.Time, line int, file string, function string) {
	if l.IsHook {
		log.Emit(&log.Content{
			Message:  msg,
			Date:     date,
			Line:     line,
			FileName: file,
			Function: function,
		})
	}
}

// INFO create a info level log
func (l *Logger) INFO(msg string) {
	file, line, function := trace()
	defer l.Zap.Sync()
	sugar := l.Zap.Sugar()
	sugar.Infow(msg, "date", time.Now(), "line", line, "file_name", file, "func", function)
	l.hook(msg, time.Now(), line, file, function)
}

// DEBUG create a debug level log
func (l *Logger) DEBUG(msg string) {
	file, line, function := trace()
	defer l.Zap.Sync()
	sugar := l.Zap.Sugar()
	sugar.Debugw(msg, "date", time.Now, "line", line, "file_name", file, "func", function)
	l.hook(msg, time.Now(), line, file, function)
}

// WARN create a warn level log
func (l *Logger) WARN(msg string) {
	file, line, function := trace()
	defer l.Zap.Sync() // flushes buffer, if any
	sugar := l.Zap.Sugar()
	sugar.Warnw(msg, "date", time.Now, "line", line, "file_name", file, "func", function)
	l.hook(msg, time.Now(), line, file, function)
}

// ERROR create a error level log
func (l *Logger) ERROR(msg string) {
	file, line, function := trace()
	defer l.Zap.Sync() // flushes buffer, if any
	sugar := l.Zap.Sugar()
	sugar.Errorw(msg, "date", time.Now, "line", line, "file_name", file, "func", function)
	l.hook(msg, time.Now(), line, file, function)
}

// PANIC create a panic level log
func (l *Logger) PANIC(msg string) {
	file, line, function := trace()
	defer l.Zap.Sync() // flushes buffer, if any
	sugar := l.Zap.Sugar()
	sugar.Panicw(msg, "date", time.Now, "line", line, "file_name", file, "func", function)
	l.hook(msg, time.Now(), line, file, function)
}

// FATAL create a fatal level log
func (l *Logger) FATAL(msg string) {
	file, line, function := trace()
	defer l.Zap.Sync() // flushes buffer, if any
	sugar := l.Zap.Sugar()
	sugar.Fatalw(msg, "date", time.Now, "line", line, "file_name", file, "func", function)
	l.hook(msg, time.Now(), line, file, function)
}
