package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
	LogFile   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:    cfg.Section("api").Key("api_key").String(),
		ApiSecret: cfg.Section("api").Key("api_secret").String(),
		LogFile:   cfg.Section("trade").Key("log_file").String(),
	}
}
