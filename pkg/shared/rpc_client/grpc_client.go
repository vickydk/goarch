package rpc_client

import (
	"context"
	"encoding/json"
	"time"

	ctxSess "goarch/pkg/shared/utils/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type RpcConnection struct {
	options Options
	Conn    *grpc.ClientConn
}

func (rpc *RpcConnection) CreateContext(parent context.Context, session *ctxSess.Context) (ctx context.Context) {
	ctx, _ = context.WithTimeout(parent, rpc.options.Timeout*time.Second)
	ctx = context.WithValue(ctx, ctxSess.AppSession, session)

	header, _ := json.Marshal(session.Header)

	ctx = metadata.AppendToOutgoingContext(ctx,
		XRequestID, session.XRequestID,
		SourceService, session.AppName,
		Header, string(header),
	)

	return
}

func NewGRpcConnection(options Options) *RpcConnection {
	var conn *grpc.ClientConn
	var err error

	conn, err = grpc.Dial(options.Address, grpc.WithInsecure(), withClientUnaryInterceptor())

	if err != nil {
		panic(err)
	}

	return &RpcConnection{
		Conn:    conn,
		options: options,
	}
}

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	session := ctx.Value(ctxSess.AppSession).(*ctxSess.Context)
	ctxWithMetadata := metadata.AppendToOutgoingContext(ctx, XRequestID, session.XRequestID)
	md, _ := metadata.FromOutgoingContext(ctxWithMetadata)
	processTime := session.Lv2("[request]", method, md, req)
	err := invoker(ctxWithMetadata, method, req, reply, cc, opts...)
	if err != nil {
		session.Lv3(processTime, "[response][error]", method, err)
		return err
	}
	session.Lv3(processTime, "[response]", method, reply)
	return err
}

func withClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}
