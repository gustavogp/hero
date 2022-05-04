package main

import "fmt"

type AHero struct {
	life     int
	name     string
	speed    int
	position int
	power    int
}

func (ah *AHero) implementsIHero() {}

func (ah *AHero) move() {
	ah.position = ah.position + ah.speed
}

func (ah *AHero) die() {
	fmt.Printf("%v\n", "Nooooo, I'm d...!!!")
}

func (ah *AHero) status() {
	fmt.Printf("Life Remainnig: %v\n", ah.life)
}

func (ah *AHero) shout() {
	fmt.Printf("I am %v\n", ah.name)
}

func (ah *AHero) attack(oponentAH *AHero) {
	oponentAH.life = oponentAH.life - ah.power
	if oponentAH.life == 0 {
		oponentAH.die()
	}
}
