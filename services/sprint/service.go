package sprintsvc

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/sjdweb/go-kit-protoc-template/services/sprint/gen/pb/sprint"
)

type Service struct{}

func New() pb.SprintServiceServer {
	return &Service{}
}

func (svc *Service) AddSprint(ctx context.Context, in *pb.AddSprintRequest) (*pb.AddSprintResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) CloseSprint(ctx context.Context, in *pb.CloseSprintRequest) (*pb.CloseSprintResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) GetSprint(ctx context.Context, in *pb.GetSprintRequest) (*pb.GetSprintResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
