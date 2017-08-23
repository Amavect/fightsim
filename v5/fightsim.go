// A refactoring of fightsim v1 for batch processing.
// Hopefully there is a speedup. Will time it.
// Battle is based off of http://i.imgur.com/DxMejXv.png
// Shoutouts to >>9648141

//package main
package v5

import (
	"fmt"
	"math/rand"
	"time"
)

type Fighter struct {
	Name                string
	Health              int
	Attack              int
	Defense             int
	DodgeChance         float64
	IgnoreDefenseChance float64
}

func newMinotaur() *Fighter {
	return &Fighter{"Minotaur", 1200, 100, 20, 0.0, 0.0}
}

func newKobold() *Fighter {
	return &Fighter{"Kobold", 200, 30, 10, 0.95, 0.5}
}

func attack(attacker *Fighter, defender *Fighter) {
	var attack int = 0
	if rand.Float64() >= defender.DodgeChance {
		attack = attacker.Attack
		if attack < 0 {
			attack = 0
		}
	}	
	defense := defender.Defense
	if rand.Float64() < attacker.IgnoreDefenseChance {
		defense = 0
	}
	if defense < 0 {
		defense = 0
	}
	damage := attack - defense
	if damage < 0 {
		damage = 0
	}

	defender.Health = defender.Health - damage
	return
}

func battle(bs [][2]*Fighter) (nowins int, f1wins int, f2wins int, errwins int) {
	for i := range bs {
		for (bs[i][0].Health > 0) && (bs[i][1].Health > 0) { 
			attack(bs[i][0], bs[i][1])
			attack(bs[i][1], bs[i][0])
		}
		
		
		if (bs[i][0].Health > 0) == (bs[i][1].Health > 0) { //XOR
			if bs[i][0].Health < 0 {
				nowins++
			} else {
				errwins++
			}
		} else if (bs[i][0].Health > 0) {
			f1wins++
		} else {
			f2wins++
		}
	}
	return 
}

func Main() {
	main()
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	
	var rounds int = 1000000
	
	if rounds < 1 {
		rounds = 1
	}
	
	var battles [][2]*Fighter
	
	var fighter1wins int
	var fighter2wins int
	var nowins int
	var errwins int
	
	batches := rounds / 1024
	extras := rounds % 1024
	
	for i := 0; i < batches; i++{
		battles = make([][2]*Fighter, 1024, 1024)
		for j := 0; j < 1024; j++ {
			battles[j][0] = newKobold()
			battles[j][1] = newMinotaur()
		}

		var nw, f1, f2, errw int = battle(battles)
		fighter1wins += f1
		fighter2wins += f2
		nowins += nw
		errwins += errw
	}
	{
		battles = make([][2]*Fighter, extras, extras)
		for j := 0; j < extras; j++ {
			battles[j][0] = newKobold()
			battles[j][1] = newMinotaur()
		}

		var nw, f1, f2, errw int = battle(battles)
		fighter1wins += f1
		fighter2wins += f2
		nowins += nw
		errwins += errw
	}
	
	fmt.Println("ERR:", errwins)
	var totalWins int = fighter1wins + fighter2wins + nowins
	fmt.Println("Total fights:", totalWins)
	fmt.Println(battles[0][0].Name, "wins:", fighter1wins)
	fmt.Println(battles[0][1].Name, "wins:", fighter2wins)
	fmt.Println("Noone wins:", nowins)
	fmt.Println(battles[0][0].Name, "won", 100.0*float64(fighter1wins)/float64(totalWins), "percent of the time")
	fmt.Println(battles[0][1].Name, "won", 100.0*float64(fighter2wins)/float64(totalWins), "percent of the time")
	fmt.Println("Nobody won", 100.0*float64(nowins)/float64(totalWins), "percent of the time")
}
