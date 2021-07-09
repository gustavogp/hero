package main

type IHero interface {
	move()
	die()
	status()
	shout()
	attack(oponent *AHero)
}
