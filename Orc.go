package main

import (
	"fmt"
)

type Orc struct {
	AHero
}

func init() {
	fmt.Printf("%v\n", "I'm born! Ugh! Ugh! Ugh!")
}

func (p Orc) move() {
	ah := new(AHero)
	ah.position, ah.speed = p.position, p.speed
	ah.move()
	fmt.Println("Ugh, come here!")
}

func (p Orc) die() {
	new(AHero).die()
}

func (p Orc) status() {
	ah := new(AHero)
	ah.life = p.life
	ah.status()
}

func (p Orc) shout() {
	ah := new(AHero)
	ah.name = p.name
	ah.shout()
}

func (p Orc) attack(oponent *Paladin) {
	fmt.Printf("%v\n", "Eat that human")
	oponent.life = oponent.life - p.power
	if oponent.life == 0 {
		oponent.die()
	}
}
