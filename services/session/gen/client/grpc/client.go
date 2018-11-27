package session_clientgrpc

import (
	context "context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	endpoints "github.com/sjdweb/go-kit-protoc-template/services/session/gen/endpoints"
	pb "github.com/sjdweb/go-kit-protoc-template/services/session/gen/pb"
)

func New(conn *grpc.ClientConn, logger log.Logger) pb.SessionServiceServer {

	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = grpctransport.NewClient(
			conn,
			"session.SessionService",
			"Login",
			EncodeLoginRequest,
			DecodeLoginResponse,
			pb.LoginResponse{},
		).Endpoint()
	}

	return &endpoints.Endpoints{

		LoginEndpoint: loginEndpoint,
	}
}

func EncodeLoginRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LoginRequest)
	return req, nil
}

func DecodeLoginResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.LoginResponse)
	return response, nil
}
