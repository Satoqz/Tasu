package docker

// ContainerStruct is a struct containing info on a tasu container
type ContainerStruct struct {
	Language   string
	Restarting bool
	Uses       uint
}

// Containers is an array storing all tasu docker containers and their states
var Containers map[string]ContainerStruct = make(map[string]ContainerStruct)
