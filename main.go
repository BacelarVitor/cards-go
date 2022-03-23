package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("flin")

	cards := newDeckFromFile("flin")
	cards.print()
}