package logger

import (
	"fmt"
	"log"
)

var LoggerInfo *log.Logger
var LoggerError *log.Logger

func init() {
	LoggerInfo = log.New(log.Writer(), "[reviewBot][INFO] ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	LoggerError = log.New(log.Writer(), "[reviewBot][ERROR] ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)

}

func Info(s string) {
	_ = LoggerInfo.Output(2, s)
}

func Infof(f string, s ...interface{}) {
	_ = LoggerInfo.Output(2, fmt.Sprintf(f, s...))
}

func Error(s string) {
	_ = LoggerError.Output(2, s)
}

func Errorf(f string, s ...interface{}) {
	_ = LoggerError.Output(2, fmt.Sprintf(f, s...))
}
