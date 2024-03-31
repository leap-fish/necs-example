package systems

import (
	"fmt"
	"github.com/leap-fish/necs-example/shared"
	"github.com/leap-fish/necs/esync"
	"github.com/leap-fish/necs/esync/srvsync"
	"github.com/leap-fish/necs/router"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"math/rand"
	"time"
)

type Name string

type PositionData struct {
	X, Y float32
}

var PositionComponent = donburi.NewComponentType[PositionData]()
var NameComponent = donburi.NewComponentType[Name]()

type BasicSubSystem struct {
	world donburi.World

	src  rand.Source
	rand *rand.Rand
}

func (s *BasicSubSystem) Initialize(world donburi.World) {
	s.world = world
	s.src = rand.NewSource(time.Now().UnixNano())
	s.rand = rand.New(s.src)

	_ = esync.RegisterComponent(10, PositionData{}, PositionComponent)
	_ = esync.RegisterComponent(11, Name(""), NameComponent)
	_ = esync.RegisterComponent(12, shared.GlobalStats{}, shared.GlobalStatsComponent)

	stats := s.world.Create(shared.GlobalStatsComponent)
	_ = srvsync.NetworkSync(s.world, &stats, shared.GlobalStatsComponent)

	for i := 0; i < 5; i++ {
		ent := s.world.Create(PositionComponent)
		PositionComponent.Set(world.Entry(ent), &PositionData{X: 5000, Y: 5000})

	}
	for i := 0; i < 10_000; i++ {
		ent := s.world.Create(PositionComponent, NameComponent)
		name := Name(fmt.Sprintf("Meow %d", i))
		NameComponent.Set(world.Entry(ent), &name)
		PositionComponent.Set(world.Entry(ent), &PositionData{X: 5000, Y: 5000})

		_ = srvsync.NetworkSync(s.world, &ent, PositionComponent)
	}

	go func() {
		deleteQuery := donburi.NewQuery(filter.And(filter.Contains(esync.NetworkIdComponent), filter.Not(filter.Contains(shared.GlobalStatsComponent))))
		for range time.Tick(time.Second * 8) {
			// Delete old
			deletionEnt, ok := deleteQuery.First(world)
			if ok {
				world.Remove(deletionEnt.Entity())
			}
		}
	}()
}
func (s *BasicSubSystem) Random() *rand.Rand {
	r := rand.New(s.src)
	return r
}

func (s *BasicSubSystem) Tick() {
	if s.world == nil {
		return
	}

	q := donburi.NewQuery(filter.Contains(PositionComponent))

	q.Each(s.world, func(entry *donburi.Entry) {
		pos := PositionComponent.Get(entry)

		pos.X = 1368 * rand.Float32() * 10
		pos.Y = 768 * rand.Float32() * 10
	})

	{
		statQuery := donburi.NewQuery(filter.Contains(shared.GlobalStatsComponent))
		st, ok := statQuery.First(s.world)
		if !ok {
			return
		}

		stats := shared.GlobalStatsComponent.Get(st)
		stats.PlayerCount = len(router.Peers())
		stats.TotalEnts = s.world.Len()
	}

}
