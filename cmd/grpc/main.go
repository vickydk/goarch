package main

import (
	"context"

	"goarch/pkg/interface/container"
	GRPC "goarch/pkg/interface/server/grpc"
)

func main() {
	if err := GRPC.StartGRPCService(context.Background(), container.Setup()); err != nil {
		panic(err)
	}
}
