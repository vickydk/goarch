package main

import (
	"goarch/pkg/interface/container"
	Http "goarch/pkg/interface/server/http"
)

func main() {
	Http.StartHttpService(container.Setup())
}
