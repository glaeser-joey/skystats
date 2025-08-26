package main

import (
	"io"
	"net/http"
	"os"
)

func Fetch() ([]byte, error) {

	url := os.Getenv("READSB_AIRCRAFT_JSON")

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}
