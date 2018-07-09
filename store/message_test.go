package store

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//test palindrome function
func TestCheckPalindrome(t *testing.T) {
	var testString = []Message{Message{Content: "qlik"}, Message{Content: "ahha"}}
	var results = make([]bool, 2)

	for i, str := range testString {
		results[i] = str.CheckPalindrome()
	}

	if results[0] != false && results[1] != true {
		t.Fail()
	}
}

//test get all messages function
func TestGetAllMessages(t *testing.T) {
	MockData()
	messages, err := GetAllMessages()
	CleanData()
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, messages, "Messages should be returned")
}

//test get one message by id function
func TestGetMessageByID(t *testing.T) {
	MockData()
	message, err := GetMessageByID(1)
	CleanData()
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, message, "Message should not be nil")
}

//test create message function
func TestCreateMessage(t *testing.T) {
	id, err := CreateMessage("test", "test create message")
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, id, "Message id should be larger than 0")
	CleanData()
}

//test delete message function
func TestDeleteMessage(t *testing.T) {
	MockData()
	err := DeleteMessage(1)
	CleanData()
	assert.Nil(t, err, "Error should be nil.")

}

var ds dataStore
var schema = `
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username varchar,
    password varchar,
    email varchar
);

CREATE TABLE IF NOT EXISTS message (
    id SERIAL PRIMARY KEY,
    content text,
    palindrome boolean
)`

func init() {
	config := dbConfig()
	pgInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])
	c, err := sqlx.Connect("postgres", pgInfo)
	if err != nil {
		panic(err)
	}
	ds = dataStore{db: c}
}

//MockData mocks data for tests
func MockData() {
	//create tables
	ds.db.MustExec(schema)
	insertMessage := `INSERT INTO message (title, content, palindrome) VALUES ($1, $2, $3)`
	insertUser := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	ds.db.MustExec(insertMessage, "test message", "test message", false)
	ds.db.MustExec(insertMessage, "test more", "test more", false)
	ds.db.MustExec(insertUser, "user1", "pass1", "user1@qlik.com")
	ds.db.MustExec(insertUser, "user2", "pass2", "user2@qlik.com")
}

//CleanData clean data after tests
func CleanData() {
	deleteMessage := `TRUNCATE TABLE message RESTART IDENTITY`
	deleteUser := `TRUNCATE TABLE users RESTART IDENTITY`
	ds.db.MustExec(deleteMessage)
	ds.db.MustExec(deleteUser)
}
