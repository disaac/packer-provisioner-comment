package main

import (
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterProvisioner(new(CommentProvisioner))
	server.Serve()
}
