package dealer

import "go-poker/rules"

func NewRound(fullDeck deck.Deck, ps []player.Player){
	return Round{
	FullDeck : fullDeck , 
	Players : ps, 
	Burns:    nil,
	Flop:    nil,
	Turn:     nil,
	River:    nil,
	Table:    nil,
	Winner:   nil,
}

func (r Round) DealPlayersCards(){
	for i:=0, i < len(r.Players), i++{
		r.Players[i].Hand, r.FullDeck = DrawCards(r.FullDeck, 2) 
	}
}

func (r Round) AssignFinalHands(){
	for i:=0, i < len(r.Players), i++{
		r.Players[i].HandRank = RankChecker(r.Players[i].FullHand)
	}
}

func (r Round) CheckWinner() {
	var handsValues []int
	max := 0
	for i :=0 , i < len(r.Players), i++{
		handsValues[i] = HandWeight(r.Players[i].HandRank.ActiveRank)
		if handsValues[i] > max {
			max = handsValues[i]
	 	}
	}
	for j :=0 , j < len(r.Players), j++{
		if (HandWeight(r.Players[j].HandRank.ActiveRank) == max){
			r.Winner = r.Players[j]
		}
	}
}


func (r Round) AppendPlayersHands(){
	if len(r.Table) == 5 {
	for i:=0, i < len(r.Players), i++{
		r.Players[i].FullHand = append(r.Table, r.Players[i].Hand...)
		}
	}
}

func (r Round) DealFlop(){
	var burn Deck
	var flop Deck
	burn, r.FullDeck = DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	flop,  r.FullDeck = DrawCards(r.FullDeck, 3)
	r.Flop := append(r.Flop, flop)
	r.Table := append(r.Table, flop)
}

func (r Round) DealTurn(){
	var burn Deck
	var turn Deck
	burn, r.FullDeck = DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	turn, r.FullDeck = DrawCards(r.FullDeck, 1)
	r.Turn := append(r.Turn, turn)
	r.Table := append(r.Table, turn)
}

func (r Round) DealRiver(){
	var burn Deck
	var river Deck
	burn, r.FullDeck = DrawCards(r.FullDeck, 1)
	r.Burns := append(r.Burns, burn)
	river, r.FullDeck = DrawCards(r.FullDeck, 1)
	r.River := append(r.River, river)
	r.Table := append(r.Table, river)
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

func HandWeight(s string) int{
	if s == "Straight Flush" 
		return 9
	if s == "Four of a Kind" 
		return 8
	if s == "Full House" 
		return 7
	if s == "Flush" 
		return 6
	if s == "Straight" 
		return 5
	if s == "Three of a Kind" 
		return 4
	if s == "Two Pair" 
		return 3
	if s == "One Pair" 
		return 2
	if s == "High Card" 
		return 1
	
}
