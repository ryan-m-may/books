package main

import (
	"fmt"
)

func main() {
	var card string = newCard()
	fmt.Printf("%s", card)
}

func newCard() string {
	return "Ace of Spades"
}
