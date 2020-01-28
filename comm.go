package main

import (
	"io/ioutil"
	"net/http"
)

var localhost string = "http://localhost:"
var port string = "21337"

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
