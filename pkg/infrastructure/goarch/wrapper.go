package goarch

import (
	pbUser "goarch/pkg/shared/grpc/pb/user"
	ctxSess "goarch/pkg/shared/utils/context"
)

type CrudUserWrapper interface {
	RegisterUser(session *ctxSess.Context, email, password, name string) (resp *pbUser.User, err error)
}
