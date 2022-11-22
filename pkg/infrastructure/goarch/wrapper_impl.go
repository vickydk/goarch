package goarch

import (
	"context"
	"errors"

	"goarch/pkg/shared/config"
	pbUser "goarch/pkg/shared/grpc/pb/user"
	RpcClient "goarch/pkg/shared/rpc_client"
	ctxSess "goarch/pkg/shared/utils/context"
)

type crudUserWrapper struct {
	config *config.GoarchConfig
	rcpCon *RpcClient.RpcConnection
	client pbUser.UserCrudHandlerClient
}

func SetupCrudUserWrapper(cfg *config.GoarchConfig) *crudUserWrapper {
	if cfg == nil {
		panic("goarchGrpc config is nil")
	}

	w := &crudUserWrapper{config: cfg}
	w.rcpCon = RpcClient.NewGRpcConnection(w.config.RpcOptions)
	w.client = pbUser.NewUserCrudHandlerClient(w.rcpCon.Conn)
	return w
}

func (w *crudUserWrapper) RegisterUser(session *ctxSess.Context, email, password, name string) (resp *pbUser.User, err error) {
	ctx := w.rcpCon.CreateContext(context.Background(), session)
	out, err := w.client.RegisterUser(ctx, &pbUser.RegisterRequest{
		RequestId: session.XRequestID,
		Email:     email,
		Password:  password,
		Name:      name,
	})
	if err != nil {
		return
	}

	if out.Status != "00" {
		err = errors.New(out.Message)
		return
	}

	resp = out.Data

	return
}
