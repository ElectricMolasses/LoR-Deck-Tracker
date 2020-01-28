package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var localhost string = "http://localhost:"
var port string = "21337"

type deck struct {
	DeckCode    string
	CardsInDeck map[string]int
}

type card struct {
	CardId   string
	Quantity int
}

func main() {
	// data, _ := getDeck()
	// fmt.Println(data)
	// Contains DeckCode and CardsInDeck
	deckList := parseDeck("{\"DeckCode\":\"CEAAEBYBAEDRMGREFYZDKCABAABQMCYSCQNB2JYCAQAQABYMFIWAMAIBBEKCAIRHFE\",\"CardsInDeck\":{\"00IO004\":1,\"00IO015\":1,\"00IO008\":1,\"00IO006\":1,\"00IO010\":1,\"00IO014\":1,\"00IO012T2\":1,\"00IO005\":1,\"00IO016\":1}}")

	fmt.Println(deckList.CardsInDeck)
	// parseDeck(data)

	for i := range deckList.CardsInDeck {
		fmt.Println(i)
	}
}

// GET http://localhost:{port}/static-decklist
func queryClient(request string) (string, error) {
	resp, err := http.Get(localhost + port + "/" + request)

	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)

	return string(data), err
}

func getDeck() (string, error) {
	data, err := queryClient("static-decklist")

	if err != nil {
		return "", err
	}

	return data, err
}

func parseDeck(data string) deck {
	// var cards [40]Card // 40 is deck size.
	var result deck

	json.Unmarshal([]byte(data), &result)

	fmt.Println(result)
	return result
}

/*
{
	"DeckCode":"CEAAEBYBAEDRMGREFYZDKCABAABQMCYSCQNB2JYCAQAQABYMFIWAMAIBBEKCAIRHFE",
	"CardsInDeck":
		{"00IO004":1,"00IO015":1,"00IO008":1,"00IO006":1,"00IO010":1,"00IO014":1,"00IO012T2":1,"00IO005":1,"00IO016":1}
}
*/
