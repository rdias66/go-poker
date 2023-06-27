package dealer

import (
	"go-poker/deck"
	"go-poker/player"
)

type Round struct {
	FullDeck    deck.Deck
	Players	    []player.Player
	Burns       *[]deck.Card
	Flop        *[]deck.Card
	Turn        *deck.Card
	River       *deck.Card
	Table 	    *[]deck.Card
	Winner 	    *player.Player
}


