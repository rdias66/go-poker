package rules

import "go-poker/deck"

type HandRank struct { // MOVE TO RULES.MODELS.GO
    Hand Deck
    ActiveRank string
}

func RankChecker(h Deck) HandRank {
  if len(h) != 5 {
    return HandRank{Hand: h, ActiveRank: "Invalid hand size"}
  }
	if isRoyalFlush(h){
		return HandRank{Hand: h, ActiveRank: "Straight Flush"}
	}
	if isFourOfAKind(h){
		return HandRank{Hand: h, ActiveRank: "Four of a Kind"}
	}
	if isFullHouse(h){
		return HandRank{Hand: h, ActiveRank: "Full House"}
	}
	if isFlush(h){
		return HandRank{Hand: h, ActiveRank: "Full House"}
	}
	if isStraight(h){
		return HandRank{Hand: h, ActiveRank: "Full House"}
	}
	if isThreeOfAKind(h){
		return HandRank{Hand: h, ActiveRank: "Full House"}
	}
	OnePair, TwoPair : = isPair(h)
	if TwoPair{
		return HandRank{Hand: h, ActiveRank: "Two Pair"}
	}
	if OnePair{
		return HandRank{Hand: h, ActiveRank: "One Pair "}
	}
	else {
		return HandRank{Hand: h, ActiveRank: "High Card"}
	}
}

func CardValueCounter(h Deck) map[string]int{ 
  counts := make(map[string]int)
	for _, card := range h {
		counts[card.value]++
	}
	return counts
}

func CardSuitCounter(h Deck) map[string]int{ 
  counts := make(map[string]int)
	for _, card := range h {
		counts[card.suit]++
	}
	return counts
}

func isRoyalFlush(h Deck) bool {
	if isStraight(h) && isFlush(h){
		return true
	}
	return false
}

func isFourOfAKind(h Deck) bool {
	counts := CardValueCounter(h)
	for _, value := range counts {
		if counts[value] == 4{
			return true
		}
	}
	return false
}

func isFullHouse(h Deck) bool {
	if isThreeOfAKind(h) && isPair(h) {
		return true
	}
	return false
}

func isFlush(h Deck) bool {
	counts := CardSuitCounter(h)
	for _, value := range counts {
		if counts[value] == 5{
			return true
		}
	}
	return false
}

func isStraight(h Deck) bool{
	for i := 1; i < len(h); i++ {
		if cardValue(h[i]) != cardValue(h[i-1])+1 {
			return false
		}
	}
	return true
}

func isThreeOfAKind(h Deck) bool {
	counts := CardValueCounter(h)
	for _, value := range counts {
		if counts[value] == 3{
			return true
		}
	}
	return false
}

func isPair(h Deck) (bool, bool) { // could be 2 pairs
	counts := CardValueCounter(h)
	var pairCounter int 
	for _, value := range counts {
		if counts[value] == 2{
			pairCounter++
		}
	}
	if pairCounter == 1 {
		return true, false
	}
	if pairCounter == 2 {
		return false, true
	}
	return false, false
}