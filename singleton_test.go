package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func Test(t *testing.T) {
	logrus.Info("Hello mowr")
	logrus.Warn("Hello mowr")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello mowr")
	logrus.Warn("Hello mowr")
}
