package store

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
//TestUserExists tests user exists function
func TestUserExists(t *testing.T) {
    MockData()
	if !userExists("user1@qlik.com") {
		t.Fail()
	}
	if userExists("test@qlik.com") {
		t.Fail()
	}
	RegisterUser("test", "test", "test@qlik.com")
	if !userExists("user1@qlik.com") {
		t.Fail()
	}
    CleanData()
	//func clear users data
}
//TestUserRegistrationValid test user registration function
func TestUserRegistrationValid(t *testing.T) {
    //insert users data
    MockData()
    //existing user user1
	u, err := RegisterUser("user1", "haha", "user1@qlik.com")
    assert.NotNil(t, err, "Error should not be nil")
    assert.Nil(t, u, "Returned user should be nil")
    //new user test
    u, err = RegisterUser("test", "haha", "test@qlik.com")
    assert.Nil(t, err, "Error should not be nil")
    assert.NotNil(t, u, "Returned user should be nil")
    CleanData()
	//func clean users data
}
