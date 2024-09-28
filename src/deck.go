// Package main defines the entry point for the Go program.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
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
	suit  Suit
	rank  Rank
	owner player
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

// Print is a receiver function that prints each card in the deck.
func (d Deck) Print() {
	for key, card := range d {
		fmt.Printf("Card %d: %s of %s\n", key, card.rank.String(), card.suit.String())
	}
}

// PrintToWeb is a receiver function that prints each card in the deck designed for to the required parameters of a web request.
func (d Deck) PrintToWeb(w http.ResponseWriter, r *http.Request) {
	for key, card := range d {
		fmt.Fprintf(w, "Card %d: %s of %s\n", key, card.rank.String(), card.suit.String())
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

// NewDeckFromFile reads a Deck from a file and recreates it.
// It reads the file's content, converts it to a string, and then splits it to recreate the Deck.
func NewDeckFromFile(filename string) Deck {
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
func deletePlayer(slice []player, index int) []player {
	return append(slice[:index], slice[index+1:]...)
}

func deleteHand(slice []Hand, index int) []Hand {
	return append(slice[:index], slice[index+1:]...)
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

type Hand struct {
	player    player
	handCards [13]Card
}

func (h Hand) ShowHand() {
	for key, card := range h.handCards {
		fmt.Printf("Card %d: %s of %s\n", key, card.rank.String(), card.suit.String())
	}
}

// deal is a function that splits a deck into two parts based on the handSize:
// The first part is a hand of cards, and the second part is the remaining deck.
// Since Deck is now a map, we'll return two maps representing the hand and the remaining deck.
func (d Deck, P1 player, P2 player, P3 player, P4 player) deal(Hand, Hand, Hand, Hand) {
	var H1 Hand
	var H2 Hand
	var H3 Hand
	var H4 Hand
	handowners := []player{P1, P2, P3, P4}
	availableHands := []Hand{H1, H2, H3, H4}
	for len(handowners) > 0 {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(handowners))
		for len(availableHands) > 0 {
			availableHands[randomIndex].player = *handowners[randomIndex].cards
			handowners = deletePlayer(handowners, randomIndex)
			availableHands = deleteHand(availableHands, randomIndex)
		}
	}
	counter := 52
	for _, card := range d {
		if counter != 0 {
			for i := 0; i <= 52; i++ {
				if H1.handCards[12].rank.String() == " " {
					H1.handCards[i] = card
				} else if H1.handCards[12].rank.String() == " " {
					if i >= 12 && i != 12 {
						i = 0
					}
					H2.handCards[i] = card
				} else if H3.handCards[12].rank.String() == " " {
					if i >= 12 && i != 12 {
						i = 0
					}
					H3.handCards[i] = card
				} else if H4.handCards[12].rank.String() == " " {
					if i >= 12 && i != 12 {
						i = 0
					}
					H4.handCards[i] = card
				}

				counter--
			}

		}
	}
	return H1, H2, H3, H4

}
