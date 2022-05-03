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

func (p Orc) move() {
	p.AHero.move()
	fmt.Println("Ugh, come here!")
}

func (p Orc) die() {
	p.AHero.die()
}

func (p Orc) status() {
	p.AHero.status()
}

func (p Orc) shout() {
	p.AHero.shout()
}

func (p Orc) attack(oponent *Paladin) {
	fmt.Printf("%v\n", "Eat that human")
	oponent.life = oponent.life - p.power
	if oponent.life == 0 {
		oponent.AHero.die()
	}
}
