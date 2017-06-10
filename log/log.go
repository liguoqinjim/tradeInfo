package log

import (
	logrus "github.com/sirupsen/logrus"
)

func init() {
	//file, err := os.Create("tradeInfo.log")
	//if err != nil {
	//	panic(err)
	//}
	//logrus.SetOutput(file)
}

func Debugf(format string, a ...interface{}) {
	logrus.Debugf(format, a)
}

func Infof(format string, a ...interface{}) {
	logrus.Infof(format, a)
}

func Warnf(format string, a ...interface{}) {
	logrus.Warnf(format, a)
}

func Errorf(format string, a ...interface{}) {
	logrus.Errorf(format, a)
}

func Fatalf(format string, a ...interface{}) {
	logrus.Fatalf(format, a)
}

func Panicf(format string, a ...interface{}) {
	logrus.Panicf(format, a)
}

func Debug(message string) {
	logrus.Debug(message)
}
func Info(message string) {
	logrus.Info(message)
}
func Warn(message string) {
	logrus.Warn(message)
}
func Error(message string) {
	logrus.Error(message)
}
func Fatal(message string) {
	logrus.Fatal(message)
}
func Panic(message string) {
	logrus.Panic(message)
}
