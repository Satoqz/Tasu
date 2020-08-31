package docker

import (
	"log"
	"time"
)

// RestartContainers restarts all docker containers listes in `Containers`
func RestartContainers() {
	for k, v := range Containers {
		RestartContainer(k, v)
	}
}

// RestartContainer restarts a single container
func RestartContainer(k string, v ContainerStruct) {
	Containers[k] = ContainerStruct{
		Restarting: true,
		Language:   v.Language,
		Uses:       v.Uses,
	}
	log.Printf("Restarting container %s\n", k)
	time.Sleep(time.Minute)
	Containers[k] = ContainerStruct{
		Restarting: false,
		Language:   v.Language,
		Uses:       0,
	}
}
