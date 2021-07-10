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

func (p Orc) implementsIHero() {}

func (p Orc) move() {
	p.position = p.position + p.speed
	fmt.Println("Ugh, come here!")
}

func (p Orc) die() {
	fmt.Printf("%v\n", "Nooooo!!!")
}

func (p Orc) status() {
	fmt.Printf("Life Remainnig: %v\n", p.life)
}

func (p Orc) shout() {
	fmt.Printf("I am %v\n", p.name)
}

func (p Orc) attack(oponent *Paladin) {
	fmt.Printf("%v\n", "Eat that human")
	oponent.life = oponent.life - p.power
	if oponent.life == 0 {
		oponent.die()
	}
}
