package main

import (
	"github.com/leap-fish/necs-example/server/core"
	"github.com/leap-fish/necs-example/shared"
	"github.com/sirupsen/logrus"
)

func init() {
	if shared.EnvEnvironment == shared.EnvModeDevelopment {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	}
}

func main() {
	s := core.NewServer()
	s.Start()
}
