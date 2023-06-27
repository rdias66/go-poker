package player

type Player struct {
    Hand []deck.Card
    FullHand []deck.Card
    Id string
    Name string
    Result *FinalHand
}

type FinalHand struct {
    Hand[]deck.Card
    HandResult string
}

func (p Player) ShowHand() []deck.Card {
  return p.Hand
}
