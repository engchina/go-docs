package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
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

type LogFormatter struct {
	Prefix     string
	TimeFormat string
}

func (lf LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var color int
	switch entry.Level {
	case logrus.ErrorLevel:
		color = cRed
	case logrus.WarnLevel:
		color = cYellow
	case logrus.InfoLevel:
		color = cBlue
	case logrus.DebugLevel:
		color = cCyan
	default:
		color = cGray
	}

	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}

	formatTime := entry.Time.Format(lf.TimeFormat)
	fileLine := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	fmt.Fprintf(b, "[%s] \033[3%dm[%s]\033[0m [%s] %s %s\n", lf.Prefix, color, entry.Level, formatTime, fileLine, entry.Message)

	return b.Bytes(), nil
}

func main() {
	//example-1
	//logfile, _ := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//writes := []io.Writer{
	//	logfile,
	//	os.Stdout,
	//}
	//logrus.SetOutput(io.MultiWriter(writes...))
	//logrus.Infoln("hello")

	//example-2
	//logfile, _ := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//logrus.SetOutput(io.MultiWriter(logfile, os.Stdout))
	//logrus.Infoln("hello")

	//example-3
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{Prefix: "ORACLE", TimeFormat: "2006-01-02 15:04:06"})
	logrus.Infof("hello")
}
