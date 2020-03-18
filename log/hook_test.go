package log_test

import (
	"Edwardz43/tgbot/log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmitThenReturnSuccess(t *testing.T) {
	c := &log.Content{
		Level:   "Info",
		Message: "Test ES Log",
		Date:    time.Now(),
		Caller:  "zaplogger/zaplogger.go:104",
	}
	err := log.Emit(c)
	assert.Nil(t, err)
}
