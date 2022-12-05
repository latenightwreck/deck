package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.shuffle()
	cards.print()

}
