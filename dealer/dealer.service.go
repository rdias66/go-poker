package dealer

func NewRound(fullDeck deck.Deck, ps []player.Player){
	return Round{FullDeck : fullDeck , Players : ps}
}

func (r Round) DealPlayersCards(){
	for i:=0, i < len(r.Players), i++{
		r.Players[i].Hand, r.FullDeck = DrawCards(r.FullDeck, 2) 
	}
}

func (r Round) DealFlop(){
	var burn Deck
	var flop Deck
	burn, r.FullDeck = DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	flop,  r.FullDeck = DrawCards(r.FullDeck, 3)
	r.Flop := append(r.Flop, flop)
}

func (r Round) DealTurn(){
	var burn Deck
	var turn Deck
	burn, r.FullDeck = DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	turn, r.FullDeck = DrawCards(r.FullDeck, 1)
	r.Turn := append(r.Turn, turn)
}

func (r Round) DealRiver(){
	var burn Deck
	var river Deck
	burn, r.FullDeck = DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	river, r.FullDeck = DrawCards(r.FullDeck, 1)
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
