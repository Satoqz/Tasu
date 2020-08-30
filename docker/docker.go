package docker

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/bwmarrin/snowflake"
	"github.com/satoqz/tasu/config"
)

func createNode() (node *snowflake.Node) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Snowflake for eval container directory snowflakes
var Snowflake *snowflake.Node = createNode()

// Run a docker command and return its output
func Run(args []string) (out string, err error) {
	cmd := exec.Command(
		"docker",
		args...,
	)
	result, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	out = fmt.Sprintf("%s", result)
	return
}

// StartContainers starts containers for all languages in config
func StartContainers() {
	for _, language := range config.Config.Languages {
		startContainer(language)
	}
}

func startContainer(language string) {
	log.Printf("Starting container: %s\n", language)
	out, err := exec.Command(
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
		fmt.Printf("%s container: %s\n", language, out)
	}
}

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
		fmt.Printf("%s container: %s\n", language, out)
	}
}
