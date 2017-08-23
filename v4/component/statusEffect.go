package component

import (

)

type Afflicted struct {
	Statuses []StatusEffect
}

type StatusEffect struct {
	StatusStrat StatusID
	Duration int
}

type StatusAttack func(a Attacker) int
type StatusDefend func(d Defender) int
type StatusHealth func(h Health) int
type StatusID int

const (
	//AttackBasic is the ID for the attackBasic() attack strategy.
	STATUS_BLEEDING StatusID = iota
)

var bleeding StatusStrategy = func() int {
	return baseAtk
}

func StatusList() []StatusStrategy {
	var statusList []StatusStrategy= make([]StatusStrategy, 10, 10)
	statusList[BLEEDING] = bleeding
	return statusList
}