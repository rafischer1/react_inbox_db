package models

import (
	"database/sql"
	"fmt"
	"log"
)

// Message the psql table messages
type Message struct {
	ID        int      `json:"id"`
	Read      bool     `json:"read"`
	Starred   bool     `json:"starred"`
	Selected  bool     `json:"selected"`
	Subject   string   `sql:"type:varchar(255)"`
	Body      string   `json:"body"`
	Labels    []string `sql:",array"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// GetAllMessages function
func GetAllMessages() []Message {
	connStr := "user=artiefischer dbname=reactinboxdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM messages")

	defer rows.Close()
	var messages []Message
	for rows.Next() {
		message := Message{}
		rows.Scan(&message.ID, &message.Read, &message.Starred, &message.Selected, &message.Subject, &message.Body, &message.Labels)
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return messages
}

// PostMessage function
func PostMessage() []Message {

	db, err := sql.Open("postgres", "user=artiefischer dbname=reactinboxdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	var message []Message
	rows, err := db.Query(
		"INSERT INTO messages (title) VALUES ($1)",
		message,
	)
	defer rows.Close()

	// fmt.Println("In the model", message)
	// defer rows.Close()
	// var messages []Message
	// for rows.Next() {
	//   message := Message{}
	//   rows.Scan(&message.id, &message.title)
	//   messages = append(messages, message)
	// }
	// if err := rows.Err(); err != nil {
	//   log.Fatal(err)
	// }

	return message
}

// EditMessage function
func EditMessage() []Message {
	fmt.Println("In the model edit")
	db, err := sql.Open("postgres", "user=artiefischer dbname=reactinboxdb sslmode=disable")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM messages")

	defer rows.Close()
	var messages []Message
	for rows.Next() {
		message := Message{}
		rows.Scan(&message.ID, &message.Subject)
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return messages
}

// DeleteMessage Model function
func DeleteMessage() []Message {
	var messages []Message
	fmt.Println("In the model delete")
	return messages
}

// 	db, err := sql.Open("postgres", "user=artiefischer dbname=reactinboxdb sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}

// 	rows, err := db.Query("SELECT * FROM messages")

// 	defer rows.Close()
// 	var messages []Message
// 	for rows.Next() {
// 		message := Message{}
// 		rows.Scan(&Message.id, &Message.title)
// 		messages = append(messages, Message)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	return messages
// }
