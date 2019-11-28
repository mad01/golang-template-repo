package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = &logrus.Entry{}

func initLog(debug bool) {
	logger = newLogger(debug)
}

func newLogger(debug bool) *logrus.Entry {
	l := logrus.Logger{
		Out:       os.Stdout,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}

	if debug == true {
		l.SetLevel(logrus.DebugLevel)
	}

	entry := logrus.NewEntry(&l)
	return entry
}

func main() {
	initLog(false)
	runCmd()
}
