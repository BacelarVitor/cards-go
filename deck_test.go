package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 deck length of 52, but got %v", len(d))
	}
}

func TestFirstCardIsAceOfSpades(t *testing.T) {
	d := newDeck()

	e := "Ace of Spades"
	testCardPosition(t, d, 0, e)
}


func TestLastCardIsAceOfSpades(t *testing.T) {
	d := newDeck()
	
	i := len(d) -1
	e := "King of Diamonds"
	testCardPosition(t, d, i, e)
}

func testCardPosition(t *testing.T, d deck, i int, e string) {
	if d[i] !=  e {
		t.Errorf("Expected first card of %v, but got %v", e, d[i])
	}
}