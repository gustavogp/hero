package main

import (
	"fmt"
)

type AHero struct {
	life       int
	name       string
	speed      int
	position   int
	power      int
	createYell string
	heroType   string
	attackYell string
	moveYell   string
}

func (p AHero) move() {
	p.position = p.position + p.speed
	if p.heroType == "Orc" {
		fmt.Println(p.moveYell)
	}
}

func (p AHero) die() {
	fmt.Printf("%v\n", "Nooooo!!!")
}

func (p AHero) status() {
	fmt.Printf("Life Remainnig: %v\n", p.life)
}

func (p AHero) shout() {
	fmt.Printf("I am %v\n", p.name)
}

func (p AHero) attack(oponent *AHero) {
	fmt.Printf("%v\n", p.attackYell)
	oponent.life = oponent.life - p.power
	if oponent.life == 0 {
		oponent.die()
	}
}
