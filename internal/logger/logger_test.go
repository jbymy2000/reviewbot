package logger

import "testing"

func Test_1(t *testing.T) {
	Info("abc")
	Infof("abc%s , %s ", "abc", "bcd")
	Error("abc")
	Errorf("abc%s , %s ", "abc", "bcd")

}
