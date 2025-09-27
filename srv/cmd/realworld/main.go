package main

import (
	"log"

	"github.com/tim2006/realworld/private/api"
	"github.com/tim2006/realworld/private/db"
)

func main() {

	database, err := db.InitDB("realworld.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	api.Server{
		Address: ":8080",
	}.Run()
}
