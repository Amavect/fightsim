package base

import (
	"fmt"
	"math/rand"
	"time"
)

//Runs the simulation.
func RunSimulation(fighter1 Fighting, fighter2 Fighting, numSims int) {
	// Randomness based on time.
	var dice rand.Rand = *rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	// Counting wins.
	var fighter1wins int
	var fighter2wins int
	var bothLose int

	for rounds := numSims; rounds > 0; rounds-- {
		fmt.Println(rounds)
		for fighter1.IsAlive() && fighter2.IsAlive() {
			// Attacks happen simultaneously.
			fighter2.TakeDamage(fighter1.GetAttack(&dice), &dice)
			fighter1.TakeDamage(fighter2.GetAttack(&dice), &dice)
		}

		// Who won?
		if fighter1.IsAlive() && fighter2.IsAlive() == true {
			fmt.Println("idk they had tea or something like wtf make a good fight for the bloody crowd")
		} else if fighter1.IsAlive() {
			fighter1wins++
		} else if fighter2.IsAlive() {
			fighter2wins++
		} else {
			bothLose++
		}
		
		fighter1.Reset()
		fighter2.Reset()
	}

	// Print the everything.
	var totalWins int = fighter1wins + fighter2wins + bothLose
	fmt.Println("Total fights:", totalWins)
	fmt.Println(fighter1.GetName(), "wins:", fighter1wins)
	fmt.Println(fighter2.GetName(), "wins:", fighter2wins)
	fmt.Println("Times both died: ", bothLose)
	fmt.Println(fighter1.GetName(), "won", 100.0*float64(fighter1wins)/float64(totalWins), "percent of the time.")
	fmt.Println(fighter2.GetName(), "won", 100.0*float64(fighter2wins)/float64(totalWins), "percent of the time.")
	fmt.Println("Both died", 100.0*float64(bothLose)/float64(totalWins), "percent of the time.")
}
