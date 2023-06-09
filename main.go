package main

import (
	"go-poker/deck"
)

func main() {
	clubs := deck.SuitBuilder("Clubs")
	diamonds := deck.SuitBuilder("Diamonds")
	hearts := deck.SuitBuilder("Hearts")
	spades := deck.SuitBuilder("Spades")

	suits := []deck.Suit{clubs, diamonds, hearts, spades}

	deck := deck.DeckBuilder(suits)

	deck.Printer()

	deck.Shuffle()

	deck.Printer()

}
