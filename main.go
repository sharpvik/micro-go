package main

import (
	"github.com/sharpvik/log-go/v2"

	"github.com/sharpvik/micro-go/configs"
	"github.com/sharpvik/micro-go/server"
)

var config *configs.Config

func init() {
	log.SetLevel(log.LevelDebug)
	log.Debug("initialising ...")
	config = configs.MustInit()
	log.Debug("init successfull")
}

func main() {
	serv := server.MustInit(config)
	done := make(chan bool, 1)
	go serv.ServeWithGrace(done)
	<-done
	log.Debug("server stopped")
}
