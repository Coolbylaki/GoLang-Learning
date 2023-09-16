package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(handSize int, d deck) (deck, deck) {
	hand := d[:handSize]
	remainingHand := d[handSize:]

	return hand, remainingHand
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
