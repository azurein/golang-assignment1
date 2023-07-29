package main

import (
	"assignment03/monitoring/pkg/database"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

const (
	SAFE    string = "aman"
	WARNING        = "siaga"
	DANGER         = "bahaya"
)

func main() {
	db, err := database.ConnectPostgres(database.DBOption{
		Host:     "localhost",
		Port:     "5432",
		User:     "...",
		Password: "...",
		DBName:   "assignment03",
		SSLMode:  "disable",
	})
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db not connected!")
	}
	query := `
		SELECT id FROM "monitoring"
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	if !rows.Next() {
		query = `
			INSERT INTO "monitoring" ("weather", "values", "units", "status", "updated_at") values
			('water', 0, 'meter', '', now()),
			('wind', 0, 'meter/sec', '', now())
		`
		stmt, err = db.Prepare(query)
		if err != nil {
			panic(err)
		}
		_, err = stmt.Query()
		if err != nil {
			panic(err)
		}
	}
	stmt.Close()

	for i := 0; i < 1; i++ {
		water, wind := checkWeather()
		waterStatus := getWaterStatus(water)
		windStatus := getWindStatus(wind)
		log.Println("water", water, "meter, status:", waterStatus)
		log.Println("wind", wind, "meter/sec, status:", windStatus)

		query = `UPDATE "monitoring" SET "values" = $1, "status" = $2 WHERE "weather" = 'water'`
		stmt, err = db.Prepare(query)
		if err != nil {
			panic(err)
		}
		_, err = stmt.Query(water, waterStatus)
		if err != nil {
			panic(err)
		}

		query = `UPDATE "monitoring" SET "values" = $1, "status" = $2 WHERE "weather" = 'wind'`
		stmt, err = db.Prepare(query)
		if err != nil {
			panic(err)
		}
		_, err = stmt.Query(wind, windStatus)
		if err != nil {
			panic(err)
		}
		stmt.Close()

		time.Sleep(3 * time.Second)
		i--
	}

}

func checkWeather() (water int, wind int) {
	resp, err := http.Get("http://127.0.0.1:8080/api/weather/get")
	if err != nil {
		return water, wind
	}
	if resp.StatusCode != http.StatusOK {
		return water, wind
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return water, wind
	}

	var weather = Weather{}
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return water, wind
	}

	return weather.Water, weather.Wind
}

func getWaterStatus(water int) (status string) {
	switch {
	case water <= 5:
		return SAFE
	case water >= 6 && water <= 8:
		return WARNING
	case water > 8:
		return DANGER
	default:
		return
	}
}

func getWindStatus(wind int) (status string) {
	switch {
	case wind <= 6:
		return SAFE
	case wind >= 7 && wind <= 15:
		return WARNING
	case wind > 15:
		return DANGER
	default:
		return
	}
}
