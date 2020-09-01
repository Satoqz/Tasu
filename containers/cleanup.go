package containers

import (
	"log"
	"sync"
	"time"

	"github.com/satoqz/tasu/config"
)

var shouldCleanup bool = true

var wg sync.WaitGroup

// StartCleanupInterval calls restartMap in the interval specified in the config
func StartCleanupInterval() {
	time.Sleep(config.CleanupInterval)
	if shouldCleanup {
		log.Println("Beginning container cleanup")
		// schedules the Kill method to run after finished cleanup
		wg.Add(1)
		RestartAll()
		log.Println("Finished container cleanup")
		wg.Done()
		StartCleanupInterval()
	}
}

// Kill kills all remaining alive containers, returns all killed containers
func Kill() (killed []string) {
	wg.Wait() // wait for possible cleanup to finish before killing
	shouldCleanup = false
	RefreshMap()

	for k, v := range Map {
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
				Map[k] = Container{
					Alive:    false,
					Language: v.Language,
				}
			}
		}
	}
	return
}
