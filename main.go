package main

import (
	"os"

	"github.com/satoqz/tasu/config"
	"github.com/satoqz/tasu/docker"
	"github.com/satoqz/tasu/router"
)

func main() {
	config.LoadConfig()
	for _, arg := range os.Args {
		if arg == "--buildContainers" || arg == "-bc" {
			docker.BuildContainers()
			break
		}
	}
	docker.StartContainers()
	router.Setup()
}
