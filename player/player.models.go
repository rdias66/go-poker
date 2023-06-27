package player

import "go-poker/rules"

type Player struct {
    Hand []deck.Card
    FullHand []deck.Card
    Id string
    Name string
    HandRank HandRank
}

func (p Player) ShowHand() []deck.Card {
  return p.Hand
}
