package user

import (
	"context"
	"log"
)

type service struct {
	logger log.Logger
}

// Service interface describes actions on Users
// Users(Store) - defines the interface we expect our database implementation to follow
type Users interface {
	Authenticate(id string) error
	CreateUser(pwdHash string, name string, age int, addInfo string) (User, error)
	UpdateUser(usr User) error
	GetUser(id int32) (User, error)
	DeleteUser(id int32) error
}

// NewService returns a Service with all of the expected dependencies
func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s Service) CreateUser(ctx context.Context, pwdHash string, name string, age int, addInfo string) (User, error) {
	usr, err := s.User.CreateUser(pwdHash, name, age, addInfo)
	if err != nil {
		return User{}, err
	}
	return usr, nil
}
