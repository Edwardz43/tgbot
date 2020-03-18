package zaplogger

import (
	"Edwardz43/tgbot/config"
	"Edwardz43/tgbot/log"
	"fmt"
	"runtime"
	"strconv"

	"go.uber.org/zap/zapcore"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// GetInstance return a logger instance with options
func GetInstance(args ...interface{}) log.Logger {

	// init
	l := &Logger{}

	all := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	esEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	w, _, err := zap.Open("foo.log")

	if err != nil {
		fmt.Println(err)
	}

	core := zapcore.NewTee(zapcore.RegisterHooks(zapcore.NewCore(esEncoder, w, all), l.Hook))

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync()

	l.Zap = logger

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

// Hook implements the hook function of zap logger hook
func (l *Logger) Hook(e zapcore.Entry) error {
	if l.IsHook {
		err := log.Emit(&log.Content{
			Level:   e.Level.CapitalString(),
			Message: e.Message,
			Date:    e.Time,
			Caller:  e.Caller.TrimmedPath(),
			Stack:   e.Stack,
		})
		return err
	}
	return nil
}

// INFO create a info level log
func (l *Logger) INFO(msg string) {
	defer l.Zap.Sync()
	sugar := l.Zap.Sugar()
	sugar.Infow(msg)
}

// DEBUG create a debug level log
func (l *Logger) DEBUG(msg string) {
	defer l.Zap.Sync()
	sugar := l.Zap.Sugar()
	sugar.Debugw(msg)
}

// WARN create a warn level log
func (l *Logger) WARN(msg string) {
	defer l.Zap.Sync()
	sugar := l.Zap.Sugar()
	sugar.Warnw(msg)
}

// ERROR create a error level log
func (l *Logger) ERROR(msg string) {
	defer l.Zap.Sync()
	sugar := l.Zap.Sugar()
	sugar.Errorw(msg)
}

// PANIC create a panic level log
func (l *Logger) PANIC(msg string) {
	defer l.Zap.Sync()
	sugar := l.Zap.Sugar()
	sugar.Panicw(msg)
}

// FATAL create a fatal level log
func (l *Logger) FATAL(msg string) {
	defer l.Zap.Sync()
	sugar := l.Zap.Sugar()
	sugar.Fatalw(msg)
}
