package deck

import (
	"fmt"
	"math/rand"
	"time"
)

func SuitBuilder(suitName string) Suit {
	newSuit := Suit{}
	for i := 1; i <= 13; i++ {
		if i == 1 {
			var newCard Card = NewCard("Ace of ", suitName)
			newSuit = append(newSuit, newCard)
		} else if i == 11 {
			var newCard Card = NewCard("Jack of ", suitName)
			newSuit = append(newSuit, newCard)
		} else if i == 12 {
			var newCard Card = NewCard("Queen of ", suitName)
			newSuit = append(newSuit, newCard)
		} else if i == 13 {
			var newCard Card = NewCard("King of ", suitName)
			newSuit = append(newSuit, newCard)
		} else {
			var newCard Card = NewCard(fmt.Sprint(i), suitName)
			newSuit = append(newSuit, newCard)
		}
	}
	return newSuit
}

func NewCard(value string, suit string) Card {
	return Card{value, suit}
}

func DeckBuilder(suits []Suit) Deck {
	newDeck := Deck{}
	for _, s := range suits {
		for _, c := range s {
			newDeck = append(newDeck, c)
		}
	}
	return newDeck
}

func (deck Deck) Printer() {
	for _, d := range deck {
		fmt.Print(d)
		fmt.Print(", ")
	}
}

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}
