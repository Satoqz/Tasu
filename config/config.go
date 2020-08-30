package config

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Languages []string
}

// Config stores all available languages in an array
var Config config

// LoadConfig loads "config.json" into Config
func LoadConfig() {
	configFile, err := os.Open("config.json")
	defer configFile.Close()

	if err != nil {
		log.Fatalln("Error opening config.json")
	}

	json.NewDecoder(configFile).Decode(&Config)
	log.Println("Successfully loaded config")
}
