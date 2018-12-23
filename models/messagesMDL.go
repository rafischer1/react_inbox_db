package models

import (
	"database/sql"
	"fmt"
	"log"

	d "github.com/rafischer1/react_inbox_db/db"
)

// Message the psql table messages
type Message struct {
	ID       int    `json:"id"`
	Read     bool   `json:"read"`
	Starred  bool   `json:"starred"`
	Selected bool   `json:"selected"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Labels   string `json:"labels"`
}

// GetAllMessages function
func GetAllMessages() []Message {
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	rows, err := db.Query("SELECT * FROM messages")

	defer rows.Close()

	var messages []Message

	for rows.Next() {
		message := Message{}

		// gotta get all the fields!
		rows.Scan(&message.ID, &message.Read, &message.Starred, &message.Selected, &message.Subject, &message.Body, &message.Labels)
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return messages
}

// GetOneMessage Selects by ID fom db
func GetOneMessage(id string) []Message {
	fmt.Println("In the get one model", id)

	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query(`SELECT * FROM messages where id =` + id)

	var entry []Message

	for row.Next() {
		message := Message{}
		// gotta get all the fields!
		row.Scan(&message.ID, &message.Read, &message.Starred, &message.Selected, &message.Subject, &message.Body, &message.Labels)
		entry = append(entry, message)
	}

	if err := row.Err(); err != nil {
		log.Fatal(err)
	}

	return entry
}

// PostMessage function
func PostMessage(Read bool, Starred bool, Selected bool, Subject string, Body string, Labels string) (int, string, error) {
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	//Create
	var messageID int
	var messageSubject string
	errTwo := db.QueryRow(`INSERT INTO messages(read, starred, selected, subject, body, labels) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`, Read, Starred, Selected, Subject, Body, Labels).Scan(&messageID, &messageSubject)

	if errTwo != nil {
		return 0, "undefined", errTwo
	}

	fmt.Printf("Last inserted ID: %v\n", messageID)
	return messageID, messageSubject, errTwo
}

// EditMessage function
func EditMessage() []Message {
	fmt.Println("In the model edit")

	db, err := sql.Open("postgres", d.ConnStr)
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
func DeleteMessage(id string) string {
	fmt.Println("In the delete model", id)

	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row, err := db.Query(`Delete FROM messages where id =` + id)
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}
	return id
}
