package main

import "fmt"

// create a new type of deck, which is a slice of strings

type deck []string

//CREATE NEW TYPE OF DECK WHICH IS SLICE OF STRINGS
// type of deck is borrowing all of 'string' behavior
// will store all logic in this file

//before our function name, add in a RECEIEVER
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
