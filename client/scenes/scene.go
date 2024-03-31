package scenes

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
	Configure()
}
