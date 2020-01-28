package main

import "fmt"

type board struct {
	PlayerName   string
	OpponentName string
	GameState    string
	// Screen Object
	Rectangles []fieldCard
}

type fieldCard struct {
	LocalPlayer bool
	TopLeftX    int
	TopLeftY    int
	Width       int
	Height      int
	Card        card
}

var playerDeck deck

// Store the discard as an array to track order.
// The same will go for the field, but they'll contain
// maps that contain cards, to track board position, etc.
var playerDiscard []card
var playerField []int

var opponentField []int

func getField() {
	data, err := queryClient("positional-rectangles")
	fmt.Println(data, err)
}

func initDeck(deck deck) {
	playerDeck = deck
}
