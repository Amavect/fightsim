package system

import (
	"fmt"
	"math/rand"
	"github.com/Amavect/fightsim/v4/entity"
	"github.com/Amavect/fightsim/v4/component"
)

type WinFlag int

const (
	Error WinFlag = iota - 1
	NoWin
	P1win
	P2win
)

//Battle simulates a battle between the two fighters.
//If fighter 1 wins, returns true. If fighter 2 wins, returns false.
func Battle(f1 *entity.Fighter, f2 *entity.Fighter) WinFlag {
	{
		atkList := component.AttackList()
		defList := component.DefendList()
		statList := component.StatusList()
		
		//Attack removes health from the defending fighter.
		attack := func(atk *entity.Fighter, def *entity.Fighter) {
			var damage int = atkList[atk.AtkStrat](atk.AtkBase) - defList[def.DefStrat](def.DefBase)
			if damage < 0 {
				damage = 0
			}
			def.Health -= damage
			 
		}
		
		applyStats := func(f *entity.Fighter) {
			
		}
		
		for isAlive(f1) && isAlive(f2) {

			attack(f2, f1)
			attack(f1, f2)
			// Attacks happen simultaneously.
			if !(isAlive(f1) && isAlive(f2)){
				break;
			}
			
			
		}
	}

	fmt.Println(isAlive(f1), isAlive(f2))
	// Who won?
	if isAlive(f1) && isAlive(f2) {
		return Error
	} else if isAlive(f1) {
		return P1win
	} else if isAlive(f2) {
		return P2win
	} else {
		return NoWin
	}
}

func isAlive(f *entity.Fighter) bool {
	if f.Health > 0 {
		return true
	} else {
		return false
	}
}
