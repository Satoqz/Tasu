package containers

import (
	"fmt"
	"log"
	"strings"

	"github.com/satoqz/tasu/config"
)

// Container is a struct containing info on a tasu container
type Container struct {
	Language string
	Alive    bool
}

// Map is a map storing all tasu docker containers and their states
var Map map[string]Container = make(map[string]Container)

func includes(arr []string, val string) (res bool) {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return
}

// MakeMap creates a map of containers, this is used once on startup
func MakeMap() {
	for _, lang := range config.Languages {
		Map[fmt.Sprintf("tasu_%s", lang)] = Container{
			Language: lang,
			Alive:    false,
		}
	}
}

// RefreshMap refreshes the container map
func RefreshMap() {
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

	for k, v := range Map {
		if !includes(alive, k) {
			Map[k] = Container{
				Alive:    false,
				Language: v.Language,
			}
		} else {
			Map[k] = Container{
				Alive:    true,
				Language: v.Language,
			}
		}
	}
}
