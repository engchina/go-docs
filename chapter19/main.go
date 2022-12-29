package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LogHook struct {
	ErrorLogFile string
}

func (lh LogHook) Levels() []logrus.Level {
	//example-1
	//return logrus.AllLevels
	return []logrus.Level{logrus.ErrorLevel}
}
func (lh LogHook) Fire(entry *logrus.Entry) error {
	file, _ := os.OpenFile(lh.ErrorLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	line, _ := entry.String()
	file.Write([]byte(line))
	return nil
}

func main() {
	logrus.AddHook(&LogHook{ErrorLogFile: "error.log"})

	logrus.Warnln("hello")
	logrus.Errorln("hello")
}
