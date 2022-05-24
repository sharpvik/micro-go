package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/sharpvik/log-go/v2"

	"github.com/sharpvik/micro-go/configs"
	"github.com/sharpvik/micro-go/database"
	"github.com/sharpvik/micro-go/server"
)

var (
	config *configs.Config
	db     *sqlx.DB
)

func init() {
	log.SetLevel(log.LevelDebug)
	log.Debug("initialising ...")
	config = configs.MustInit()
	db = database.MustInit(config.Database)
	log.Debug("init successfull")
}

func main() {
	if err := server.Runtime(db).Echo().Start(config.Server.Address); err != nil {
		log.Error("server shutdown with error:", err)
	} else {
		log.Debug("server stopped")
	}
}
