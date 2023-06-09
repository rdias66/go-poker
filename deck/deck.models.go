package deck

type Card struct {
	value string
	suit  string
}

type Suit []Card

type Deck []Card
