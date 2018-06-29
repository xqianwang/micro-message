package store

import (
	"errors"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

const (
	dbhost = "PGHOST"
	dbport = "PGPORT"
	dbuser = "PGUSER"
	dbpass = "PGPASS"
	dbname = "DBNAME"
)

//Store interface defines 4 methods
//CreateMessage, GetMessages, DeleteMessage, GetMessageByID
//this will interact with message API to retrieve data from database
type Store interface {
	createMessage(message *Message) error
	getMessages() ([]Message, error)
	deleteMessage(int) error
	getMessageByID(int) (*Message, error)
}

//dataStore maintain db connection
type dataStore struct {
	db *sqlx.DB
}

//DBStore singleton db instance
var DBStore Store

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
	DBStore = dataStore{db: c}
}

//Get database configuration from envs
func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

//CreateMessage creates a message in database
func (s dataStore) createMessage(message *Message) error {
	//here we will trigger func to judge if message is palindrome or not
	pl := message.checkPalindrome()
	createMessage := `INSERT INTO message(content, palindrome) VALUES (?, ?)`
	s.db.MustExec(createMessage, message.Content, pl)
	return nil
}

func (s dataStore) getMessages() ([]Message, error) {
	var messages = []Message{}
	getMessages := `SELECT * FROM message`
	_ = s.db.Select(&messages, getMessages)
	if messages[0].IsEmpty() {
		return nil, errors.New("No messages")
	}
	return messages, nil
}

func (s dataStore) getMessageByID(id int) (*Message, error) {
	var message = Message{}
	getMessage := `SELECT * FROM message WHERE ID=$1`
	err := s.db.Select(&message, getMessage, id)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (s dataStore) deleteMessage(id int) error {
	deleteMessage := `DELETE FROM message WHERE ID:id`
	_, err := s.db.NamedQuery(deleteMessage, id)
	if err != nil {
		return err
	}
	return nil
}
