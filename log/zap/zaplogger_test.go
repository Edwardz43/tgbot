package zaplogger_test

import (
	"Edwardz43/tgbot/log"
	zaplogger "Edwardz43/tgbot/log/zap"
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
	c := &log.Content{
		Message: "test message",
		Date:    time.Now(),
		Line:    123,
	}
	logger.INFO(c)
	t.Log("done")
}
