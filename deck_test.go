package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected NewDeck lenth of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first element to be Ace of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"
	_ = os.Remove(testFileName)

	deck := newDeck()
	err := deck.saveToFile(testFileName)
	if err != nil {
		t.Errorf("Error when saving to file: %v", err)
	}

	deck = newDeckFromFile(testFileName)
	if len(deck) != 52 {
		t.Errorf("Expected newDeckFromFile lenth of 52, but got %v", len(deck))
	}

	_ = os.Remove(testFileName)
}
