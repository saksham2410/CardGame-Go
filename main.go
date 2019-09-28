package main

import "fmt"

func main() {
	// var card string = "Ace of cards"

	// cards := "Ace of cards"
	cards := []string{"Meow", newCard()}
	cards = append(cards, "saksham")
	//only define colon when initially you have to define a value
	// card = "Meow"
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// fmt.Println(card)
}
func newCard() string {
	return "Meow modified"
}
