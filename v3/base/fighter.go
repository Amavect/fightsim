package base

import ()

//Fighter is a helper struct with general implementations for the Fighting interface.
type Fighter struct {
	Name    string
	Alive   bool
	Health  int
	Attack  int
	Defense int
	StatusEffectList []StatusEffect
}

//Newfighter initializes a Fighter.
func Newfighter() Fighter {
	return Fighter{"", false, 0, 0, 0, make([]StatusEffect, 0)}
}

func (f *Fighter) GetName() string {
	return f.Name
}

func (f *Fighter) IsAlive() bool {
	return f.Alive
}

//CheckAlive sets the Alive bool to false if Health is less than 0.
func (f *Fighter) CheckAlive() {
	if f.Health <= 0 {
		f.Alive = false
	}
}

func (f *Fighter) AddStatusEffect(se StatusEffect) {
	f.StatusEffectList = append(f.StatusEffectList, se)
}

func (f *Fighter) RemoveStatusEffect(index int) {
    f.StatusEffectList[len(f.StatusEffectList)-1], f.StatusEffectList[index] = f.StatusEffectList[index], f.StatusEffectList[len(f.StatusEffectList)-1]
    f.StatusEffectList = f.StatusEffectList[:len(f.StatusEffectList)-1]
}

func (f *Fighter) UpdateStatusEffects() {
	for index, stat := range f.StatusEffectList {
		if stat.IsExpired() {
			f.RemoveStatusEffect(index)
		}
	}
}