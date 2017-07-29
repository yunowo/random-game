package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	AppEndpoint  string
	ApiEndpoint  string
	CookieDomain string
}

func loadConfig() Config {
	var config Config
	var path string
	if os.Getenv("RND_DEBUG") == "1" {
		path = "config_debug.json"
	} else {
		path = "config.json"
	}
	configFile, err := os.Open(os.Getenv("RND_CONFIG") + path)
	defer configFile.Close()
	if err != nil {
		log.Fatalf(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

var ClientConfigFile = os.Getenv("RND_CONFIG") + "client.json"
