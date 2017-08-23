package fighters

import (
	"github.com/Amavect/fightsim/v3/base"
	"math/rand"
)

//baaa
type sheep struct {
	base.Fighter
}

//NewSheep returns a *sheep.
func NewSheep() *sheep {
	return &sheep{base.Fighter{"Sheep", true, 300, 60, 20, make([]base.StatusEffect, 0)}}
}

func (s *sheep) GetAttack(dice *rand.Rand) int {
	return s.Attack
}

func (s *sheep) TakeDamage(attack int, dice *rand.Rand) {
	var damage int
	if dice.Float64() < 0.50 {
		damage = attack 
	}
	
	s.Health -= damage
	
	s.CheckAlive()
}

func (s *sheep) Reset() {
	s.Health = 300
	s.Alive = true
}