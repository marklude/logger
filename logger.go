package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	env string
}

var dt time.Time

func Init() {
	dt = time.Now()
}

func (Logger) Info(args ...interface{}) {
	fmt.Printf("[INFO] %s %s", append([]interface{}{dt.Format("01-02-2006 15:04:05")}, args...)...)
}

func (Logger) Warn(args ...interface{}) {
	fmt.Printf("[WARN] %s %s", append([]interface{}{dt.Format("01-02-2006 15:04:05")}, args...)...)
}

func (Logger) Error(args ...interface{}) {
	fmt.Printf("[ERROR] %s %s", append([]interface{}{dt.Format("01-02-2006 15:04:05")}, args...)...)
}

func NewLogger(env string) *Logger {
	return &Logger{env: env}
}
