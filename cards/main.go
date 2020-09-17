package main

func main() {
	cards := newDeck()

	// hand, remainingCards := cards.deal(5)

	cards.saveToFile("My cards!")

	
}
