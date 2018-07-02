package store

import (
    "testing"
)

func TestUserExists(t *testing.T) {
    if userExists("user1@qlik.com") {
        t.Fail()
    }

    if !userExists("xqian@qlik.com") {
        t.Fail()
    }

    registerUser("user1", "haha", "user1@qlik.com") 

    if !userExists("user1@qlik.com") {
        t.Fail()
    }

    //func clear users data
}

func TestUserRegistrationValid(t *testing.T) {
    //insert users data

    u, err := registerUser("user1", "haha", "user1@qlik.com")

    if err != nil || u != nil {
        t.Fail()
    }

    //func clean users data
}
