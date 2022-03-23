package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) saveToFile(filename string) error {
  return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
  bs, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }

  return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
  rand.Seed(time.Now().UnixNano())
  rand.Shuffle(len(d), func(i, j int) { 
    d[i], d[j] = d[j], d[i] 
  })
}