package models

//Message is the structure for messages sent by users
type Message struct {
	ID      int    `db:"id" json:"id"`
	Content string `db:"content" json:"content"`
}

var messageList = []Message{
	Message{ID: 1, Content: "Qlik message center1."},
	Message{ID: 2, Content: "Qlik message center2."},
}

//GetAllMessages return a list of all the messages
func GetAllMessages() []Message {
	return messageList
}
