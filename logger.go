package systemgo

import (
	"fmt"
)

// logger

func logger(level string, msg string, args...interface {}) {
	fmt.Printf("[SG %d %s] %s\n", CurrentTime(), level, fmt.Sprintf(msg, args...))
}

func Info(msg string, args...interface {}) {
	logger("I", msg, args...)
}

func Err(msg string, args...interface {}) {
	logger("E", msg, args...)
}

func Debug(msg string, args...interface {}) {
	logger("D", msg, args...)
}

func Warn(msg string, args...interface {}) {
	logger("W", msg, args...)
}

