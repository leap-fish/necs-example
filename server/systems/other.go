package systems

import (
	"github.com/leap-fish/necs/router"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
)

type OtherSubSystem struct {
	world donburi.World
}

func (s *OtherSubSystem) Initialize(world donburi.World) {
	s.world = world

	router.OnConnect(func(sender *router.NetworkClient) {
		log.Infof("Client %s connected to the server!", sender.Id())
	})
	router.OnDisconnect(func(sender *router.NetworkClient, err error) {
		log.Infof("Client %s disconnected from the server! / Reason [%s]", sender.Id(), err)
	})
}
func (s *OtherSubSystem) Tick() {
	if s.world == nil {
		return
	}
}
