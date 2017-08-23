package entity

import (
	"github.com/Amavect/fightsim/v4/component"
)

type Fighter struct {
	Name   string
	component.Health
	component.Defender
	component.Attacker
	component.Afflicted
}

func NewMinotaur() *Fighter {
	return &Fighter{
		"Minotaur",
		1200,
		component.Defender{
			20,
			component.DEFEND_BASIC,
		},
		component.Attacker{
			100,
			component.ATTACK_BASIC,
		},
		component.Afflicted{
			make([]component.StatusEffect, 5),
		},
	}
}

func NewKobold() *Fighter {
	return &Fighter{
		"Kobold",
		200,
		component.Defender{
			10,
			component.DEFEND_BASIC,
		},
		component.Attacker{
			100,
			component.ATTACK_BASIC,
		},
		component.Afflicted{
			make([]component.StatusEffect, 5),
		},
	}
}