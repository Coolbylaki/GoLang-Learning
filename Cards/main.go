package main

import "fmt"

func main() {
	// Make the deck
	deck := newDeck()

	// Deal the deck
	hand, remainingDeck := deal(5, deck)

	hand.print()
	fmt.Println("________________________")
	remainingDeck.print()
}
