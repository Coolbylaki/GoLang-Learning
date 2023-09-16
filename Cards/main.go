package main

func main() {
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "King of Spades")

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
