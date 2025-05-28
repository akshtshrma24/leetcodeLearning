package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const API_KEY = "384530a26df84465ffe7e9419c8ef8db"

type TemperatureData struct {
	TempStruct struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func fetchWeather(city string, countryCode string) interface{} {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s",
		city,
		countryCode,
		API_KEY)

	fmt.Println(url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	var t TemperatureData
	json.NewDecoder(response.Body).Decode(&t)

	return t
}

func main() {
	fmt.Println((fetchWeather("Fremont", "us")))
}
