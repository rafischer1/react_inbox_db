package models

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/subosito/gotenv"
)

// Message the psql table messages
type Message struct {
	ID        int64  `json:"id"`
	Read      bool   `json:"read"`
	Starred   bool   `json:"starred"`
	Selected  bool   `json:"selected"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Labels    string `json:"labels"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// GetAllMessages function
func GetAllMessages() []Message {
	connStr := dbInit()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("db:%v", db)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM messages")

	defer rows.Close()

	var messages []Message
	for rows.Next() {
		message := Message{}
		// gotta get all the fields!
		rows.Scan(&message.ID, &message.Read, &message.Starred, &message.Selected, &message.Subject, &message.Body, &message.Labels, &message.CreatedAt, &message.UpdatedAt)
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
	connStr := dbInit()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query(`SELECT * FROM messages where id =` + id)

	var entry []Message
	for row.Next() {
		message := Message{}
		// gotta get all the fields!
		row.Scan(&message.ID, &message.Read, &message.Starred, &message.Selected, &message.Subject, &message.Body, &message.Labels, &message.CreatedAt, &message.UpdatedAt)
		entry = append(entry, message)
	}
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}
	return entry
}

// PostMessage function
func PostMessage(req io.ReadCloser) []Message {
	fmt.Println("in pOSTmessages:", req)
	connStr := dbInit()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	var message []Message
	rows, err := db.Query(
		`INSERT INTO messages(ID, Read, Starred, Selected, Subject, Body, Labels, CreatedAt, UpdatedAt) VALUES`,
		message,
	)
	defer rows.Close()
	return message
}

// EditMessage function
func EditMessage() []Message {
	fmt.Println("In the model edit")
	connStr := dbInit()
	db, err := sql.Open("postgres", connStr)
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
	connStr := dbInit()
	db, err := sql.Open("postgres", connStr)
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

/************
*  initialize connection string
*  for db using .env
**************/
func dbInit() string {
	gotenv.Load()
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	connStr := fmt.Sprintf("user=%[1]v "+
		"dbname=%[2]v sslmode=disable", dbuser, dbname)
	return connStr
}
