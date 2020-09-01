package containers

import "log"

// RestartAll starts or restarts all tasu containers
func RestartAll() {
	RefreshMap()
	for k, v := range Map {
		if v.Alive {
			restart(k, v)
		} else {
			start(v.Language)
		}
	}
}

func restart(k string, v Container) {
	Map[k] = Container{
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
	Map[k] = Container{
		Language: v.Language,
		Alive:    true,
	}
}
