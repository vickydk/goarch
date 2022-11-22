package grpc

import (
	"context"
	pbUser "goarch/pkg/shared/grpc/pb/user"
	ctxSess "goarch/pkg/shared/utils/context"
	userSvc "goarch/pkg/usecase/user"
)

type crudUserGRPC struct {
	pbUser.UnimplementedUserCrudHandlerServer
	svc userSvc.Service
}

func SetupCrudUser(svc userSvc.Service) *crudUserGRPC {
	h := &crudUserGRPC{svc: svc}
	if h.svc == nil {
		panic("please provide cash out service")
	}
	return h
}

func (c crudUserGRPC) RegisterUser(ctx context.Context, req *pbUser.RegisterRequest) (res *pbUser.RegisterResponse, err error) {
	sess := ctx.Value(ctxSess.AppSession).(*ctxSess.Context).SetXRequestID(req.RequestId)
	sess.Lv1("incoming request")

	in := &userSvc.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	out, serviceErr := c.svc.RegisterUserGrpc(sess, in)
	status, message := "00", "Success"
	if serviceErr != nil {
		status, message = "01", serviceErr.Error()
	}

	data := &pbUser.User{}
	if out != nil {
		data.Email = out.Email
		data.Name = out.Name
	}

	res = &pbUser.RegisterResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	sess.Lv4(res)

	return
}
