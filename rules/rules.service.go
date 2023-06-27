package rules

import "go-poker/deck"

type 

type HandRank struct {
    Hand Deck
    ActiveRank string
}

var Ranks = []string{
  "Royal Flush",
	"Straight Flush",
	"Four of a Kind",
	"Full House",
	"Flush",
	"Straight",
	"Three of a Kind",
	"Two Pair",
	"One Pair",
	"High Card",
}

func RankChecker(h Deck) HandRank {
  if len(h) != 5 {
    return HandRank{Hand: h, ActiveRank: "Invalid hand size"}
  }
    
}
