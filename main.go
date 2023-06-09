package main

import (
	"go-poker/deck"
	"go-poker/player"
	"go-poker/dealer"
	"go-poker/rules"
	"fmt"
)

func main() {
	//DECK BUILDING SECTION 
	clubs := deck.SuitBuilder("Clubs")
	diamonds := deck.SuitBuilder("Diamonds")
	hearts := deck.SuitBuilder("Hearts")
	spades := deck.SuitBuilder("Spades")
	suits := []deck.Suit{clubs, diamonds, hearts, spades}
	deck := deck.DeckBuilder(suits)
	deck.Shuffle()
	deck.Printer()

	//PLAYERS CREATION
	player1 := Player{[]deck.Card, Id: '1', Name: 'Player1'}
	player2 := Player{[]deck.Card, Id: '2', Name: 'Player2'}
	players := []Player{player1, player2}

	//ROUND CREATION
	round1 := NewRound(deck, players)

	//DEAL CARDS
	round1.DealPlayersCards()

	//GAMEFLOW
	
	//FIRST BETS - hypothetical
	//FLOP
	round1.DealFlop()
	
	//POSTFLOP BETS - hypothetical
	//TURN
	round1.DealTurn()

	//POSTTURN BETS - hypothetical
	//RIVER
	round1.DealRiver()

	//CHECK HANDS
	round1.AppendPlayersHands()
	round1.AssignFinalHands()
	round1.CheckWinner()

	//SHOW INFO ON ROUND
	fmt.Printf(round1)

}
