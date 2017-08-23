// A battle between a kobold and a minotaur.
// Battle is based off of http://i.imgur.com/DxMejXv.png
// Shoutouts to >>9648141
// How to use:
// 0: Look for the //comments for help, or ask a question.
// 1: Copy the example fighter.
// 2: Give it a new variable name.
// 3: Edit the fields in the example fighter.
// 4: Change what fighter1 and fighter 2 are set to. Look for the //comment.
// 5: Run.

//package main
package v1

import (
	"fmt"
	"math/rand"
	"time"
)

type Fighter struct {
	name                string
	isAlive             bool
	health              int
	attack              int
	defense             int
	dodgeChance         float64
	ignoreDefenseChance float64
}

var (
	//{name in quotes, isAlive true/false, health integer, attack integer, defense integer, dodgeChance decimal, ignoreDefenseChance decimal}
	example Fighter = Fighter{"Example Fighter", true, 500, 50, 10, 0.30, 0.5}

	minotaur Fighter = Fighter{"Minotaur", true, 1200, 100, 20, 0.0, 0.0}
	kobold   Fighter = Fighter{"Kobold", true, 200, 30, 10, 0.95, 0.5}
)

func Main() {
	main()
}

func main() {
	var dice rand.Rand = *rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	var fighter1 Fighter
	var fighter2 Fighter

	var fighter1wins int
	var fighter2wins int
	var nowins int
	
	for rounds := 1000000; rounds > 0; rounds-- {
		fighter1 = kobold   //Change the first attacker.
		fighter2 = minotaur //Change the second attacker.

		for fighter1.health > 0 && fighter2.health > 0 {
			attack(&fighter1, &fighter2, &dice)
			attack(&fighter2, &fighter1, &dice)
		}
		
		if fighter1.health <= 0 && fighter2.health <= 0 {
			nowins++
		} else if fighter1.health > fighter2.health {
			fighter1wins++
		} else if fighter1.health < fighter2.health {
			fighter2wins++
		}
	}

	var totalWins int = fighter1wins + fighter2wins + nowins
	fmt.Println("Total fights:", totalWins)
	fmt.Println(fighter1.name, "wins:", fighter1wins)
	fmt.Println(fighter2.name, "wins:", fighter2wins)
	fmt.Println("Noone wins:", nowins)
	fmt.Println(fighter1.name, "won", 100.0*float64(fighter1wins)/float64(totalWins), "percent of the time")
	fmt.Println(fighter2.name, "won", 100.0*float64(fighter2wins)/float64(totalWins), "percent of the time")
	fmt.Println("Nobody won", 100.0*float64(nowins)/float64(totalWins), "percent of the time")
}

func attack(attacker *Fighter, defender *Fighter, dice *rand.Rand) {
	var attack int = 0
	if dice.Float64() >= defender.dodgeChance {
		attack = attacker.attack
		if attack < 0 {
			attack = 0
		}
	}
	defense := defender.defense
	if dice.Float64() < attacker.ignoreDefenseChance {
		defense = 0
	}
	if defense < 0 {
		defense = 0
	}
	damage := attack - defense
	if damage < 0 {
		damage = 0
	}

	defender.health = defender.health - damage
	if defender.health <= 0 {
		defender.isAlive = false
	}

	return
}
