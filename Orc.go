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

func (o Orc) attack(oponentAH *AHero) {
	fmt.Printf("%v\n", "Eat that human")
	o.AHero.attack(oponentAH)
}
