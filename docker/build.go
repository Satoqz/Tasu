package docker

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/satoqz/tasu/config"
)

// BuildContainers tries to build containers for all languages in config
func BuildContainers() {
	for _, language := range config.Config.Languages {
		buildContainer(language)
	}
}

func buildContainer(language string) {
	log.Printf("Building container: %s\n", language)
	out, err := exec.Command(
		"docker",
		"build",
		fmt.Sprintf("./languages/%s", language),
		"-t",
		fmt.Sprintf("tasu_%s", language),
	).Output()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s container: %s\n", language, out)
	}
}
