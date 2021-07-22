package main

import (
	"fmt"
)

type Paladin struct {
	AHero
}

func init() {
	fmt.Printf("%v\n", "Ready to fight for our king")
}

func (p Paladin) move() {
	ah := new(AHero)
	ah.position, ah.speed = p.position, p.speed
	ah.move()
}

func (p Paladin) die() {
	new(AHero).die()
}

func (p Paladin) status() {
	ah := new(AHero)
	ah.life = p.life
	ah.status()
}

func (p Paladin) shout() {
	ah := new(AHero)
	ah.name = p.name
	ah.shout()
}

func (p Paladin) attack(oponent *Orc) {
	fmt.Printf("%v\n", "Take that, ugly beast!")
	oponent.life = oponent.life - p.power
	if oponent.life == 0 {
		oponent.die()
	}
}
