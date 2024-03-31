package systems

import (
	"github.com/yohamta/donburi"
)

type OtherSubSystem struct {
	world donburi.World
}

func (s *OtherSubSystem) Initialize(world donburi.World) {
	s.world = world
}
func (s *OtherSubSystem) Tick() {
	if s.world == nil {
		return
	}
}
