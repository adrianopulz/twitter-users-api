package users

import (
	"fmt"
	"net/http"

	"github.com/adrianopulz/twitter-users-api/repository/mysql/users_db"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryListUser               = "SELECT	id, user_name, first_name, last_name FROM users WHERE status = 1 and user_name LIKE ?;"
	queryGetUser                = "SELECT id, user_name, email, first_name, last_name, photo, created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, created, status FROM users WHERE email=? AND password=? AND status=?"
)

// UsersRepository interface variable.
var UsersRepository usersRepositoryInterface = &usersRepository{}

type usersRepository struct{}

type usersRepositoryInterface interface {
	Get(*User) *ErrorMsg
	ListUsers(userName string) ([]User, *ErrorMsg)
}

// Get a user by ID.
func (r *usersRepository) Get(u *User) *ErrorMsg {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		fmt.Println(err)
		return &ErrorMsg{"Error when tying to prepare the query", http.StatusInternalServerError}
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.ID)

	if getErr := result.Scan(&u.ID, &u.UserName, &u.Email, &u.FirstName, &u.LastName, &u.Photo, &u.Created, &u.Status); getErr != nil {
		fmt.Println(getErr)
		return &ErrorMsg{"Error when tying to get user", http.StatusInternalServerError}
	}

	return nil
}

// ListUsers return a list of user.
func (r *usersRepository) ListUsers(name string) ([]User, *ErrorMsg) {
	stmt, err := users_db.Client.Prepare(queryListUser)
	if err != nil {
		fmt.Println(err)
		return nil, &ErrorMsg{"Error when tying to prepare the query", http.StatusInternalServerError}
	}
	defer stmt.Close()

	rows, err := stmt.Query(name + "%")
	if err != nil {
		fmt.Println(err)
		return nil, &ErrorMsg{"error when tying to get users", http.StatusInternalServerError}
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName); err != nil {
			fmt.Println(err)
			return nil, &ErrorMsg{"error when tying to gett user", http.StatusInternalServerError}
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, &ErrorMsg{fmt.Sprintf("No users matching the user name %s", name), http.StatusNotFound}
	}

	return results, nil
}
