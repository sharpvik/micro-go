package main

import (
	"log"

	"github.com/sharpvik/micro-go/configs"
	"github.com/sharpvik/micro-go/database"
	"github.com/sharpvik/micro-go/database/users"
	"github.com/sharpvik/micro-go/service"
)

func main() {
	log.Print("initializing ...")
	config := configs.MustInit()
	db := database.MustInit(config.Database)
	log.Print("init successful")

	names := users.NewRepo(db)
	server := service.New(names).Server()

	if err := server.Start(config.Server.Address); err != nil {
		log.Println("server shutdown with error:", err)
	} else {
		log.Print("server stopped")
	}
}
