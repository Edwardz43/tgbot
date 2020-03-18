package zaplogger_test

import (
	"Edwardz43/tgbot/log/zaplogger"
	"testing"
)

func TestLogWithInfoLevel(t *testing.T) {
	logger := zaplogger.GetInstance()
	logger.INFO("test info log")
	t.Log("done")
}

func TestLogWithErrorLevel(t *testing.T) {
	logger := zaplogger.GetInstance()
	logger.ERROR("test error log")
	t.Log("done")
}

func TestLogWithPanicLevel(t *testing.T) {
	logger := zaplogger.GetInstance()
	logger.PANIC("test panic log")
	t.Log("done")
}

func TestLogWithFatalLevel(t *testing.T) {
	logger := zaplogger.GetInstance()
	logger.FATAL("test panic log")
	t.Log("done")
}
