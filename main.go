package main

import (
	"log"

	"github.com/sharpvik/micro-go/configs"
	"github.com/sharpvik/micro-go/database"
	"github.com/sharpvik/micro-go/server"
)

func main() {
	log.Print("initializing ...")
	config := configs.MustInit()
	db := database.MustInit(config.Database)
	log.Print("init successful")
	if err := server.Runtime(db).Echo().Start(config.Server.Address); err != nil {
		log.Println("server shutdown with error:", err)
	} else {
		log.Print("server stopped")
	}
}
