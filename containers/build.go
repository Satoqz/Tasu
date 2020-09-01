package containers

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/satoqz/tasu/config"
)

// BuildAll tries to build containers for all languages in config
func BuildAll() {
	for _, language := range config.Languages {
		build(language)
	}
}

func build(language string) {
	containerName := fmt.Sprintf("tasu_%s", language)
	log.Printf("Building container: %s\n", language)
	_, err := exec.Command(
		"docker",
		"build",
		fmt.Sprintf("./languages/%s", language),
		"-t",
		containerName,
	).Output()
	if err != nil {
		log.Printf("Container %s failed to build\n", containerName)
		log.Fatal(err)
	} else {
		log.Printf("Container built: %s\n", containerName)
	}
}
