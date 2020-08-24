package main

import "fmt"

func main() {
	// var card string = "King of Hearts"
	cards := []string{"Queen of Spades", getNewCard()}
	cards = append(cards, "Seven of Diamond")
	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func getNewCard() string {
	return "Two of Hearts"
}
