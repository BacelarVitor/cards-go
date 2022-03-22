package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 2)

	hand.print()
	remainingCards.print()
}