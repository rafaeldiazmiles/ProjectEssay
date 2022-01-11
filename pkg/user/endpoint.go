package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/junereycasuga/gokit-grpc-demo/service"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	CreateUser endpoint.Endpoint
}

// MathReq struct holds the endpoint request definition
type UserRequest struct {
	pwdHash string
	name    string
	age     int
	addInfo string
	// parents []Parent   --> Para implementar cuando haya parents
}

// MathResp struct holds the endpoint response definition
type UserResp struct {
	id string
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Add: makeCreateUserEndpoint(s),
	}
}

func makeAddEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUser)
		result, _ := s.CreateUser(ctx, pwdHash, name, age, addInfo)
		return CreateUser{Result: result}, nil
	}
}
