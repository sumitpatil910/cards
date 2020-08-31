package main

func main() {
	myDeck := getNewDeck()
	myDeck.shuffle()
	myDeck.print()

	// hand, myDeck := deal(myDeck, 5)
	// hand.print()
	// myDeck.print()
	// myDeck.saveToFile("MyNewDeck.txt")
	// newDeck := readDeckFromFile("MyDeck.txt")

	// newDeck.print()
}
