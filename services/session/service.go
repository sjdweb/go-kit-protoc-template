package sessionsvc

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/sjdweb/go-kit-protoc-template/services/session/gen/pb"
)

type Service struct{}

func New() pb.SessionServiceServer {
	return &Service{}
}

func (svc *Service) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
