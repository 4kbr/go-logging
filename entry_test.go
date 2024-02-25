package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)

	entry.WithField("username", "user1").Info("hello world")
	entry.WithField("username", "user22").WithField("name", "user").Info("hello world")
}
