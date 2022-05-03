package main

import (
	"fmt"
)

type Paladin struct {
	*AHero
}

func init() {
	fmt.Printf("%v\n", "Ready to fight for our king")
}

func (p Paladin) attack(oponent *Orc) {
	fmt.Printf("%v\n", "Take that, ugly beast!")
	oponent.life = oponent.life - p.power
	if oponent.life == 0 {
		oponent.die()
	}
}
