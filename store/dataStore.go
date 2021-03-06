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
	createMessage(string, string) (int64, error)
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
		fmt.Println("Using defualt postgresql host: localhost")
		host = "localhost"
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		fmt.Println("Using defualt postgresql port: 5432")
		port = "5432"
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		fmt.Println("Using defualt postgresql user: postgres")
		user = "postgres"
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		fmt.Println("Using defualt postgresql password: postgres")
		password = "postgres"
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		fmt.Println("Using defualt postgresql db name: postgres")
		name = "postgres"
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

//CreateMessage creates a message in database
func (s dataStore) createMessage(title, content string) (int64, error) {
	var id int
	//here we will trigger func to judge if message is palindrome or not
	message := Message{
		Title:   title,
		Content: content,
	}
	pl := message.CheckPalindrome()
	createMessage := `INSERT INTO message(title, content, palindrome) VALUES ($1, $2, $3) RETURNING message.id`
	err := s.db.QueryRow(createMessage, message.Title, message.Content, pl).Scan(&id)
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
	deleteMessage := `DELETE FROM message WHERE ID=$1`
	_, err := s.db.Exec(deleteMessage, id)
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
