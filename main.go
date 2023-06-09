package main

import (
      "go-poker/deck"
      "fmt"
)

func main() {
  clubs := SuitBuilder("Clubs")
	diamonds := SuitBuilder("Diamonds")
	hearts := SuitBuilder("Hearts")
	spades := SuitBuilder("Spades")
  
	suits := []suit{clubs, diamonds, hearts, spades}
  
	deck := DeckBuilder(suits)
  
  deck.Printer()
  
  deck = deck.Shuffle()
  
  deck.Printer()
  
}