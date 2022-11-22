package auth

import (
	ctxSess "goarch/pkg/shared/utils/context"
)

type Service interface {
	Login(ctxSess *ctxSess.Context, req *LoginRequest) (res *LoginResponse, err error)
}
