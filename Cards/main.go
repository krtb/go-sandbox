package main

import "fmt"

func main() {
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of spades")
	//append returns a new array, does not modify cards() slice

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	//need to inform the go compiler that when this func is run, return a type of string
	return "five of diamonds"
}
