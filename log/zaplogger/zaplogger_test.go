package zaplogger_test

import (
	"Edwardz43/tgbot/log/zaplogger"
	"testing"
	"time"
)

type test struct {
	v1 string
	v2 bool
	v3 int
	v4 time.Time
}

func TestLog(t *testing.T) {
	logger := zaplogger.GetInstance()
	logger.INFO("test message")
	t.Log("done")
}
