package users

import (
	"encoding/json"
	"net/http"
)

// User struct custom type.
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:user_name`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Photo     string `json:"photo"`
	Created   string `json:"date_created"`
	Status    int    `json:"status"`
}

// Users list of User.
type Users []User

// ErrorMsg to return HTTP errors trhu the layers.
type ErrorMsg struct {
	Msg  string
	Code int
}

// Get a user by Id.
func (u *User) Get() (*User, *ErrorMsg) {
	err := UsersRepository.Get(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// SearchUsers a user by Id.
func (u *User) SearchUsers() (Users, *ErrorMsg) {
	users, err := UsersRepository.ListUsers(u.UserName)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Marshal returns User as byte array.
func (u *User) Marshall() []byte {
	userJSON, err := json.Marshal(u)
	if err != nil {
		errObj := ErrorMsg{"Internal error!", http.StatusInternalServerError}
		err, _ := json.Marshal(errObj)
		return err
	}

	return userJSON
}

// Marshall returns a List of Users as byte array.
func (users Users) Marshall() []byte {
	usersJSON, err := json.Marshal(users)
	if err != nil {
		errObj := ErrorMsg{"Internal error!", http.StatusInternalServerError}
		err, _ := json.Marshal(errObj)
		return err
	}

	return usersJSON
}
