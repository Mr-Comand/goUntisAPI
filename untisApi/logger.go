package untisApi

import (
	"fmt"
	"log"
	"strings"
)

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	NONE
)

type LogLevel int

const (
	// ANSI escape codes for colors
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// logger structure
type logger struct {
	level  LogLevel
	logger *log.Logger
}

// New logger constructor
func newLogger(logLogger *log.Logger, level LogLevel) *logger {
	return &logger{
		level:  level,
		logger: logLogger,
	}
}

// Helper function to log messages with color and level
func (l *logger) logWithLevel(color, level string, msg ...string) {
	if l.logger == nil {
		return
	}

	// Save the original prefix
	originalPrefix := l.logger.Prefix()

	// Set new prefix with color and level
	l.logger.SetPrefix(fmt.Sprintf("%s[%s]\t%s", color, level, Reset))

	// Log the message
	l.logger.Println(strings.Join(msg, " "))

	// Restore the original prefix
	l.logger.SetPrefix(originalPrefix)
}

// Debug level log with blue color
func (l *logger) Debug(msg ...string) {
	if l.level <= DEBUG {
		l.logWithLevel(Blue, "DEBUG", msg...)
	}
}

// Info level log with green color
func (l *logger) Info(msg ...string) {
	if l.level <= INFO {
		l.logWithLevel(Green, "INFO", msg...)
	}
}

// Warn level log with yellow color
func (l *logger) Warn(msg ...string) {
	if l.level <= WARN {
		l.logWithLevel(Yellow, "WARN", msg...)
	}
}

// Error level log with red color
func (l *logger) Error(msg ...string) {
	if l.level <= ERROR {
		l.logWithLevel(Red, "ERROR", msg...)
	}
}
