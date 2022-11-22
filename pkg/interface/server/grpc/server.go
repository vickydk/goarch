package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"goarch/pkg/interface/container"
	pbUser "goarch/pkg/shared/grpc/pb/user"

	"github.com/spf13/cast"
	"google.golang.org/grpc"
)

func StartGRPCService(ctx context.Context, container *container.Container) error {
	cfg := container.Config

	port := cast.ToString(cfg.Apps.GRPCPort)
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer(grpc.UnaryInterceptor(middleware(container)))

	// setup handler
	handler := SetupHandlers(container)

	// register consul health check
	pbUser.RegisterUserCrudHandlerServer(server, handler.crudUserHandler)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server... " + port)
	return server.Serve(listen)
}
