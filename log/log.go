package log

import "time"

// Logger defines the log behaviors
type Logger interface {
	Debug(Content)
	Info(Content)
	Warn(Content)
	Error(Content)
	Fatal(Content)
	Panic(Content)
	hook() error
}

// Content is a log content data model
type Content struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
	Line    int       `json:"line"`
}
