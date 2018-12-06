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
	id        int64  `json:"id"`
	read      bool   `json:"read"`
	starred   bool   `json:"starred"`
	selected  bool   `json:"selected"`
	subject   string `json:"subject"`
	body      string `json:"body"`
	labels    string `json:"labels"`
	createdAt string `json:"createdAt"`
	updatedAt string `json:"updatedAt"`
}

// GetAllMessages function
func GetAllMessages() []Message {
	connStr := dbInit()
	fmt.Println("connection string:", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("db:%v", db)
	defer db.Close()

	fmt.Println("db Query:", "Select * from messages")
	rows, err := db.Query("SELECT * FROM messages")

	defer rows.Close()

	var messages []Message
	for rows.Next() {
		message := Message{}
		// gotta get all the fields!
		rows.Scan(&message.id, &message.read, &message.starred, &message.selected, &message.subject, &message.body, &message.labels, &message.createdAt, &message.updatedAt)
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
	fmt.Println("id:", id)
	row, err := db.Query(`SELECT * FROM messages where id =` + id)

	var entry []Message
	for row.Next() {
		message := Message{}
		// gotta get all the fields!
		row.Scan(&message.id, &message.read, &message.starred, &message.selected, &message.subject, &message.body, &message.labels, &message.createdAt, &message.updatedAt)
		entry = append(entry, message)
	}
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}
	return entry
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
		rows.Scan(&message.id, &message.subject)
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return messages

}

// DeleteMessage Model function
func DeleteMessage(id string) []Message {
	fmt.Println("In the delete model", id)
	connStr := dbInit()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("id:", id)
	row, err := db.Query(`Delete FROM messages where id =` + id)

	var entry []Message
	for row.Next() {
		message := Message{}
		// gotta get all the fields!
		row.Scan(&message.id, &message.read, &message.starred, &message.selected, &message.subject, &message.body, &message.labels, &message.createdAt, &message.updatedAt)
		entry = append(entry, message)
	}
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}
	return entry
}

// initialize connection string for db using.env
func dbInit() string {
	gotenv.Load()
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	connStr := fmt.Sprintf("user=%[1]v "+
		"dbname=%[2]v sslmode=disable", dbuser, dbname)
	return connStr
}

func dbOpen() {
	connStr := dbInit()
	fmt.Println("connection string:", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("db:%v", db)
	defer db.Close()
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
