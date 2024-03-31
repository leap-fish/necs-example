package core

import (
	"github.com/yohamta/donburi"
)

type SubSystem interface {
	Initialize(world donburi.World)
	Tick()
}
