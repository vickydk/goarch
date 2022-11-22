package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"goarch/pkg/usecase/user"
)

type userHandler struct {
	validate *validator.Validate
	service  user.Service
}

func SetupUserHandler(validate *validator.Validate, service user.Service) *userHandler {
	handler := &userHandler{
		validate: validate,
		service:  service,
	}
	if handler.service == nil {
		panic("service is nil")
	}
	return handler
}

func (s *userHandler) registerUser(c echo.Context) error {
	context := Parse(c)
	ctxSess := context.CtxSess
	request := &user.RegisterRequest{}
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
	resp, err := s.service.RegisterUser(ctxSess, request)
	if err != nil {
		ctxSess.Lv4()
		httpCode, errMsg := errHandler(err)
		return httpError(c, httpCode, errors.New(errMsg))
	}

	ctxSess.Lv4(resp)
	return c.JSON(http.StatusOK, resp)
}

func (s *userHandler) resetPassword(c echo.Context) error {
	context := Parse(c)
	ctxSess := context.CtxSess
	request := &user.ResetPasswordReq{}
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
	err := s.service.ResetPassword(ctxSess, request)
	if err != nil {
		ctxSess.Lv4()
		httpCode, errMsg := errHandler(err)
		return httpError(c, httpCode, errors.New(errMsg))
	}

	ctxSess.Lv4()
	return c.JSON(http.StatusOK, struct{}{})
}

func (s *userHandler) updateName(c echo.Context) error {
	context := Parse(c)
	ctxSess := context.CtxSess
	request := &user.UpdateNameReq{}
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
	err := s.service.UpdateName(ctxSess, request)
	if err != nil {
		ctxSess.Lv4()
		httpCode, errMsg := errHandler(err)
		return httpError(c, httpCode, errors.New(errMsg))
	}

	ctxSess.Lv4()
	return c.JSON(http.StatusOK, struct{}{})
}

func (s *userHandler) updatePassword(c echo.Context) error {
	context := Parse(c)
	ctxSess := context.CtxSess
	request := &user.UpdatePasswordReq{}
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
	err := s.service.UpdatePassword(ctxSess, request)
	if err != nil {
		ctxSess.Lv4()
		httpCode, errMsg := errHandler(err)
		return httpError(c, httpCode, errors.New(errMsg))
	}

	ctxSess.Lv4()
	return c.JSON(http.StatusOK, struct{}{})
}
