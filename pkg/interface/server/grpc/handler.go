package grpc

import (
	"goarch/pkg/interface/container"
	pbUser "goarch/pkg/shared/grpc/pb/user"
)

type Handler struct {
	crudUserHandler pbUser.UserCrudHandlerServer
}

func SetupHandlers(container *container.Container) *Handler {
	if container == nil {
		panic("container si nil")
	}

	crudUserHandler := SetupCrudUser(container.UserSvc)

	return &Handler{
		crudUserHandler: crudUserHandler,
	}
}
