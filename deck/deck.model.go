package deck

import "fmt"

type card struct {
	value string
	suit  string
}

type suit []card

type deck []card
