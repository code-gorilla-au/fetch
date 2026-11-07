package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/code-gorilla-au/fetch"
)

func main() {

	url := "https://icanhazdadjoke.com/"
	opts := fetch.Options{
		DefaultHeaders: map[string]string{
			"Accept": "application/json",
		},
	}

	client := fetch.New(&opts)

	resp, err := client.Get(url, nil)
	if err != nil {
		if apiErr, ok := errors.AsType[*fetch.APIError](err); ok {
			fmt.Println("API Response error", apiErr)
			os.Exit(1)
		}

		fmt.Println("Client Error", err)
		os.Exit(1)
	}

	joke := dadJoke{}
	if err = json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		fmt.Println("json decode error ", err)
		os.Exit(1)
	}
	prettyPrintJson(joke)
}

type dadJoke struct {
	ID     string `json:"id,omitempty"`
	Joke   string `json:"joke,omitempty"`
	Status int    `json:"status,omitempty"`
}

func prettyPrintJson(v any) {
	data, _ := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(data))
}
