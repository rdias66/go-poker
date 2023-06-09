package dealer

import (
	"go-poker/deck"
)

type Round struct {
	FullDeck    deck.Deck
	PlayerCount int
	Burns       []deck.Card
	Flop        []deck.Card
	Turn        deck.Card
	River       deck.Card
}

func Deal(d deck.Deck, playerCount int) {
	//dealing function code here
}
