package store

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestCheckPalindrome(t *testing.T) {
	var testString = []Message{Message{Content: "qlik"}, Message{Content: "ahha"}}
	var results = make([]bool, 2)

	for i, str := range testString {
		results[i] = str.checkPalindrome()
	}

	if results[0] != false && results[1] != true {
		t.Fail()
	}
}

func TestGetAllMessages(t *testing.T) {
	mockData()
	messages, err := GetAllMessages()
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, messages, "Messages should be returned")
	cleanData()
}

func TestGetMessageByID(t *testing.T) {

}

var ds dataStore
var schema = `
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username varchar,
    password varchar,
    email varchar
);

CREATE TABLE message (
    id SERIAL PRIMARY KEY
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

func mockData() {
	//create tables
	ds.db.MustExec(schema)
	insertMessage := `INSERT INTO message (content, palindrome) VALUES ($1, $2)`
	insertUser := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	ds.db.MustExec(insertMessage, "test message", false)
	ds.db.MustExec(insertMessage, "test more", false)
	ds.db.MustExec(insertUser, "user1", "pass1", "user1@qlik.com")
	ds.db.MustExec(insertUser, "user2", "pass2", "user2@qlik.com")
}

func cleanData() {
	deleteMessage := `TRUNCATE TABLE message CASCADE`
	deleteUser := `TRUNCATE TABLE users CASCADE`
	ds.db.MustExec(deleteMessage)
	ds.db.MustExec(deleteUser)
}
