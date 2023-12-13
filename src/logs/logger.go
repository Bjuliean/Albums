package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

type Logger interface {
	CreateLogger()
	CloseLogger()
	WriteError(err string, ip string)
}

type LogrusLogger struct {
	logsController  *logrus.Logger
	logsPath 		string
	logsFile		*os.File
}

func NewLogrusLogger(lpath string) *LogrusLogger {
	return &LogrusLogger{
		logsPath: lpath,
		logsFile: nil,
	}
}

func (l *LogrusLogger)CreateLogger() {
	file, err := os.OpenFile(l.logsPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	l.logsFile = file
	l.logsController = &logrus.Logger{
		Out: l.logsFile,
		Level: logrus.ErrorLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat: "[%lvl%]: %time% - %msg%",
		},
	}
}

func (l *LogrusLogger)CloseLogger() {
	l.logsFile.Close()
}

func (l *LogrusLogger)WriteError(err string, ip string) {
	l.logsController.Error(fmt.Sprintf("%s - %s\n", ip, err))
}