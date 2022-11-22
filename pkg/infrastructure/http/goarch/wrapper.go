package goarch

import ctxSess "goarch/pkg/shared/utils/context"

type Wrapper interface {
	GetUserDetail(ctx *ctxSess.Context, userId int64) (out GetUserDetailResponse, err error)
}
