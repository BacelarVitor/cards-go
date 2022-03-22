package main

import (
	"fmt"
	"strings"
)
type deck []string;

func newDeck() deck {
  var cards deck
  cardSuits := [4]string{"Spades", "Hearts", "Clubs", "Diamonds"}
  cardValues := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

  for _, suit := range cardSuits {  
    for _, value := range cardValues {
      cards = append(cards, value + " of " + suit)
    }
  }
  return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, hansSize int) (deck, deck) {
  return d[:hansSize], d[hansSize:]
}

func (d deck) toString() string {
  return strings.Join([]string(d), ",")
}