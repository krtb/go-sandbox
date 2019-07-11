package main

import "fmt"

// create a new type of deck, which is a slice of strings

type deck []string

//to create a new deck, don't need a receiver
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//replace variables in Go that you're not going to use with an underscore.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"of"+suit)
		}
	}
	return cards
}

//CREATE NEW TYPE OF DECK WHICH IS SLICE OF STRINGS
// type of deck is borrowing all of 'string' behavior
// will store all logic in this file

//before our function name, add in a RECEIEVER
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//can call the deal function with the first arg of type deck, and the second arg of type int
// go has support for returning multiple values from one function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // two seperate return values inside of one value

}
