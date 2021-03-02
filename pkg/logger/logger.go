package logger

import (
	"github.com/SArtemJ/WTest/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func NewLogger(viper *viper.Viper) *logrus.Entry {
	logOutput := strings.ToUpper(viper.GetString(config.ServiceLogOutputKey))
	if strings.ToUpper(logOutput) != "STDOUT" {
		logFile, err := os.OpenFile(logOutput, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logrus.WithError(err).Fatalf("can't open log file")
		}
		logrus.SetOutput(logFile)
	} else {
		logrus.SetOutput(os.Stdout)
	}

	switch strings.ToLower(config.ServiceLogFormatKey) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: false,
		})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	return logrus.NewEntry(logrus.StandardLogger())
}
