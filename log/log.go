package log

// Logger defines the log behaviors
type Logger interface {
	INFO(string)
	DEBUG(string)
	WARN(string)
	ERROR(string)
	FATAL(string)
	PANIC(string)
	//hook() error
}
