package base

import (

)

//Ouchy!
type StatusEffect interface {
	//GetName returns the name of the StatusEffect.
	GetName() string
	//GetLength returns how many turns the StatusEffect will last.
	GetLength() int
	//SetLength sets how many turns the StatusEffect will last.
	SetLength(i int) 
	//IsExpired checks if the StatusEffect no longer applies.
	IsExpired() bool
	//Performs the status effect on the Fighting.
	Perform(f Fighting)
}