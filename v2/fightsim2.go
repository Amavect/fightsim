// A battle between a knight and a wingless dragon.
// Battle is based off of http://i.imgur.com/P0rxUy8.png
// Shoutouts to >>10218103
// All of the freaking ampersands and stars because pass by reference needed to happen.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Fighter struct {
	name    string
	alive   bool
	health  int
	attack  int
	defense int
}

func (f *Fighter) checkAlive() {
	if f.health <= 0 {
		f.alive = false
	}
}

type Fighting interface {
	isAlive() bool
	getName() string
	getAttack(dice *rand.Rand) int
	takeDamage(attack int, dice *rand.Rand)
}

// Knight block
type Knight struct {
	Fighter
}

func (k *Knight) isAlive() bool {
	return k.alive
}

func (k *Knight) getName() string {
	return k.name
}

func (k *Knight) getAttack(dice *rand.Rand) int {
	if dice.Float64() <= 0.33 { 
		return k.attack * 3 // Knight hits 3x damage 33% of the time.
	} else {
		return k.attack
	}
}

func (k *Knight) takeDamage(attack int, dice *rand.Rand) {
	var damage int = attack - k.defense
	if damage < 0 {
		damage = 0 // Bound damage.
	}
	if dice.Float64() >= 0.90 {
		k.health = k.health - damage // 90% chance to dodge.
	}
	k.checkAlive()
}

// Dragon block
type Dragon struct {
	Fighter
}

func (d *Dragon) getName() string {
	return d.name
}

func (d *Dragon) isAlive() bool {
	return d.alive
}

func (d *Dragon) getAttack(dice *rand.Rand) int {
	if dice.Float64() <= 0.25 {
		return d.attack * 1000 // 25% chance to insta-kill. Calculation doesn't matter whether it's before or after the hit.
	} else {
		return d.attack
	}

}

func (d *Dragon) takeDamage(attack int, dice *rand.Rand) {
	var damage int = attack - d.defense
	if damage < 0 { // Bound damage.
		damage = 0
	}
	d.health = d.health - damage // Dragon has no % to dodge.
	d.checkAlive()
}

// Constructors
func newKnight() *Knight {
	return &Knight{Fighter{"Knight", true, 300, 80, 30}} // Knight has 300 hp, 80 atk, 30 def.
}

func newDragon() *Dragon {
	return &Dragon{Fighter{"Dragon", true, 1500, 120, 60}} // Dragon has 1500 hp, 120 atk, 60 def.
}

func main() {
	// Randomness based on time.
	var dice rand.Rand = *rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	// Counting wins.
	var fighter1wins int
	var fighter2wins int
	var bothLose int

	var fighter1 Fighting
	var fighter2 Fighting

	for rounds := 1000000; rounds > 0; rounds-- {

		fighter1 = newKnight() //Change the first attacker.
		fighter2 = newDragon() //Change the second attacker.

		for fighter1.isAlive() && fighter2.isAlive() {
			// Attacks happen simultaneously.
			fighter2.takeDamage(fighter1.getAttack(&dice), &dice)
			fighter1.takeDamage(fighter2.getAttack(&dice), &dice)
		}

		// Who won?
		if fighter1.isAlive() && fighter2.isAlive() == true {
			fmt.Println("idk they had tea or something like wtf make a good fight for the bloody crowd")
		} else if fighter1.isAlive() {
			fighter1wins++
		} else if fighter2.isAlive() {
			fighter2wins++
		} else {
			bothLose++
		}
	}

	// Print the everything.
	var totalWins int = fighter1wins + fighter2wins + bothLose
	fmt.Println("Total fights:", totalWins)
	fmt.Println(fighter1.getName(), "wins:", fighter1wins)
	fmt.Println(fighter2.getName(), "wins:", fighter2wins)
	fmt.Println("Times both died: ", bothLose)
	fmt.Println(fighter1.getName(), "won", 100.0*float64(fighter1wins)/float64(totalWins), "percent of the time.")
	fmt.Println(fighter2.getName(), "won", 100.0*float64(fighter2wins)/float64(totalWins), "percent of the time.")
	fmt.Println("Both died", 100.0*float64(bothLose)/float64(totalWins), "percent of the time.")
}