package user

import (
	"context"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/junereycasuga/gokit-grpc-demo/endpoints"
	"github.com/junereycasuga/gokit-grpc-demo/pb"
)

type gRPCServer struct {
	add gt.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.ScalarTypes {
	return &gRPCServer{
		add: gt.NewServer(
			endpoints.Add,
			decodeUserRequest,
			encodeUserResponse,
		),
	}
}

func (s *gRPCServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponse), nil
}

func decodeUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserRequest)
	return endpoints.UserReq{ctx, pwdHash, name, age, addInfo}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.MathResp)
	return &pb.UserResponse{
		pwdHash: pwdHash,
		name:    name,
		age:     age,
		addInfo: addInfo,
	}, nil
}
