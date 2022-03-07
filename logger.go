package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	env string
}

var dt time.Time

func (Logger) Info(args ...interface{}) {
	dt = time.Now()
	fmt.Printf("[INFO] %s \n", append([]interface{}{dt.Format("02-01-2006 15:04:05")}, args...)...)
}

func (Logger) Warn(args ...interface{}) {
	dt = time.Now()
	fmt.Printf("[WARN] %s \n", append([]interface{}{dt.Format("02-01-2006 15:04:05")}, args...)...)
}

func (Logger) Error(args ...interface{}) {
	dt = time.Now()
	fmt.Printf("[ERROR] %s \n", append([]interface{}{dt.Format("02-01-2006 15:04:05")}, args...)...)
}

func NewLogger(env string) *Logger {
	return &Logger{env: env}
}
