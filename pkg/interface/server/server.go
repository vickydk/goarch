package server

import (
	"context"

	"goarch/pkg/interface/container"
	GRPC "goarch/pkg/interface/server/grpc"
	Http "goarch/pkg/interface/server/http"
)

func StartService(container *container.Container) {
	// start grpc server
	go func() {
		err := GRPC.StartGRPCService(context.Background(), container)
		if err != nil {
			panic(err)
		}
	}()

	// start http server
	Http.StartHttpService(container)
}
