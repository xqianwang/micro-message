package store

import (
    "errors"
)

type user struct {
    UserName string `db:"username" json:"username"`
    Password string `db:"password" json:""password`
    Email    string `db:"email" json:"email"`
}

func registerUser(username, password, email string) (*user, error) {
    return nil, errors.New("placeholder error")
}

func userExists(email string) bool {
    return false
}
