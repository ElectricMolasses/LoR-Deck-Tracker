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
var player []fieldCard

var opponent []fieldCard

func getBoard() (board, error) {
	data, err := queryClient("positional-rectangles")
	if err != nil {
		panic(err)
	}

	var currentBoard board
	err = json.Unmarshal([]byte(data), &currentBoard)
	if err != nil {
		panic(err)
	}

	playerCards := getPlayerCards(currentBoard.Rectangles)
	fmt.Println(playerCards)

	return currentBoard, err
}

func initDeck(deck deck) {
	playerDeck = deck
}

// Will return an array of all cards that were not
// on the board or handalready, and then an array of all
// cards that were on the board or hand, but have since been
// removed.
func boardChanges(oldBoard, newBoard board) ([]card, []card) {
	var newCards []card
	var removedCards []card

	return newCards, removedCards
}

func getPlayerCards(fieldCards []fieldCard) []fieldCard {
	var playerCards []fieldCard
	for i := range fieldCards {
		if fieldCards[i].LocalPlayer {
			playerCards = append(playerCards, fieldCards[i])
		}
	}

	return playerCards
}

func handleRemovedCard(removedCard card) {
	// Mostly just removes from field for extra functionality later.
	// Immediate functionality required will be parsing card description
	// in case the card adds or returns something to the players library.
}

func handleAddedCard(newCard fieldCard) {
	if newCard.LocalPlayer {
		playerDeck.CardsInDeck[newCard.Card.CardCode]--
		fmt.Println("You played:", newCard.Card.Name)
	}
}
