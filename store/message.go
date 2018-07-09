package store

//Message is the structure for messages sent by users
type Message struct {
	ID         int    `db:"id" json:"id"`
	Title      string `db:"title" json:"title"`
	Content    string `db:"content" json:"content"`
	Palindrome bool   `db:"palindrome" json:"palindrome"`
}

// var messageList = []Message{
// 	Message{ID: 1, Content: "Qlik message center1."},
// 	Message{ID: 2, Content: "Qlik message center2."},
// }

//GetAllMessages return a list of all the messages
func GetAllMessages() ([]Message, error) {
	messageList, err := DBStore.getMessages()
	if err != nil {
		return nil, err
	}
	return messageList, nil
}

//GetMessageByID will get message accoording to message id
func GetMessageByID(id int) (*Message, error) {
	message, err := DBStore.getMessageByID(id)
	if err != nil {
		return nil, err
	}

	return message, nil
}

//CreateMessage create a message that a user posted
func CreateMessage(title, content string) (int64, error) {
	id, err := DBStore.createMessage(title, content)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//DeleteMessage deletes a message
func DeleteMessage(id int) error {
	err := DBStore.deleteMessage(id)
	if err != nil {
		return err
	}
	return nil
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
