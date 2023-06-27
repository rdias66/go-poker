package dealer

func NewRound(fullDeck deck.Deck, ps []player.Player){
	return Round{FullDeck : fullDeck , Players : ps}
}

func (r Round) DealFlop(){
	burn, deck := DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	flop, deck := DrawCards(r.FullDeck, 3)
	r.Flop := append(r.Flop, flop)
}

func (r Round) DealTurn(){
	burn, deck := DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	turn, deck := DrawCards(r.FullDeck, 1)
	r.Turn := append(r.Turn, turn)
}

func (r Round) DealRiver(){
	burn, deck := DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	river, deck := DrawCards(r.FullDeck, 1)
	r.River := append(r.River, river)
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
