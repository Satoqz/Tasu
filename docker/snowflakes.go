package docker

import (
	"log"

	"github.com/bwmarrin/snowflake"
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
