package main

import (
	"os"

	"github.com/satoqz/tasu/config"
	"github.com/satoqz/tasu/containers"
	"github.com/satoqz/tasu/router"
)

func main() {

	config.LoadConfig()

	// check if containers should be built first
	for _, arg := range os.Args {
		if arg == "--buildContainers" || arg == "-bc" {
			containers.BuildAll()
			break
		}
	}

	// make a map of containers
	containers.MakeMap()
	// start all containers wanted in config
	containers.RestartAll()
	// start container cleanup interval in different goroutine
	go containers.StartCleanupInterval()
	// finally, start webserver
	router.Setup()
}
