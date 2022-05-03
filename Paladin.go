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

func (p Paladin) move() {
	p.AHero.move()
}

func (p Paladin) die() {
	p.AHero.die()
}

func (p Paladin) status() {
	p.AHero.status()
}

func (p Paladin) shout() {
	p.AHero.shout()
}

func (p Paladin) attack(oponent *Orc) {
	fmt.Printf("%v\n", "Take that, ugly beast!")
	oponent.life = oponent.life - p.power
	if oponent.life == 0 {
		oponent.AHero.die()
	}
}
