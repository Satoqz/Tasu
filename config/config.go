package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type config struct {
	Languages       []string `json:"languages" binding:"required"`
	CleanupInterval string   `json:"cleanupInterval" binding:"required"`
	RAM             uint     `json:"ram" binding:"required"`
	SWAP            uint     `json:"swap" binding:"required"`
}

// CleanupInterval in config.json
var CleanupInterval time.Duration

// Languages in config.json
var Languages []string

// RAM to use per container
var RAM uint

// SWAP to use per container
var SWAP uint

// LoadConfig loads "config.json" into Config
func LoadConfig() {
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		log.Fatalln("Error opening config.json")
	}

	var jsonConfig config
	json.NewDecoder(configFile).Decode(&jsonConfig)

	interval, err := time.ParseDuration(jsonConfig.CleanupInterval)
	if err != nil {
		log.Fatalln("Error parsing cleanupInterval")
	}

	CleanupInterval = interval
	Languages = jsonConfig.Languages
	RAM = jsonConfig.RAM
	SWAP = jsonConfig.SWAP

	log.Println("Config loaded")
}
