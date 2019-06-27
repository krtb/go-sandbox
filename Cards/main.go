package main

import "fmt"

func main() {
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of spades")
	//append returns a new array, does not modify cards() slice

	cards.print()

	fmt.Println(cards)
}
