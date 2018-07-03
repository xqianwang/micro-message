package store

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	createMessage(string) (int64, error)
	getMessages() ([]Message, error)
	deleteMessage(int) error
	getMessageByID(int) (*Message, error)
	createUser(*User) error
	getUserByEmail(string) (*User, error)
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
		panic("PGHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("PGPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("PGUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("PGPASS environment variable required but not set")
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
func (s dataStore) createMessage(content string) (int64, error) {
	var id int
	//here we will trigger func to judge if message is palindrome or not
	message := Message{Content: content}
	pl := message.checkPalindrome()
	createMessage := `INSERT INTO message(content, palindrome) VALUES ($1, $2) RETURING message.id`
	err := s.db.QueryRow(createMessage, message.Content, pl).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

func (s dataStore) getMessages() ([]Message, error) {
	var messages = []Message{}
	getMessages := `SELECT * FROM message`
	err := s.db.Select(&messages, getMessages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (s dataStore) getMessageByID(id int) (*Message, error) {
	var message = Message{}
	getMessage := `SELECT * FROM message WHERE ID=$1`
	rows, err := s.db.Queryx(getMessage, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&message)
		if err != nil {
			return nil, err
		}
	}
	return &message, nil
}

func (s dataStore) deleteMessage(id int) error {
	deleteMessage := `DELETE FROM message WHERE ID=:id`
	_, err := s.db.NamedQuery(deleteMessage, id)
	if err != nil {
		return err
	}
	return nil
}

func (s dataStore) createUser(user *User) error {
	var id int
	createUserQuery := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING users.id`
	err := s.db.QueryRow(createUserQuery, user.UserName, user.Password, user.Email).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (s dataStore) getUserByEmail(email string) (*User, error) {
	var user = User{}
	getUserByEmail := `Select * FROM users WHERE email=$1`
	rows, err := s.db.Queryx(getUserByEmail, email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
