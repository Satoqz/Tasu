package docker

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/satoqz/tasu/config"
)

// ContainerStruct is a struct containing info on a tasu container
type ContainerStruct struct {
	Language string
	Alive    bool
}

// Containers is an array storing all tasu docker containers and their states
var Containers map[string]ContainerStruct = make(map[string]ContainerStruct)

func includes(arr []string, val string) (res bool) {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return
}

// RefreshContainers refreshes the "alive" status of all containers in `Containers`
func RefreshContainers() {
	res, err := Run([]string{
		"ps",
		"--filter",
		"name=tasu_",
		"--format",
		"{{.Names}}",
	})
	if err != nil {
		log.Fatal(err)
	}
	alive := strings.Split(res, "\n")
	alive = alive[:len(alive)-1]

	for k, v := range Containers {
		if !includes(alive, k) {
			Containers[k] = ContainerStruct{
				Alive:    false,
				Language: v.Language,
			}
		}
	}
}

func restartContainers() {
	RefreshContainers()
	for k, v := range Containers {
		if v.Alive {
			restartContainer(k, v)
		} else {
			startContainer(v.Language)
		}
	}
}

func restartContainer(k string, v ContainerStruct) {
	Containers[k] = ContainerStruct{
		Alive:    false,
		Language: v.Language,
	}
	log.Printf("Restarting container %s\n", k)
	_, err := Run([]string{
		"restart",
		k,
	})
	if err != nil {
		log.Printf("Container %s failed to restart\n", k)
		log.Fatal(err)
	}
	Containers[k] = ContainerStruct{
		Language: v.Language,
		Alive:    true,
	}
}

var shouldCleanup bool = true

var wg sync.WaitGroup

// StartCleanupInterval calls restartContainers in the interval specified in the config
func StartCleanupInterval() {
	if shouldCleanup {
		time.Sleep(config.CleanupInterval)
		log.Println("Beginning container cleanup")
		// schedules the KillContainers method to run after finished cleanup
		wg.Add(1)
		restartContainers()
		log.Println("Finished container cleanup")
		wg.Done()
		StartCleanupInterval()
	}
}

// KillContainers kills all remaining alive containers
func KillContainers() (killed []string) {
	wg.Wait() // wait for possible cleanup to finish before killing
	shouldCleanup = false
	RefreshContainers()

	for k, v := range Containers {
		if v.Alive {
			log.Printf("Killing container: %s\n", k)
			_, err := Run([]string{
				"kill",
				k,
			})
			if err != nil {
				log.Printf("Failed killing container: %s\n", k)
				log.Fatal(err)
			} else {
				log.Printf("Killed container: %s\n", k)
				killed = append(killed, k)
				Containers[k] = ContainerStruct{
					Alive:    false,
					Language: v.Language,
				}
			}
		}
	}
	return
}
