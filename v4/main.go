package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/Amavect/fightsim/v4/system"
	"github.com/Amavect/fightsim/v4/entity"
)

func main() {
	rand.Seed(rand.NewSource(int64(time.Now().Nanosecond())))
	
	fighter1 := entity.NewMinotaur()
	fighter2 := entity.NewKobold()
	
	var (
		errorwins int
		bothLose int
		fighter1wins int
		fighter2wins int
	)
	
	for i := 10; i > 0; i-- {
		switch system.Battle(fighter1, fighter2){
		case system.Error:
		errorwins++
		case system.NoWin:
		bothLose++
		case system.P1win:
		fighter1wins++
		case system.P2win:
		fighter2wins++
		}
	}
	
	var totalWins int = fighter1wins + fighter2wins + bothLose
	fmt.Println("Total fights:", totalWins)
	fmt.Println(fighter1.Name, "wins:", fighter1wins)
	fmt.Println(fighter2.Name, "wins:", fighter2wins)
	fmt.Println("Times both died: ", bothLose)
	fmt.Println(fighter1.Name, "won", 100.0*float64(fighter1wins)/float64(totalWins), "percent of the time.")
	fmt.Println(fighter2.Name, "won", 100.0*float64(fighter2wins)/float64(totalWins), "percent of the time.")
	fmt.Println("Both died", 100.0*float64(bothLose)/float64(totalWins), "percent of the time.")
	
}