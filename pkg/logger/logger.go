package logger

import (
	"github.com/ArWhale/WTest/internal/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func NewLogger(viper *viper.Viper) logrus.FieldLogger {
	logOutput := strings.ToUpper(viper.GetString(consts.ServiceLogOutputKey))

	if strings.ToUpper(logOutput) != "STDOUT" {
		logFile, err := os.OpenFile(logOutput, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logrus.WithError(err).Fatalf("can't open log file")
		}
		logrus.SetOutput(logFile)
	} else {
		logrus.SetOutput(os.Stdout)
	}

	switch strings.ToLower(consts.ServiceLogFormatKey) {
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
