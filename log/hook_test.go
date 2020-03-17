package log_test

import (
	"Edwardz43/tgbot/log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmitThenReturnSuccess(t *testing.T) {
	c := &log.Content{
		Message:  "Test ES Log",
		Date:     time.Now(),
		Line:     123,
		FileName: "hook_test.go",
		Function: "TestEmitThenReturnSuccess",
	}
	err := log.Emit(c)
	assert.Nil(t, err)
}
