// Package main defines the entry point for the Go program.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Suit is a custom type representing the different card suits.
type Suit int

// Enum for the four card suits.
const (
	Hearts Suit = iota + 1
	Spades
	Diamonds
	Clubs
)

// String returns the string representation of the Suit.
func (s Suit) String() string {
	return [...]string{"Hearts", "Spades", "Diamonds", "Clubs"}[s-1]
}

// Rank is a custom type representing the value of a card.
type Rank int

// Enum for the card values from Ace to King.
const (
	Ace Rank = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// String returns the string representation of the Rank.
func (r Rank) String() string {
	return [...]string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}[r]
}

// Card represents a single playing card, with both a suit and a value.
type Card struct {
	suit Suit
	rank Rank
}

// Deck is a custom type representing a map of Cards.
type Deck map[int]Card

// newDeck is a constructor function that creates and returns a new Deck.
// The deck is populated with cards for all combinations of four suits and thirteen values.
func newDeck() Deck {
	newDeck := make(Deck)
	suits := []Suit{Hearts, Spades, Diamonds, Clubs}
	ranks := []Rank{Ace, One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

	// Counter for assigning the key in the map
	counter := 0

	// Nested loops to generate each combination of value and suit
	for _, suit := range suits {
		for _, rank := range ranks {
			newDeck[counter] = Card{suit: suit, rank: rank}
			counter++
		}
	}

	return newDeck
}

// deal is a function that splits a deck into two parts based on the handSize:
// The first part is a hand of cards, and the second part is the remaining deck.
// Since Deck is now a map, we'll return two maps representing the hand and the remaining deck.
func deal(d Deck, handSize int) (Deck, Deck) {
	hand := make(Deck)
	remaining := make(Deck)

	counter := 0
	for key, card := range d {
		if counter < handSize {
			hand[key] = card
		} else {
			remaining[key] = card
		}
		counter++
	}

	return hand, remaining
}

// print is a receiver function that prints each card in the deck.
func (d Deck) print() {
	for key, card := range d {
		fmt.Printf("Card %d: %s of %s\n", key, card.rank.String(), card.suit.String())
	}
}

// toString is a receiver function that converts a Deck into a single string.
// The cards are joined together as a comma-separated string.
func (d Deck) toString() string {
	cardStrings := []string{}
	for _, card := range d {
		cardStrings = append(cardStrings, fmt.Sprintf("%s of %s", card.rank.String(), card.suit.String()))
	}
	return strings.Join(cardStrings, ",")
}

// saveToFile saves the current Deck to a file.
// The file is created with the specified fileName, and the deck's string representation is written into it.
func (d Deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

// newDeckFromFile reads a Deck from a file and recreates it.
// It reads the file's content, converts it to a string, and then splits it to recreate the Deck.
func newDeckFromFile(filename string) Deck {
	deckInBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error:", err)
	}
	stringifiedDeck := strings.Split(string(deckInBytes), ",")
	cards := make(Deck)
	counter := 0
	for _, cardStr := range stringifiedDeck {
		cardDetails := strings.Split(cardStr, " of ")
		cards[counter] = Card{
			rank: parseCardValue(cardDetails[0]),
			suit: parseSuit(cardDetails[1]),
		}
		counter++
	}
	return cards
}

// Helper function to parse card value from string to Rank.
func parseCardValue(valueStr string) Rank {
	switch valueStr {
	case "One":
		return One
	case "Two":
		return Two
	case "Three":
		return Three
	case "Four":
		return Four
	case "Five":
		return Five
	case "Six":
		return Six
	case "Seven":
		return Seven
	case "Eight":
		return Eight
	case "Nine":
		return Nine
	case "Ten":
		return Ten
	case "Jack":
		return Jack
	case "King":
		return King
	case "Queen":
		return Queen
	default:
		return Ace // default fallback
	}

}

// Helper function to parse suit from string to Suit.
func parseSuit(suitStr string) Suit {
	switch suitStr {
	case "Hearts":
		return Hearts
	case "Spades":
		return Spades
	case "Diamonds":
		return Diamonds
	default:
		return Clubs // default fallback
	}
}

// shuffle randomizes the order of the cards in the Deck.
// Since Deck is a map, we will randomize the keys and shuffle the corresponding values.
func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	keys := make([]int, 0, len(d))

	// Get all the keys from the deck
	for key := range d {
		keys = append(keys, key)
	}

	// Shuffle the keys and remap the cards
	for i := range keys {
		newPosition := r.Intn(len(keys))
		keys[i], keys[newPosition] = keys[newPosition], keys[i]
	}

	shuffledDeck := make(Deck)
	for i, key := range keys {
		shuffledDeck[i] = d[key]
	}

	// Copy back the shuffled deck into the original deck
	for key, card := range shuffledDeck {
		d[key] = card
	}
}
