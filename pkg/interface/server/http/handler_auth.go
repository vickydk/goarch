package http

import (
	"errors"
	"fmt"
	"net/http"

	authSvc "goarch/pkg/usecase/auth"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	validate *validator.Validate
	service  authSvc.Service
}

func SetupAuthHandler(validate *validator.Validate, service authSvc.Service) *authHandler {
	handler := &authHandler{
		validate: validate,
		service:  service,
	}
	if handler.service == nil {
		panic("service is nil")
	}
	return handler
}

func (s *authHandler) login(c echo.Context) error {
	context := Parse(c)
	ctxSess := context.CtxSess
	request := &authSvc.LoginRequest{}
	if err := c.Bind(request); err != nil {
		ctxSess.ErrorMessage = err.Error()
		ctxSess.Lv4()
		return httpError(c, http.StatusBadRequest, fmt.Errorf("bind request: %w", err))
	}
	if err := s.validate.Struct(request); err != nil {
		ctxSess.ErrorMessage = err.Error()
		ctxSess.Lv4()
		return httpError(c, http.StatusBadRequest, fmt.Errorf("validate: %w", err))
	}

	ctxSess.Request = request
	resp, err := s.service.Login(ctxSess, request)
	if err != nil {
		ctxSess.Lv4()
		httpCode, errMsg := errHandler(err)
		return httpError(c, httpCode, errors.New(errMsg))
	}

	ctxSess.Lv4()
	return c.JSON(http.StatusOK, resp)
}
