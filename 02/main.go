package main

import (
	"assignment02/domain/order"
	"assignment02/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectPostgres(database.DBOption{
		Host:     "localhost",
		Port:     "5432",
		User:     "...",
		Password: "...",
		DBName:   "assignment02",
		SSLMode:  "disable",
	})

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db not connected!")
	}

	log.Println("db connected")

	router := gin.New()

	api := router.Group("api")
	order.RegisterRoute(api, db)

	router.Run(":8888")
}
