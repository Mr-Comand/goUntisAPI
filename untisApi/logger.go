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

// Debug level log with blue color
func (l *logger) Debug(msg ...string) {
	if l.logger == nil {
		return
	}
	if l.level <= DEBUG {
		l.logger.SetPrefix(fmt.Sprintf("%s[DEBUG]\t%s", Blue, Reset))
		l.logger.Println(strings.Join(msg, " "))
	}
}

// Info level log with green color
func (l *logger) Info(msg ...string) {
	if l.logger == nil {
		return
	}
	if l.level <= INFO {
		l.logger.SetPrefix(fmt.Sprintf("%s[INFO]\t%s", Green, Reset))
		l.logger.Println(strings.Join(msg, " "))
	}
}

// Warn level log with yellow color
func (l *logger) Warn(msg ...string) {
	if l.logger == nil {
		return
	}
	if l.level <= WARN {
		l.logger.SetPrefix(fmt.Sprintf("%s[WARN]\t%s", Yellow, Reset))
		l.logger.Println(strings.Join(msg, " "))
	}
}

// Error level log with red color
func (l *logger) Error(msg ...string) {
	if l.logger == nil {
		return
	}
	if l.level <= ERROR {
		l.logger.SetPrefix(fmt.Sprintf("%s[ERROR]\t%s", Red, Reset))
		l.logger.Println(strings.Join(msg, " "))
	}
}
