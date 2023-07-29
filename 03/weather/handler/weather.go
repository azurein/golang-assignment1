package handler

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

type Weather struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var weather Weather
	weather.Water = rand.Intn(49) + 1
	weather.Wind = rand.Intn(49) + 1
	log.Println(r.URL, "| water:", weather.Water, "| wind:", weather.Wind)

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(weather)
	} else {
		w.Write([]byte("method not allowed"))
	}
}
