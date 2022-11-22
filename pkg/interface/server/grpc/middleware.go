package grpc

import (
	"context"
	"goarch/pkg/shared/logger"

	"goarch/pkg/interface/container"
	ctxSess "goarch/pkg/shared/utils/context"
	"google.golang.org/grpc"
)

func middleware(container *container.Container) grpc.UnaryServerInterceptor {
	cfg := container.Config.Apps

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		sess := ctxSess.New(logger.GetLogger()).
			SetAppName(cfg.Name).
			SetPort(cfg.GRPCPort).
			SetRequest(req).
			SetURL(info.FullMethod).
			SetMethod("GRPC")

		c := context.WithValue(ctx, ctxSess.AppSession, sess)
		h, err := handler(c, req)
		return h, err
	}
}
