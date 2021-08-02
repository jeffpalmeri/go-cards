package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card is not Ace of Spades... but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("First card is not Four of Club... but got %v", d[len(d) - 1])
	}
}

func TestSaveNewDeckAndTestNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	deck := newDeck()
	deck.saveToFile("_decktesting.txt")

	loadedDeck := newDeckFromFile("_decktesting.txt")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, instead got %v", len(loadedDeck))
	}

	os.Remove("_decktesting.txt")
}