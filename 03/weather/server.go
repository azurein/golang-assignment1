package main

import (
	"assignment03/weather/handler"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	log.Println("server running at port", PORT)
	http.HandleFunc("/api/weather/get", handler.GetWeather)
	http.ListenAndServe(PORT, nil)
}
