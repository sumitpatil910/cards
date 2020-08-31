package main

import (
	"testing"
)

func TestGetNewDeck(t *testing.T) {
	d := getNewDeck()
	// fmt.Println(len(d))
	if len(d) != 52 {
		t.Errorf("Expected length of 52 but got %v", len(d))
	}
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected 'Ace of Hearts' but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {

		t.Errorf("Expected 'King of Clubs' but got %v", d[0])
	}
}
