package main

func main() {

	var player1 = &Paladin{&AHero{8, "Pal the Great", 2, 0, 1}}
	var player2 = &Orc{&AHero{4, "Ugly the Orc", 2, 0, 1}}

	player1.move()
	player2.move()
	player1.shout()
	player2.shout()
	player1.attack(player2)
	player2.attack(player1)
	player1.attack(player2)
	player2.status()
	player1.attack(player2)
	player1.attack(player2)
}
