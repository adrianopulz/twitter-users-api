package users

// UsersService interface variable.
var UsersService usersServiceInterface = &usersService{}

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*User, *ErrorMsg)
	SearchUsers(string) (*Users, *ErrorMsg)
}

// GetUser returns a User by ID.
func (s *usersService) GetUser(userID int64) (*User, *ErrorMsg) {
	dao := &User{ID: userID}
	user, err := dao.Get()

	if err != nil {
		return nil, err
	}

	return user, nil
}

// SearchUsers return users by the user_name.
func (s *usersService) SearchUsers(userName string) (*Users, *ErrorMsg) {
	dao := &User{UserName: userName}
	users, err := dao.SearchUsers()

	if err != nil {
		return nil, err
	}

	return &users, nil
}
