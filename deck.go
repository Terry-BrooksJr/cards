package main

import "fmt"

// Type -

type Deck []string

func newDeck() Deck {
	cards := Deck{}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
func (d Deck) print() {
	for idx, card := range d {
		fmt.Println(idx, card)
	}
}
