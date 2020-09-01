package containers

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/satoqz/tasu/config"
)

func start(language string) {
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
		fmt.Sprintf("-m=%dm", config.RAM),
		fmt.Sprintf("--memory-swap=%dm", config.SWAP),
		fmt.Sprintf("tasu_%s:latest", language),
		"/bin/sh",
	).Output()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Started container: %s\n", language)
		container := Container{
			Language: language,
			Alive:    true,
		}
		Map[fmt.Sprintf("tasu_%s", language)] = container
	}
}
