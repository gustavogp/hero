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
