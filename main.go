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
	deckList := parseDeck("{\"DeckCode\":\"CEAAEBYBAEDRMGREFYZDKCABAABQMCYSCQNB2JYCAQAQABYMFIWAMAIBBEKCAIRHFE\",\"CardsInDeck\":{\"01IO004\":1,\"01IO015\":1,\"01IO008\":1,\"01IO006\":1,\"01IO010\":1,\"01IO014\":1,\"01IO012T2\":1,\"01IO005\":1,\"01IO016\":1}}")

	library, error := populateCardLibrary()

	if error != nil {
		fmt.Println("Library population failed with:", error)
	}

	for i := range deckList.CardsInDeck {
		fmt.Println(library[i])
	}

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

func populateCardLibrary() (map[string]card, error) {
	var cardsArray []card
	var cardsMap = make(map[string]card)

	content, error := ioutil.ReadFile("en_us/data/set1-en_us.json")

	if error != nil {
		fmt.Println("Library population failed with:", error)
		return nil, error
	}

	error = json.Unmarshal(content, &cardsArray)
	if error != nil {
		fmt.Println(error)
	}

	for i := range cardsArray {
		cardID := cardsArray[i].CardCode
		cardsMap[cardID] = cardsArray[i]
	}

	return cardsMap, error
}

// func findCardByID(string) (card, error) {

// }
