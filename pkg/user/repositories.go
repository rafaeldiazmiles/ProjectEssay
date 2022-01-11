//go:generate mockgen -destination=repositories_mocks_test.go -package=rocket https://github.com/rafaeldiazmiles/ProjectEssay Users
package user

// User - should contain the definition of our user
type User struct {
	id      int32
	pwdHash string
	name    string
	age     int
	addInfo string
	// parents []Parent   --> Para implementar cuando haya parents
}

// // Service - our Users service, responsible for updating the Users DB
// type Service struct {
// 	User Users
// }

// New - returns a new instance of our Users service
// func New(user Users) Service {
// 	return Service{
// 		User: user,
// 	}
// }

// // GetUser - retrieves a user based on the ID from the
// func (s Service) GetUser(ctx context.Context, id int32) (User, error) {
// 	usr, err := s.User.GetUser(id)
// 	if err != nil {
// 		return User{}, err
// 	}
// 	return usr, nil
// }

// // DeleteRocket - deletes a rocket from our inventory
// func (s Service) DeleteRocket(ctx context.Context, id int32) error {
// 	err := s.User.DeleteUser(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
