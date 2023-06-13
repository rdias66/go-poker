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
	Winner 	    *player.Player
}

func newRound(fullDeck deck.Deck, ps []player.Player){
	return Round{FullDeck : fullDeck , Players : ps}
}

func  (Round) DealCards(d deck.Deck, ps []player.Player){
	for _, p := range ps {
		p.Hand, d := DrawCards(d, 2)
	}
	r := newRound(d, ps)
	return r
}

func (r Round) DealFlop(r Round) {
	burn, deck := DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	flop, deck := DrawCards(r.FullDeck, 3)
	r.Flop := append(r.Flop, flop)
	return r
}

func (r Round) DealTurn(r Round) {
	burn, deck := DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	turn, deck := DrawCards(r.FullDeck, 1)
	r.Turn := append(r.Turn, turn)
	return r
}

func (r Round) DealRiver(r Round) {
	burn, deck := DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	river, deck := DrawCards(r.FullDeck, 1)
	r.River := append(r.River, river)
	return r
}

func DrawCards(d deck.Deck, numCards int) (deck.Deck, deck.Deck){
	deckSize := len(d)
	if numCards > deckSize {
		numCards = deckSize // If the requested number of cards exceeds the deck size, draw all the remaining cards
	}

	draw := deck[deckSize-numCards:] // Extract the last 'numCards' from the deck
	deck = deck[:deckSize-numCards]  // Remove the drawn cards from the deck

	return draw, deck
}
