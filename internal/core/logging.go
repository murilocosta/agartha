package core

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type LogLevel logrus.Level
type LogEvent map[string]interface{}

const (
	LogInfo  LogLevel = LogLevel(logrus.InfoLevel)
	LogWarn  LogLevel = LogLevel(logrus.WarnLevel)
	LogErr   LogLevel = LogLevel(logrus.ErrorLevel)
	LogFatal LogLevel = LogLevel(logrus.FatalLevel)
	LogDebug LogLevel = LogLevel(logrus.DebugLevel)
)

func InitLogger(level LogLevel, logPath string, linkPath string) error {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.Level(level))
	writer, err := rotatelogs.New(
		logPath,
		rotatelogs.WithLinkName(linkPath),
		rotatelogs.WithMaxAge(-1),
	)
	if err != nil {
		return err
	}
	logrus.SetOutput(writer)
	return nil
}

func Log(level LogLevel, msg string, ctx LogEvent) {
	reqLogger := logrus.WithFields(logrus.Fields(ctx))
	logWithLevel(reqLogger, level, msg)
}

func LogError(level LogLevel, msg string, err error) {
	reqLogger := logrus.WithError(err)
	logWithLevel(reqLogger, level, msg)
}

func logWithLevel(entry *logrus.Entry, level LogLevel, msg string) {
	switch level {
	case LogInfo:
		entry.Info(msg)
	case LogWarn:
		entry.Warn(msg)
	case LogErr:
		entry.Error(msg)
	case LogFatal:
		entry.Fatal(msg)
	case LogDebug:
		entry.Debug(msg)
	}
}
