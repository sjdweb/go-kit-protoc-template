package user_clientgrpc

import (
	context "context"

        "github.com/go-kit/kit/log"
        "google.golang.org/grpc"
        grpctransport "github.com/go-kit/kit/transport/grpc"
        "github.com/go-kit/kit/endpoint"

        pb "github.com/sjdweb/go-kit-protoc-template/services/user/gen/pb"
        endpoints "github.com/sjdweb/go-kit-protoc-template/services/user/gen/endpoints"
)



func New(conn *grpc.ClientConn, logger log.Logger) pb.UserServiceServer {
        
		
			var createuserEndpoint endpoint.Endpoint
			{
				createuserEndpoint = grpctransport.NewClient(
					conn,
					"user.UserService",
					"CreateUser",
					EncodeCreateUserRequest,
					DecodeCreateUserResponse,
					pb.CreateUserResponse{},
				).Endpoint()
			}
		
        
		
			var getuserEndpoint endpoint.Endpoint
			{
				getuserEndpoint = grpctransport.NewClient(
					conn,
					"user.UserService",
					"GetUser",
					EncodeGetUserRequest,
					DecodeGetUserResponse,
					pb.GetUserResponse{},
				).Endpoint()
			}
		
        

        return &endpoints.Endpoints {
                
			
				CreateUserEndpoint: createuserEndpoint,
			
                
			
				GetUserEndpoint: getuserEndpoint,
			
                
        }
}


	
		func EncodeCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.CreateUserRequest)
			return req, nil
		}

		func DecodeCreateUserResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.CreateUserResponse)
			return response, nil
		}
	

	
		func EncodeGetUserRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.GetUserRequest)
			return req, nil
		}

		func DecodeGetUserResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.GetUserResponse)
			return response, nil
		}
	

