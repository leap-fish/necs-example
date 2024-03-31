package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/leap-fish/necs-example/client"
	"github.com/leap-fish/necs-example/client/scenes"
	"github.com/leap-fish/necs-example/shared/archetype"
	log "github.com/sirupsen/logrus"
	"sync"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewMenuScene(switchScene scenes.Scene) *MenuScene {
	return &MenuScene{
		ecs:         ecs.NewECS(donburi.NewWorld()),
		systems:     []ecs.System{},
		renderers:   []ecs.RendererWithArg[ebiten.Image]{},
		once:        &sync.Once{},
		switchScene: switchScene,
	}
}

type MenuScene struct {
	ecs  *ecs.ECS
	once *sync.Once

	systems   []ecs.System
	renderers []ecs.RendererWithArg[ebiten.Image]

	switchScene scenes.Scene
}

func (s *MenuScene) Update() {
	s.once.Do(s.Configure)

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		client.Instance.SetScene(s.switchScene)
		return
	}
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	s.ecs.Draw(screen)

	ebitenutil.DebugPrint(screen, "Menu (press SPACE to join)")
}

func (s *MenuScene) Configure() {
	log.Info("Game started..")

	for _, system := range s.systems {
		s.ecs.AddSystem(system)
	}

	for _, renderer := range s.renderers {
		s.ecs.AddRenderer(archetype.DefaultEcsLayer, renderer)
	}
}
