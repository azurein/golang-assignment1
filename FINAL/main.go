package main

import (
	"finalproject/domain/comment"
	"finalproject/domain/photo"
	"finalproject/domain/socialmedia"
	"finalproject/domain/user"
	"finalproject/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectPostgres(database.DBOption{
		Host:     "localhost",
		Port:     "5432",
		User:     "...",
		Password: "...",
		DBName:   "finalproject",
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

	user.RegisterRoute(api, db)
	socialmedia.RegisterRoute(api, db)
	photo.RegisterRoute(api, db)
	comment.RegisterRoute(api, db)

	router.Run(":8001")
}
