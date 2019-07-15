package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

type Suit string

const (
	HEARTS Suit = "♡"
	DIAMONDS Suit = "♢"
	CLUBS Suit = "♣"
	SPADES Suit = "♠"
)

type Card struct {
	number int
	suit Suit
}

func (card1 *Card) hasSameNumber(card2 Card) bool{
	return card1.number == card2.number
}

func (card1 *Card) hasSameSuit(card2 Card) bool{
	return card1.suit == card2.suit
}

func (card1 *Card) String() string{
	var cardName string = strconv.Itoa(card1.number)
	var cardSuit string = string(card1.suit)

	if card1.number == 1 {
		cardName = "A"
	} else if card1.number == 11 {
		cardName = "J"
	} else if card1.number == 12 {
		cardName = "Q"
	} else if card1.number == 13 {
		cardName = "K"
	}
	return cardName + cardSuit

}

func printCards(cards []Card) {
	for _, card := range cards {
		fmt.Print(card.String(), " ")
	}
	fmt.Println()
}

func getCardDeck() []Card{
	var cardDeck []Card

	for i := 1 ; i <= 13 ; i++ {
		cardDeck = append(cardDeck, Card{suit: HEARTS, number: i})
		cardDeck = append(cardDeck, Card{suit: SPADES, number: i})
		cardDeck = append(cardDeck, Card{suit: CLUBS, number: i})
		cardDeck = append(cardDeck, Card{suit: DIAMONDS, number: i})
	}
	return cardDeck
}

func shuffleCards(vals []Card) []Card {
  r := rand.New(rand.NewSource(time.Now().Unix()))
  ret := make([]Card, len(vals))
  perm := r.Perm(len(vals))
  for i, randIndex := range perm {
    ret[i] = vals[randIndex]
  }
  return ret
}



func main() {
	mycard := Card{suit: HEARTS, number:10}
	fmt.Println(mycard.number == 10)

	mycard2 := Card{suit: SPADES, number:10}
	fmt.Println(mycard.hasSameNumber(mycard2) == true)

	mycard3 := Card{suit: HEARTS, number:12}
	fmt.Println(mycard2.hasSameNumber(mycard3) == false)

	fmt.Println(mycard3.hasSameSuit(mycard) == true)
	fmt.Println(mycard3.hasSameSuit(mycard2) == false)

	fmt.Println(mycard3.String())

	printCards([]Card{mycard, mycard2, mycard3})

	var totalCards []Card = getCardDeck()
	printCards(totalCards)
	totalCards = shuffleCards(totalCards)
	totalCards = shuffleCards(totalCards)
	printCards(totalCards)



}

