package client

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/necs-example/client/scenes"
	"github.com/leap-fish/necs/router"
)

var Instance GameInstance

var Network *router.NetworkClient

type GameInstance interface {
	ebiten.Game

	SetScene(scene scenes.Scene)
}
