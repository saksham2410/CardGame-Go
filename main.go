package main

func main() {
	// var card string = "Ace of cards"

	// cards := "Ace of cards"
	cards := deck{"Meow", newCard()}
	cards = append(cards, "saksham")
	//only define colon when initially you have to define a value
	// card = "Meow"
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards.print()
	// fmt.Println(card)
}
func newCard() string {
	return "Meow modified"
}
