package core

import (
	"fmt"
	"github.com/leap-fish/necs-example/server/systems"
	"github.com/leap-fish/necs-example/shared"
	"github.com/leap-fish/necs-example/shared/sliceutil"
	"github.com/leap-fish/necs/esync/srvsync"
	"github.com/leap-fish/necs/router"
	"github.com/leap-fish/necs/transports"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"nhooyr.io/websocket"
	"runtime"
	"time"
)

const (
	TickRate = 16
)

type Server struct {
	host       transports.NetworkTransport
	subSystems []SubSystem
	ecs        *ecs.ECS
}

func NewServer() *Server {
	return &Server{
		ecs: ecs.NewECS(donburi.NewWorld()),
		subSystems: []SubSystem{
			&systems.BasicSubSystem{},
			&systems.OtherSubSystem{},
		},
		host: transports.NewWsServerTransport(
			uint(shared.EnvServerPort),
			shared.EnvBindAddress,
			&websocket.AcceptOptions{
				InsecureSkipVerify: true,
			},
		),
	}
}

func (s *Server) Start() {
	log.
		WithField("arch", runtime.GOARCH).
		WithField("goos", runtime.GOOS).
		WithField("envMode", shared.EnvEnvironment).
		WithField("bindAddr", fmt.Sprintf("%s:%d", shared.EnvBindAddress, uint(shared.EnvServerPort))).
		Info("Starting the server")

	log.
		WithField("subSystemCount", len(s.subSystems)).
		WithField("subSystems", sliceutil.TypeList(s.subSystems)).
		WithField("worldId", s.ecs.World.Id()).
		Info("Initializing sub systems")

	srvsync.UseEsync(s.ecs.World)

	srvsync.AddNetworkFilter(func(client *router.NetworkClient, entry *donburi.Entry) bool {
		if !entry.HasComponent(systems.PositionComponent) {
			return true
		}

		pos := systems.PositionComponent.Get(entry)
		return pos.X < 1368 && pos.Y < 768
	})

	for _, sys := range s.subSystems {
		sys.Initialize(s.ecs.World)
	}

	go s.startTicking()

	err := s.host.Start()
	if err != nil {
		log.WithError(err).Error("Could not start server host")
	}
}

func (s *Server) startTicking() {
	for range time.Tick(time.Second / TickRate) {
		for _, sys := range s.subSystems {
			sys.Tick()
		}

		err := srvsync.DoSync()
		if err != nil {
			log.
				WithError(err).
				Error("Unable to perform esync.DoSync")
		}
	}
}
