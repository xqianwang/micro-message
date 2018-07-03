package store

import (
	"fmt"
	"strings"
)

//User for micro message users
type User struct {
	ID       int    `db:"id" json:"id"`
	UserName string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Email    string
}

//RegisterUser registers new user
func RegisterUser(username, password, email string) (*User, error) {
	if userExists(email) {
		return nil, fmt.Errorf("The email account %s already exists", email)
	}
	if strings.TrimSpace(password) == "" {
		return nil, fmt.Errorf("Password cannot be empty")
	}
	user := User{
		UserName: username,
		Password: password,
		Email:    email,
	}

	if err := DBStore.createUser(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

//UserExists check if user aleady exists based on email account
func userExists(email string) bool {
	user, _ := DBStore.getUserByEmail(email)
	if user != nil && user.UserName != "" && user.Password != "" && user.Email != "" {
		return true
	}
	return false
}

//ValidUser validates if user is available or not
func ValidUser(username, password, email string) bool {
	user, err := DBStore.getUserByEmail(email)
	if err != nil {
		return false
	}
	if user != nil && user.UserName == username && user.Password == password && user.Email == email {
		return true
	}
	return false
}
