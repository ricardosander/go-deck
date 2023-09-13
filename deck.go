package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	var cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	var cardsValues = []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
