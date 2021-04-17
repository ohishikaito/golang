package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("failed")
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:    cfg.section("bitflayer").key("api_key").String(),
		ApiSecret: cfg.section("bitflayer").Key("api_secret").String(),
	}
}
