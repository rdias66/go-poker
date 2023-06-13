package player

type Player struct {
    Hand []deck.Card
    Id string
}

func (p Player) ShowHand() []deck.Card {
  return p.Hand
}
