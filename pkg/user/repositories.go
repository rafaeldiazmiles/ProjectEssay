//go:generate mockgen -destination=repositories_mocks_test.go -package=rocket https://github.com/rafaeldiazmiles/ProjectEssay Users
package user

import "context"

// User - should contain the definition of our user
type User struct {
	id      int32
	pwdHash string
	name    string
	age     int
	addInfo string
	// parents []Parent   --> Para implementar cuando haya parents
}

// Users(Store) - defines the interface we expect our database implementation to follow
type Users interface {
	Authenticate(id string) error
	CreateUser(id int32, pwdHash string, name string, age int, addInfo string) (User, error)
	UpdateUser(usr User) error
	GetUser(id int32) (User, error)
	DeleteUser(id int32) error
}

// Service - our Users service, responsible for updating the Users DB
type Service struct {
	User Users
}

// New - returns a new instance of our Users service
func New(user Users) Service {
	return Service{
		User: user,
	}
}

// GetUser - retrieves a user based on the ID from the
func (s Service) GetUser(ctx context.Context, id int32) (User, error) {
	usr, err := s.User.GetUser(id)
	if err != nil {
		return User{}, err
	}
	return usr, nil
}

// InsertRocket - inserts a new rocket into the store
func (s Service) CreateUser(ctx context.Context, id int32, pwdHash string, name string, age int, addInfo string) (User, error) {
	usr, err := s.User.CreateUser(id, pwdHash, name, age, addInfo)
	if err != nil {
		return User{}, err
	}
	return usr, nil
}

// DeleteRocket - deletes a rocket from our inventory
func (s Service) DeleteRocket(ctx context.Context, id int32) error {
	err := s.User.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
