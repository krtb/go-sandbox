package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	//need to inform the go compiler that when this func is run, return a type of string
	return "five of diamonds"
}
