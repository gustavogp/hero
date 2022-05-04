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

func (p Paladin) attack(oponentAH *AHero) {
	fmt.Printf("%v\n", "Take that, ugly beast!")
	p.AHero.attack(oponentAH)
}
