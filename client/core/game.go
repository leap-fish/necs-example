package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/necs-example/client"
	"github.com/leap-fish/necs-example/client/cfg"
	"github.com/leap-fish/necs-example/client/scenes"
	"github.com/leap-fish/necs-example/client/scenes/menu"
	"github.com/leap-fish/necs-example/client/scenes/primary"
	"github.com/leap-fish/necs/transports"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewGame() *Game {

	log.WithField("appconfig", cfg.AppConfig).Info("AppConfig contents initialized")

	addr := cfg.AppConfig["server_url"].(string)

	g := &Game{
		client: &client.Client{
			Transport: transports.NewWsClientTransport(addr),
		},
		ecs:         ecs.NewECS(donburi.NewWorld()),
		scaleFactor: 1,
	}

	gameScene := primary.NewPrimaryScene(g.client, g.ecs)
	g.scene = menu.NewMenuScene(gameScene)

	return g
}

type Game struct {
	scene scenes.Scene

	client *client.Client

	ecs *ecs.ECS

	scaleFactor int
}

func (g *Game) Update() error {
	if g.scene != nil {
		g.scene.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.scene != nil {
		g.scene.Draw(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / g.scaleFactor, outsideHeight / g.scaleFactor
}

func (g *Game) SetScene(scene scenes.Scene) {
	g.scene = scene
}
