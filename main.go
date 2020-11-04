package main

import (
	"flag"

	"github.com/satoqz/tasu/config"
	"github.com/satoqz/tasu/containers"
	"github.com/satoqz/tasu/router"
)

var doBuild = []*bool{
	flag.Bool("buildContainers", false, "Build all containers before startup"),
	flag.Bool("bc", false, "Build all containers before startup"),
}

func main() {
	flag.Parse()
	config.LoadConfig()

	// check flags if containers should be built first
	for _, v := range doBuild {
		if *v == true {
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
