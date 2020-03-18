package zaplogger

import (
	"Edwardz43/tgbot/config"
	"Edwardz43/tgbot/log"
	"strconv"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// GetInstance as the factory creating a logger instance with options
func GetInstance() log.Logger {

	// init
	logger := &Logger{}

	all := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return true })

	esEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	logrotateWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.GetLogPath(),
		MaxSize:    20, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	core := zapcore.NewTee(zapcore.RegisterHooks(zapcore.NewCore(esEncoder, logrotateWriter, all), logger.Hook))

	logger.Zap = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	defer logger.Zap.Sync()

	if isHooked, err := strconv.ParseBool(config.GetLogHook()); err != nil {
		errors.Errorf("zaplogger GetInstance failed")
	} else {
		logger.IsHook = isHooked
	}

	return logger
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
	l.Zap.Info(msg)
}

// DEBUG create a debug level log
func (l *Logger) DEBUG(msg string) {
	defer l.Zap.Sync()
	l.Zap.Debug(msg)
}

// WARN create a warn level log
func (l *Logger) WARN(msg string) {
	defer l.Zap.Sync()
	l.Zap.Warn(msg)
}

// ERROR create a error level log
func (l *Logger) ERROR(msg string) {
	defer l.Zap.Sync()
	l.Zap.Error(msg)
}

// PANIC create a panic level log
func (l *Logger) PANIC(msg string) {
	defer l.Zap.Sync()
	l.Zap.DPanic(msg)
}

// FATAL create a fatal level log
func (l *Logger) FATAL(msg string) {
	defer l.Zap.Sync()
	l.Zap.Fatal(msg)
}
