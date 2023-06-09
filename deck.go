package deck

import "fmt"

type card struct {
	value string
	suit  string
}

type suit []card

type deck []card

func SuitBuilder(suitName string) suit {
	newSuit := suit{}
	for i := 1; i <= 13; i++ {
		if i == 1 {
			var newCard card = NewCard("Ace of ", suitName)
			newSuit = append(newSuit, newCard)
		} else if i == 11 {
			var newCard card = NewCard("Jack of ", suitName)
			newSuit = append(newSuit, newCard)
		} else if i == 12 {
			var newCard card = NewCard("Queen of ", suitName)
			newSuit = append(newSuit, newCard)
		} else if i == 13 {
			var newCard card = NewCard("King of ", suitName)
			newSuit = append(newSuit, newCard)
		} else {
			var newCard card = NewCard(fmt.Sprint(i), suitName)
			newSuit = append(newSuit, newCard)
		}
	}
	return newSuit
}

func NewCard(value string, suit string) card {
	return card{value, suit}
}

func DeckBuilder(suits []suit) deck {
	newDeck := deck{}
	for _, s := range suits {
		for _, c := range s {
			newDeck = append(newDeck, c)
		}
	}
	return newDeck
}

func (deck deck) DeckPrinter() {
	for _, d := range deck {
		fmt.Println(d)
	}
}
