package component

import (

)

//Attack block

type Attacker struct {
	AtkBase int
	AtkStrat AttackID
}

type AttackStrategy func(baseAtk int) int
type AttackID int

const (
	//AttackBasic is the ID for the attackBasic() attack strategy.
	ATTACK_BASIC AttackID = iota
	ATTACK_SUPER
)

var attackBasic AttackStrategy = func(baseAtk int) int {
	return baseAtk
}

var attackSuper AttackStrategy = func(baseAtk int) int{
	return baseAtk * 10000000000
}

func AttackList() []AttackStrategy {
	var attackList []AttackStrategy= make([]AttackStrategy, 10, 10)
	attackList[ATTACK_BASIC] = attackBasic
	attackList[ATTACK_SUPER] = attackSuper
	return attackList
}

//Defense block

type Defender struct {
	DefBase int
	DefStrat DefenseID
}

type DefenseStrategy func(baseAtk int) int
type DefenseID int

const (
	//DefendBasic is the ID for the defendBasic() defence strategy.
	DEFEND_BASIC DefenseID = iota
)

var defendBasic DefenseStrategy = func(baseDef int) int {
	return baseDef
}

func DefendList() []DefenseStrategy {
	var defendList []DefenseStrategy = make([]DefenseStrategy, 10, 10)
	defendList[DEFEND_BASIC] = defendBasic
	return defendList
}
