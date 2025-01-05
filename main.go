package ObjectNimesislogger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

const (
	// log levels
	INFO  = "INFO"
	ERROR = "ERROR"
	WARN  = "WARN"
	DEBUG = "DEBUG"
	TRACE = "TRACE"
	FATAL = "FATAL"
)

// logger structure
type logger struct {
	scope string
}

// New creates a new instance of logger with a given prefix
func New(scope string) *logger {
	return &logger{scope: scope}
}

// Scope creates a new logger instance with a specific scope
func (l *logger) Scope(scope string) *logger {
	// Create a new instance of logger with the provided scope
	return &logger{scope: scope}
}

// log formats and prints messages based on the log level, time, and scope.
func (l *logger) log(level, msg, scope string) {

	currentTime := time.Now().Format("15:04:05") // Time without date
	levelColor := getLogLevelColor(level)
	bold := color.New(color.Bold)

	// Format the log message
	message := fmt.Sprintf("[%s] %s (%s): %s",
		color.New(levelColor).Sprintf(bold.Sprint(level)),
		bold.Sprint(currentTime),
		scope,
		msg)

	fmt.Println(message)

}

// Info logs an info message
func (l *logger) Info(msg string) {
	l.log(INFO, msg, l.scope)
}

// Trace logs an info message
func (l *logger) Trace(msg string) {
	l.log(TRACE, msg, l.scope)
}

// Error logs an error message
func (l *logger) Error(msg string) {
	l.log(ERROR, msg, l.scope)
}

// Fatal logs an info message
func (l *logger) Fatal(msg string) {
	l.log(FATAL, msg, l.scope)
}

// Warn logs a warning message
func (l *logger) Warn(msg string) {
	l.log(WARN, msg, l.scope)
}

// Debug logs a debug message
func (l *logger) Debug(msg string) {
	l.log(DEBUG, msg, l.scope)
}

func getLogLevelColor(level string) color.Attribute {
	switch level {
	case INFO:
		return color.FgGreen
	case ERROR:
		return color.FgRed
	case WARN:
		return color.FgYellow
	case DEBUG:
		return color.FgCyan
	case TRACE:
		return color.FgMagenta
	case FATAL:
		return color.FgHiRed
	default:
		return color.FgWhite
	}
}
