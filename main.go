package main

import (
	"fmt"
	"github.com/netooo/trade/bitflyer"
	"github.com/netooo/trade/config"
	"github.com/netooo/trade/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
