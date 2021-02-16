package main

import (
	"github.com/netooo/trade/config"
	"github.com/netooo/trade/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
