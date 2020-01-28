package main

import (
	"io/ioutil"
	"net/http"
)

var Localhost string = "http://localhost:"
var Port string = "21337"

// GET http://localhost:{port}/static-decklist
func QueryClient(request string) (string, error) {
	resp, err := http.Get(Localhost + Port + "/" + request)

	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)

	return string(data), err
}

func GetDeck() (string, error) {
	data, err := QueryClient("static-decklist")

	if err != nil {
		return "", err
	}

	return data, err
}
