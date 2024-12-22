package main

import (
	"net/http"

	"gopkg.in/dnaeon/go-vcr.v3/cassette"
)

// probe calls the API server to check what we can do
func probe() error {
	// define http calls here, e.g.:
	_, err := http.Get("https://api.ipify.org/?format=json")
	return err
}

// mask any secrets the API might return, e.g. in the response body
func maskSecrets(i *cassette.Interaction) error {
	return nil
}
