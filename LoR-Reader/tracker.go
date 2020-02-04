package main

import (
	"encoding/json"
	"fmt"
)

type board struct {
	PlayerName   string
	OpponentName string
	GameState    string
	Screen       screen
	Rectangles   []fieldCard
}

type screen struct {
	ScreenWidth  int
	ScreenHeight int
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

func getBoard() (board, error) {
	data, err := queryClient("positional-rectangles")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	var fieldCards board
	err = json.Unmarshal([]byte(data), &fieldCards)
	if err != nil {
		panic(err)
	}
	fmt.Println(fieldCards)

	return fieldCards, err
}

func initDeck(deck deck) {
	playerDeck = deck
}

func processFieldChange(board board) {

}
