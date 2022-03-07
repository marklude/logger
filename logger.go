package logger

import (
	"fmt"
	"time"

	"github.com/rotisserie/eris"
)

type Logger struct {
	env string
}

var dt time.Time

func (Logger) Info(args ...interface{}) {
	dt = time.Now()
	fmt.Printf("[INFO] %v %v\n", append([]interface{}{dt.Format("02-01-2006 15:04:05")}, args...)...)
}

func (Logger) Warn(args ...interface{}) {
	dt = time.Now()
	fmt.Printf("[WARN] %v %v\n", append([]interface{}{dt.Format("02-01-2006 15:04:05")}, args...)...)
}

func (Logger) Error(err error, args ...interface{}) {
	dt = time.Now()
	eris.Wrapf(err, "[ERROR] %v %v\n")
}

func NewLogger(env string) *Logger {
	return &Logger{env: env}
}
