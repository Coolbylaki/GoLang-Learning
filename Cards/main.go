package main

func main() {
	// Make the deck
	deck := newDeck()

	// Deal the deck
	// hand, _ := deal(5, deck)

	// Write to file
	// hand.writeToFile("currentHand")

	// Read deck from file
	// newDeck := newDeckFromFile("currentHand")

	// Shuffle the deck
	deck.shuffle()

	deck.print()
}
