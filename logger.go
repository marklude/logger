package logger

import (
	"fmt"
	"time"
)

type logger struct {
	env string
}

var dt time.Time

func Init() {
	dt = time.Now()
}

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

func (*logger) Info(args ...interface{}) {
	fmt.Printf("[INFO] %s %s", append([]interface{}{dt.Format("01-02-2006 15:04:05")}, args...)...)
}

func (*logger) Warn(args ...interface{}) {
	fmt.Printf("[WARN] %s %s", append([]interface{}{dt.Format("01-02-2006 15:04:05")}, args...)...)
}

func (*logger) Error(args ...interface{}) {
	fmt.Printf("[ERROR] %s %s", append([]interface{}{dt.Format("01-02-2006 15:04:05")}, args...)...)
}

func NewLogger(env string) Logger {
	return &logger{env: env}
}
