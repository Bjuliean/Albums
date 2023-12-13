package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

type Logger interface {
	CreateLogger()
	CloseLogger()
	WriteError(err string)
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
	file, _ := os.OpenFile(l.logsPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	l.logsFile = file
	l.logsController = &logrus.Logger{
		Out: l.logsFile,
		Level: logrus.ErrorLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2001-01-04 15:04:05",
			LogFormat: "[%lvl%]: %time% - %msg%",
		},
	}
}

func (l *LogrusLogger)CloseLogger() {
	l.logsFile.Close()
}

func (l *LogrusLogger)WriteError(err string) {
	l.logsController.Error(err + "\n")
}