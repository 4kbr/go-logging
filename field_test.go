package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "user1").Info("hello world")
	logger.WithField("username", "user22").WithField("name", "user").Info("hello world")
}
func TestFields(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{"username": "user1",
		"name": "user",
	}).Info("hello world")

}
