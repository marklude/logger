package logger

import (
	"time"

	"github.com/rotisserie/eris"
)

type Logger struct {
	env string
}

var dt time.Time

func (Logger) Info(err error, args ...interface{}) {
	dt = time.Now()
	eris.Wrapf(err, "[INFO] %v %v\n", append([]interface{}{dt.Format("02-01-2006 15:04:05")}, args...)...)
}

func (Logger) Warn(err error, args ...interface{}) {
	dt = time.Now()
	eris.Wrapf(err, "[WARN] %v %v\n", append([]interface{}{dt.Format("02-01-2006 15:04:05")}, args...)...)
}

func (Logger) Error(err error, args ...interface{}) {
	dt = time.Now()
	eris.Wrapf(err, "[ERROR] %v %v\n", append([]interface{}{dt.Format("02-01-2006 15:04:05")}, args...)...)
}

func NewLogger(env string) *Logger {
	return &Logger{env: env}
}
