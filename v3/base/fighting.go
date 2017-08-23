package base

import (
	"math/rand"
)

type Fighting interface {
	//IsAlive returns whether the Fighting is alive.
	IsAlive() bool
	//GetName returns the name of the Fighting.
	GetName() string
	//SendAttack sends an attack to a recipient.
	SendAttack(a Attack, recipient Fighting)
	RecieveAttack()
	//Reset sets the health and alive back to default values.
	Reset()
}
