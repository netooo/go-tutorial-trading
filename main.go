package main

import (
	"fmt"
	"github.com/netooo/trade/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
