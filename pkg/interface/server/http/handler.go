package http

import "goarch/pkg/interface/container"

type Handler struct {
	userHandler *userHandler
	authHandler *authHandler
}

func SetupHandlers(container *container.Container) *Handler {
	userHandler := SetupUserHandler(container.Validate, container.UserSvc)
	authHandler := SetupAuthHandler(container.Validate, container.AuthSvc)
	return &Handler{
		userHandler: userHandler,
		authHandler: authHandler,
	}
}
