package main

import (
	"github.com/Amavect/fightsim/v3/base"
	"github.com/Amavect/fightsim/v3/fighters"
)

func main() {
	var fighter1 base.Fighting = fighters.NewWolf()
	var fighter2 base.Fighting = fighters.NewSheep()
	
	fighter1.GetName()
	base.RunSimulation(fighter1, fighter2, 1)
}