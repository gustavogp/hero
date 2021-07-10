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

func (p Paladin) implementsIHero() {}

func (p Paladin) move() {
	p.position = p.position + p.speed
}

func (p Paladin) die() {
	fmt.Printf("%v\n", "Nooooo!!!")
}

func (p Paladin) status() {
	fmt.Printf("Life Remainnig: %v\n", p.life)
}

func (p Paladin) shout() {
	fmt.Printf("I am %v\n", p.name)
}

func (p Paladin) attack(oponent *Orc) {
	fmt.Printf("%v\n", "Take that, ugly beast!")
	oponent.life = oponent.life - p.power
	if oponent.life == 0 {
		oponent.die()
	}
}
