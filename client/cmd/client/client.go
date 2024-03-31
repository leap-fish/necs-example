package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/necs-example/client"
	"github.com/leap-fish/necs-example/client/core"
	"github.com/leap-fish/necs-example/shared"
	"log"
	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {
	if shared.EnvEnvironment == shared.EnvModeDevelopment {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: runtime.GOARCH != "wasm"})
	}
}

func main() {
	ebiten.SetWindowSize(1366, 768)
	ebiten.SetWindowTitle("Leapfish Technologies - NECS example")

	game := core.NewGame()
	client.Instance = game

	if err := ebiten.RunGame(client.Instance); err != nil {
		log.Fatal(err)
	}
}
