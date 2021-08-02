package main

func main() {
	cards := newDeck()
	cards.shuffle()

	// cards.saveToFile(("myCards.txt"))


	cards.print()
}

// func newCard() string {
// 	return "Ace of Diamonds"
// }