package main

// In the main function
// Create a cards variable that contains a slice
// slice is an array that can grow or shrink
// deck is alias of []string created in the deck.go file
// We also append Six of Spades to the end of cards
// This does not add to the existing cards, but creates a new slice of cards
func main() {
	cards := newDeck()

	// Using the variable cards
	// Pass cards to the receiver function print in the deck.go file
	cards.print()

}
