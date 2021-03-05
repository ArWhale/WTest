package server

import "github.com/sirupsen/logrus"

func LoggingActionMessage(logger logrus.FieldLogger, message string) {
	logger.WithFields(logrus.Fields{
		"context": "handlers",
		"message": message,
	}).Error()
}

func LoggingActionError(logger logrus.FieldLogger, action string, err error) {
	logger.WithFields(logrus.Fields{
		"context": "handlers",
		"message": action,
	}).Error(err)
}

func LoggingRepoError(logger logrus.FieldLogger, action string, err error) {
	logger.WithFields(logrus.Fields{
		"context": "db",
		"message": action,
	}).Error(err)
}
