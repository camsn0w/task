package main

import "encoding/json"

type Joke struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

// TODO: add input validation for cases where it doesn't match the expected json.
// bytesToJoke() takes bytes from a request and returns
// a Joke and error.
func bytesToJoke(rawName []byte) (Joke, error) {
	joke := Joke{}
	err := json.Unmarshal(rawName, &joke)
	return joke, err
}
