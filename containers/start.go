package containers

import (
	"fmt"
	"log"
	"runtime"

	"github.com/satoqz/tasu/config"
)

func start(language string) {

	base := []string{"run"}
	if runtime.GOOS == "windows" {
		base = append(base, "--runtime=runsc")
	}
	cmd := []string{
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
	}
	base = append(base, cmd...)

	log.Printf("Starting container: %s\n", language)
	_, err := Run(base)
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
