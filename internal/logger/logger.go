package logger

import (
	"fmt"
	"log"
)

var Logger *log.Logger

func init() {
	Logger = log.New(log.Writer(), "[MyApp] ", log.LstdFlags|log.Lshortfile)
}

func Info(s string) {
	Logger.Println("[INFO]" + s)
}

func Infof(f string, s ...interface{}) {
	Logger.Println(fmt.Sprintf("[INFO]"+f, s...))
}

func Error(s string) {
	Logger.Println("[ERROR]" + s)
}

func Errorf(f string, s ...interface{}) {
	Logger.Println(fmt.Sprintf("[ERROR]"+f, s...))
}
