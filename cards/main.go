package main

func main() {
	cards := newDeck()

	// hand, remainingCards := cards.deal(5)

	// cards.saveToFile("my_cards")

	// newDeckFromFile("my_cards").print()
	cards.shuffle()

	cards.print()
}
