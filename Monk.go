package main

import (
	"fmt"
)

type Monk struct {
	*AHero
}

func init() {
	fmt.Printf("%v\n", "Monks bring peace")
}

func (p Monk) attack(oponentAH *AHero) {
	fmt.Printf("%v\n", "Do you want some coffe?")
	p.AHero.attack(oponentAH)
}
