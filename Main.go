package main

func main() {

	player1 := new(Paladin)
	player1.AHero = AHero{8, "Pal the Great", 2, 0, 1, "Ready to fight for our king", "Paladin", "Take that, ugly beast!", ""}
	player2 := new(Orc)
	player2.AHero = AHero{4, "Ugly the Orc", 2, 0, 1, "I'm born! Ugh! Ugh! Ugh!", "Orc", "Eat that human", "Ugh, come here!"}

	player1.move()
	player2.move()
	player1.shout()
	player2.shout()
	player1.attack(&player2.AHero)
	player2.attack(&player1.AHero)
	player1.attack(&player2.AHero)
	player2.status()
	player1.attack(&player2.AHero)
	player1.attack(&player2.AHero)
}
