package session_endpoints

import (
	context "context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/sjdweb/go-kit-protoc-template/services/session/gen/pb/session"
	oldcontext "golang.org/x/net/context"
)

var _ = endpoint.Chain
var _ = fmt.Errorf
var _ = context.Background

type StreamEndpoint func(server interface{}, req interface{}) (err error)

type Endpoints struct {
	LoginEndpoint endpoint.Endpoint
}

func (e *Endpoints) Login(ctx oldcontext.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	out, err := e.LoginEndpoint(ctx, in)
	if err != nil {
		return &pb.LoginResponse{Err: err.Error()}, err
	}
	return out.(*pb.LoginResponse), err
}

func MakeLoginEndpoint(svc pb.SessionServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.LoginRequest)
		rep, err := svc.Login(ctx, req)
		if err != nil {
			return &pb.LoginResponse{Err: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeEndpoints(svc pb.SessionServiceServer) Endpoints {
	return Endpoints{

		LoginEndpoint: MakeLoginEndpoint(svc),
	}
}
