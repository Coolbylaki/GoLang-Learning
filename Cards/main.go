package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "King of Spades")

	for i := 0; i < len(cards); i++ {
		fmt.Println(cards[i])
	}
}

func newCard() string {
	return "Ace of Spades"
}
