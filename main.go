package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.print()

	cards, hand := deal(cards, 3)

	fmt.Println("\nMy Hand:")
	hand.print()

	fmt.Println("\nRemaining deck")
	cards.print()

	cards.saveToFile("test")
}
