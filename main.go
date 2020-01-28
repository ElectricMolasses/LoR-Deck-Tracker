package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type deck struct {
	DeckCode    string
	CardsInDeck map[string]int
}

type card struct {
	CardCode           string
	Description        string
	LevelUpDescription string
	Attack             int
	Health             int
	Cost               int
	Keywords           []string
	SpellSpeed         string
	Type               string
}

func main() {
	deckList := parseDeck("{\"DeckCode\":\"CEAAEBYBAEDRMGREFYZDKCABAABQMCYSCQNB2JYCAQAQABYMFIWAMAIBBEKCAIRHFE\",\"CardsInDeck\":{\"00IO004\":1,\"00IO015\":1,\"00IO008\":1,\"00IO006\":1,\"00IO010\":1,\"00IO014\":1,\"00IO012T2\":1,\"00IO005\":1,\"00IO016\":1}}")

	fmt.Println(deckList.CardsInDeck)

	for i := range deckList.CardsInDeck {
		fmt.Println(i, deckList.CardsInDeck[i])
	}

	populateCardLibrary()

	// // Code tested and working with actual client
	// data, err := GetDeck()

	// fmt.Println(data)
	// fmt.Println(err)
	// if err == nil {
	// 	deck := parseDeck(data)
	// 	fmt.Println(deck)
	// 	for i := range deck.CardsInDeck {
	// 		fmt.Println(i, deck.CardsInDeck[i])
	// 	}
	// }
}

func parseDeck(data string) deck {
	// var cards [40]Card // 40 is deck size.
	var result deck

	json.Unmarshal([]byte(data), &result)

	fmt.Println(result)
	return result
}

func populateCardLibrary() ([]card, error) {
	var cards []card

	content, error := ioutil.ReadFile("en_us/data/set1-en_us.json")

	if error != nil {
		fmt.Println("Library population failed with:", error)
		return nil, error
	}

	error = json.Unmarshal(content, &cards)
	if error != nil {
		fmt.Println(error)
	}

	return cards, error
}

// func findCardByID(string) (card, error) {

// }
