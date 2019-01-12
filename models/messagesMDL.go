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

	fmt.Println("in the get all messages db:", db)

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
	// fmt.Println("messages out from DB get all:", messages)
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
	defer row.Close()
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
func PostMessage(Subject string, Body string, Labels string) error {
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	message := Message{}
	var entry []Message
	//Create
	errTwo := db.QueryRow(`INSERT INTO messages(read, starred, selected, subject, body, labels) VALUES($1, $2, $3, $4, $5, $6) RETURNING *`, false, false, false, Subject, Body, Labels).Scan(&message.ID, &message.Read, &message.Starred, &message.Selected, &message.Subject, &message.Body, &message.Labels)
	entry = append(entry, message)
	if errTwo != nil {
		return errTwo
	}

	fmt.Printf("Last inserted ID: %v\n", message)
	return errTwo
}

// EditMessage function
func EditMessage(ID int, Body string, Read bool) ([]Message, error) {
	fmt.Println("In the model edit id and body:", ID, Body, Read)

	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if len(Body) == 0 {
		Body = ""
	}

	message := Message{}
	var entry []Message
	sqlStatement := `UPDATE messages SET read = $3, labels = $2 WHERE id = $1 RETURNING id, labels, read;`

	err = db.QueryRow(sqlStatement, ID, Body, Read).Scan(&message.ID, &message.Labels, &message.Read)
	if err != nil {
		panic(err)
	}
	entry = append(entry, message)

	return entry, err
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
	defer row.Close()
	return id
}
