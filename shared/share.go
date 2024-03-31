package shared

import "github.com/yohamta/donburi"

type GlobalStats struct {
	PlayerCount int
	TotalEnts   int
}

var GlobalStatsComponent = donburi.NewComponentType[GlobalStats]()
