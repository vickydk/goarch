package main

import (
	"goarch/pkg/interface/container"
	"goarch/pkg/interface/server"
)

func main() {
	server.StartService(container.Setup())
}
