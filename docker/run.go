package docker

import (
	"fmt"
	"os/exec"
)

// Run a docker command and return its output
func Run(args []string) (out string, err error) {
	cmd := exec.Command(
		"docker",
		args...,
	)
	result, err := cmd.CombinedOutput()
	out = fmt.Sprintf("%s", result)
	return
}
