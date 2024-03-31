package primary

import (
	"bytes"
	"context"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/leap-fish/necs-example/client"
	"github.com/leap-fish/necs-example/server/systems"
	"github.com/leap-fish/necs-example/shared"
	"github.com/leap-fish/necs-example/shared/archetype"
	"github.com/leap-fish/necs/esync"
	"github.com/leap-fish/necs/esync/clisync"
	"github.com/leap-fish/necs/router"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi/features/debug"
	"github.com/yohamta/donburi/filter"
	"image/color"
	"nhooyr.io/websocket"
	"sync"
	"time"

	"github.com/yohamta/donburi"
	decs "github.com/yohamta/donburi/ecs"
)

func NewPrimaryScene(client *client.Client, ecs *decs.ECS) *PrimaryScene {
	return &PrimaryScene{
		client:    client,
		ecs:       ecs,
		world:     ecs.World,
		systems:   []decs.System{},
		renderers: []decs.RendererWithArg[ebiten.Image]{},
		once:      &sync.Once{},
	}
}

func renderer(ecs *decs.ECS, image *ebiten.Image) {
	var out bytes.Buffer
	for _, c := range debug.GetEntityCounts(ecs.World) {
		out.WriteString(c.String())
		out.WriteString("\n")
	}
	out.WriteString("\n")

	entCol := color.RGBA{R: 180, G: 100, A: 255}

	var posData string
	q := donburi.NewQuery(filter.Contains(systems.PositionComponent, esync.NetworkIdComponent))
	q.Each(ecs.World, func(entry *donburi.Entry) {

		var name string
		if entry.HasComponent(systems.NameComponent) {
			name = fmt.Sprintf(" (Name: %s)", string(*systems.NameComponent.Get(entry)))
		}

		nid := esync.GetNetworkId(entry)
		pos := systems.PositionComponent.Get(entry)

		posData += fmt.Sprintf("> [%d/netID:%d] POS: %#v%s\n", entry.Id(), *nid, *pos, name)

		ebitenutil.DrawRect(image, float64(pos.X), float64(pos.Y), 6, 6, entCol)
	})

	statQuery := donburi.NewQuery(filter.Contains(shared.GlobalStatsComponent))
	var statsString string
	st, ok := statQuery.First(ecs.World)
	if ok {
		stats := shared.GlobalStatsComponent.Get(st)
		statsString = fmt.Sprintf("%#v", *stats)
	}

	output := fmt.Sprintf("%s\n%s\n%s\n%s", "Game scene "+time.Now().String(), statsString, out.String(), posData)
	ebitenutil.DebugPrint(image, output)
}

type PrimaryScene struct {
	ecs    *decs.ECS
	world  donburi.World
	once   *sync.Once
	client *client.Client

	systems   []decs.System
	renderers []decs.RendererWithArg[ebiten.Image]
}

func (s *PrimaryScene) Configure() {
	router.OnConnect(func(client *router.NetworkClient) {
		log.Info("Connected to server!")
	})

	s.renderers = append(s.renderers, renderer)

	_ = esync.RegisterComponent(10, systems.PositionData{}, systems.PositionComponent)
	_ = esync.RegisterComponent(11, systems.Name(""), systems.NameComponent)
	_ = esync.RegisterComponent(12, shared.GlobalStats{}, shared.GlobalStatsComponent)

	clisync.RegisterClient(s.world)

	go func() {
		err := s.client.Transport.Start(func(conn *websocket.Conn) {
			client.Network = router.NewNetworkClient(context.Background(), conn)
		})
		if err != nil {
			log.WithError(err).Error("Unable to dial server")
		}
	}()

	for _, system := range s.systems {
		s.ecs.AddSystem(system)
	}

	for _, renderer := range s.renderers {
		s.ecs.AddRenderer(archetype.DefaultEcsLayer, renderer)
	}
}

func (s *PrimaryScene) Draw(screen *ebiten.Image) {
	s.ecs.Draw(screen)
}

func (s *PrimaryScene) Update() {
	s.once.Do(s.Configure)
}
