package main

import (
	"fmt"
)

type Orc struct {
	*AHero
}

func init() {
	fmt.Printf("%v\n", "I'm born! Ugh! Ugh! Ugh!")
}

func (p Orc) attack(oponent *Paladin) {
	fmt.Printf("%v\n", "Eat that human")
	oponent.life = oponent.life - p.power
	if oponent.life == 0 {
		oponent.die()
	}
}
