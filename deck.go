package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// new type deck

type deck []string

func getNewDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) top() string {
	d = d[1:]
	return d[0]
}

func deal(d deck, handSize int) (deck, deck) {
	// if len(d) < handSize {
	return d[:handSize], d[handSize:]
	// }
}

func (d deck) convertToByteSlice() []byte {

	var compiledString string = strings.Join([]string(d), ",")
	// for _, card := range d {
	// 	compiledString = compiledString + "\n" + card
	// }
	byteSlice := []byte(compiledString)

	return byteSlice
}

func (d deck) saveToFile(filename string) error {
	data := d.convertToByteSlice()
	return ioutil.WriteFile(filename, data, 0644)
}

func readDeckFromFile(filename string) deck {
	bs, er := ioutil.ReadFile(filename)
	if er != nil {
		fmt.Println("Error:", er)
		os.Exit(1)
	}
	stringVal := string(bs)
	stringSlice := strings.Split(stringVal, ",")
	return deck(stringSlice)
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		randomIndex := rand.Intn(51)
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}
