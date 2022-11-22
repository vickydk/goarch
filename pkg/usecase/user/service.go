package user

import ctxSess "goarch/pkg/shared/utils/context"

type Service interface {
	RegisterUser(ctxSess *ctxSess.Context, req *RegisterRequest) (res *User, err error)
	RegisterUserGrpc(ctxSess *ctxSess.Context, req *RegisterRequest) (res *User, err error)
	ResetPassword(ctxSess *ctxSess.Context, req *ResetPasswordReq) (err error)
	UpdateName(ctxSess *ctxSess.Context, req *UpdateNameReq) (err error)
	UpdatePassword(ctxSess *ctxSess.Context, req *UpdatePasswordReq) (err error)
}
