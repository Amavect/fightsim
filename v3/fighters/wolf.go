package fighters

import (
	"github.com/Amavect/fightsim/v3/base"
	"math/rand"
)

//awoo
type wolf struct {
	base.Fighter
}

//NewWolf returns a *wolf.
func NewWolf() *wolf {
	return &wolf{base.Fighter{"Wolf", true, 500, 80, 30, make([]base.StatusEffect, 0)}} //TODO: use Fighter initializer, set other values later.
}

func (w *wolf) GetAttack(dice *rand.Rand) int {
	return w.Attack
}

func (w *wolf) TakeDamage(attack int, dice *rand.Rand) {
	var damage int
	if dice.Float64() < 0.66 {
		damage = attack 
	}
	
	w.Health -= damage
	
	w.CheckAlive()
}

func (w *wolf) Reset() {
	w.Health = 500
	w.Alive = true
}

//Causes 15 dmg per turn and lasts forever.
type bleeding struct {
	base.StatusEffect
	length int
}

func newBleeding() bleeding {
	return bleeding{}
}

func (b *bleeding) GetName() string {
	return "bleeding"
}

func (b *bleeding) GetLength() int {
	return b.length
}

func (b *bleeding) SetLength(l int) {
	b.length = l
}

func (b *bleeding) IsExpired() bool {
	if b.length > 0 {
		return false
	} else {
		return true
	}
}