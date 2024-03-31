package archetype

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

const (
	DefaultEcsLayer ecs.LayerID = iota
)

/*
var (
	Platform = newArchetype(
		tags.Platform,
		components.Object,
	)
	FloatingPlatform = newArchetype(
		tags.FloatingPlatform,
		components.Object,
		components.Tween,
	)
	Player = newArchetype(
		tags.Player,
		components.Player,
		components.Object,
	)
	Ramp = newArchetype(
		tags.Ramp,
		components.Object,
	)
	Space = newArchetype(
		components.Space,
	)
	Wall = newArchetype(
		tags.Wall,
		components.Object,
	)
)
*/

type Archetype struct {
	components []donburi.IComponentType
}

func NewArchetype(cs ...donburi.IComponentType) *Archetype {
	return &Archetype{
		components: cs,
	}
}

func (a *Archetype) Spawn(world donburi.World, additionalComponents ...donburi.IComponentType) *donburi.Entry {
	e := world.Entry(world.Create(
		append(a.components, additionalComponents...)...,
	))
	return e
}
