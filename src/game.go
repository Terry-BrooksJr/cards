package main

import (
	"fmt"
)

type player struct {
	name string
	cards Hand
	score int
}

type matchScore struct {
	humanPlayerScore int
	computerPlayer1Score int
	computerPlayer2Score int
	computerPlayer3Score int
}

type gameMatch struct {
	scores matchScore
	
}

type Game struct {
	winner string
	matches map[]

}