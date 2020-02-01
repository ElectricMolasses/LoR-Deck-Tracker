package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
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
	// Test against static data before testing against client.
	deckList := parseDeck("{\"DeckCode\":\"CEAAEBYBAEDRMGREFYZDKCABAABQMCYSCQNB2JYCAQAQABYMFIWAMAIBBEKCAIRHFE\",\"CardsInDeck\":{\"01IO004\":1,\"01IO015\":1,\"01IO008\":1,\"01IO006\":1,\"01IO010\":1,\"01IO014\":1,\"01IO012T2\":1,\"01IO005\":1,\"01IO016\":1}}")

	library, error := populateCardLibrary()

	if error != nil {
		fmt.Println("Library population failed with:", error)
	}

	for i := range deckList.CardsInDeck {
		fmt.Println(library[i])
	}

	initDeck(deckList)

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

	// Loop to continue checking board state
	// Entire app refreshes on ticker time.
	timedLoop := time.NewTicker(time.Second)
	for {
		<-timedLoop.C

		getBoard()
	}
}

func parseDeck(data string) deck {
	// var cards [40]Card // 40 is deck size.
	var result deck

	json.Unmarshal([]byte(data), &result)

	return result
}

func populateCardLibrary() (map[string]card, error) {
	var cardsArray []card
	var cardsMap = make(map[string]card)

	content, err := ioutil.ReadFile("en_us/data/set1-en_us.json")

	if err != nil {
		fmt.Println("Library population failed with:", err)
		return nil, err
	}

	err = json.Unmarshal(content, &cardsArray)
	if err != nil {
		fmt.Println(err)
	}

	for i := range cardsArray {
		cardID := cardsArray[i].CardCode
		cardsMap[cardID] = cardsArray[i]
	}

	return cardsMap, err
}
