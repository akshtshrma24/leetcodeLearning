package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type errorData struct {
	Error string `json:"error"`
	Key   string `json:"key"`
}

type isbnData struct {
	Title string `json:"title"` // title

	Cover []int `json:"covers,omitempty"`

	Works []struct {
		Key string `json:"key"`
	} `json:"works"`
	Key string `json:"key"`
}

func findISBN(isbn string) (*isbnData, error) {
	url := fmt.Sprintf("https://sce.sjsu.edu/isbn/%s", isbn)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	js, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ed errorData
	if err := json.Unmarshal(js, &ed); err == nil && ed.Error != "" {
		return nil, fmt.Errorf("Error parsing errorData")
	}

	var id isbnData
	if err := json.Unmarshal(js, &id); err == nil && id.Title != "" {
		return &id, nil
	}

	return nil, fmt.Errorf("No error related to errorData or isbnData, ")
}

func main() {
	id, ed := findISBN("OL24256462M")
	if ed != nil {
		fmt.Println("Could not find")
	}
	fmt.Println(id)
}
