package models

import (
	"errors"
)

//Message is the structure for messages sent by users
type Message struct {
	ID         int    `db:"id" json:"id"`
	Content    string `db:"content" json:"content"`
	Palindrome bool   `db:"palindrome" json:"palindrome"`
}

var messageList = []Message{
	Message{ID: 1, Content: "Qlik message center1."},
	Message{ID: 2, Content: "Qlik message center2."},
}

//GetAllMessages return a list of all the messages
func GetAllMessages() []Message {
	return messageList
}

//GetMessageByID will get message accoording to message id
func GetMessageByID(id int) (*Message, error) {
	for _, message := range messageList {
		if message.ID == id {
			return &message, nil
		}
	}

	return nil, errors.New("Cannot find the message")
}

//CheckPalindrome check if message that user posted is palidrome or not
func (m *Message) CheckPalindrome() (b bool) {
	if m.Palindrome == false && m.Content != "" {
		mid := len(m.Content) / 2
		last := len(m.Content) - 1
		for i := 0; i < mid; i++ {
			if m.Content[i] != m.Content[last-i] {
				return false
			}
		}
		return true
	}
	return true
}

//IsEmpty check if message is an empty object or not
func (m *Message) IsEmpty() (b bool) {
	if m.ID == 0 && m.Content == "" && m.Palindrome == false {
		return true
	}
	return false
}
