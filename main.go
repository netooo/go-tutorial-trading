package main

import (
	"fmt"
	"github.com/netooo/trade/app/models"
	"github.com/netooo/trade/config"
	"github.com/netooo/trade/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
}
