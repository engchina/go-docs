package main

import (
	"github.com/sirupsen/logrus"
)

const (
	cBlack  = 0
	cRed    = 1
	cGreen  = 2
	cYellow = 3
	cBlue   = 4
	cCyan   = 5
	cPurple = 6
	cGray   = 7
)

func PrintColor(colorCode int, text string, isBackGround bool) {
	if isBackGround {
		logrus.Infof("\033[4%dm %s \033[0m", colorCode, text)
	} else {
		logrus.Infof("\033[3%dm %s \033[0m", colorCode, text)
	}
}

func main() {
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	logrus.Infoln(logrus.GetLevel())
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Println("learn logrus")
	logrus.Debugln("learn logrus")
	logrus.Infoln("learn logrus")
	logrus.Warnln("learn logrus")
	logrus.Errorln("learn logrus")

	log := logrus.WithField("key1", "value1").WithField("key2", "value2").WithFields(logrus.Fields{
		"key3": "value3",
		"key4": "value4",
	})

	log.Infof("hello")

	log.Infof("\033[3%dm hello \033[0m", cBlack)
	log.Infof("\033[3%dm hello \033[0m", cRed)
	log.Infof("\033[3%dm hello \033[0m", cGreen)
	log.Infof("\033[3%dm hello \033[0m", cYellow)
	log.Infof("\033[3%dm hello \033[0m", cBlue)
	log.Infof("\033[3%dm hello \033[0m", cCyan)
	log.Infof("\033[3%dm hello \033[0m", cPurple)
	log.Infof("\033[3%dm hello \033[0m", cGray)

	log.Infof("\033[4%dm hello \033[0m", cBlack)
	log.Infof("\033[4%dm hello \033[0m", cRed)
	log.Infof("\033[4%dm hello \033[0m", cGreen)
	log.Infof("\033[4%dm hello \033[0m", cYellow)
	log.Infof("\033[4%dm hello \033[0m", cBlue)
	log.Infof("\033[4%dm hello \033[0m", cCyan)
	log.Infof("\033[4%dm hello \033[0m", cPurple)
	log.Infof("\033[4%dm hello \033[0m", cGray)

	PrintColor(1, "hello world", true)
}
