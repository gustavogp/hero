package main

func main() {

	// checking that types implement interface
	var _ IHero = (*AHero)(nil)
	var _ IHero = Paladin{}
	var _ IHero = Orc{}

	var player1 IHero = Paladin{&AHero{8, "Pal the Great", 2, 0, 1}}
	var player2 IHero = Orc{&AHero{4, "Ugly the Orc", 2, 0, 1}}

	player1.move()
	player2.move()
	player1.shout()
	player2.shout()
	player1.attack(player2.(Orc).AHero)
	player2.attack(player1.(Paladin).AHero)
	player1.attack(player2.(Orc).AHero)
	player2.status()
	player1.attack(player2.(Orc).AHero)
	player1.attack(player2.(Orc).AHero)

}
