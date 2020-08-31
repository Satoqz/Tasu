package docker

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/satoqz/tasu/config"
)

// StartContainers starts containers for all languages in config
func StartContainers() {
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
	list := strings.Split(res, "\n")
	list = list[:len(list)-1]

	for _, language := range config.Languages {
		alreadyStarted := false
		for _, containerName := range list {
			if fmt.Sprintf("tasu_%s", language) == containerName {
				alreadyStarted = true
			}
		}
		if !alreadyStarted {
			startContainer(language)
		} else {
			Containers[fmt.Sprintf("tasu_%s", language)] = ContainerStruct{
				Language: language,
				Alive:    true,
			}
		}
	}
}

func startContainer(language string) {
	log.Printf("Starting container: %s\n", language)
	_, err := exec.Command(
		"docker",
		"run",
		"--runtime=runsc",
		"--rm",
		fmt.Sprintf("--name=tasu_%s", language),
		"-u1000:1000",
		"-w/tmp/",
		"-dt",
		"--net=none",
		"--cpus=0.25",
		"-m=128m",
		"--memory-swap=128m",
		fmt.Sprintf("tasu_%s:latest", language),
		"/bin/sh",
	).Output()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Started container: %s\n", language)
		container := ContainerStruct{
			Language: language,
			Alive:    true,
		}
		Containers[fmt.Sprintf("tasu_%s", language)] = container
	}
}
