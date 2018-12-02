package models

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/subosito/gotenv"
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
	connStr := dbInit()
	fmt.Println("connStr:", connStr)
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
	connStr := dbInit()
	db, err := sql.Open("postgres", connStr)
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
func DeleteMessage() []Message {
	var messages []Message

	connStr := dbInit()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("In the model delete", db)
	return messages
}

// initialize connection string for db using.env
func dbInit() string {
	gotenv.Load()
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	connStr := fmt.Sprintf("user=%[1]v dbname=%[2]v  sslmode=disable", dbuser, dbname)
	return connStr
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
