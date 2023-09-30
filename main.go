package main

import (
	"fmt"
	"os"
)

func main() {
	cards := newDeck()
	const fileName = "my_cards"
	err := cards.saveToFile(fileName)
	if err != nil {
		fmt.Print("Error:", err)
		os.Exit(1)
	}
	cards = newDeckFromFile(fileName)
	cards.print()
}
