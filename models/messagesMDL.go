package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Message struct {
	id    uint   `json:"id"`
	title string `json:"title"`
}

/***********
* Get All
*
*************/
func GetAllMessage() []Message {
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
		rows.Scan(&message.id, &message.title)
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return messages
}

/***********
* PostMessage
*
*************/

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

/***********
* Edit
*
*************/

func EditMessagek() []Message {
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
		rows.Scan(&message.id, &message.title)
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return messages
}

/***********
* DELETE
*
*************/

// func DeleteMessage() []Message {
// 	fmt.Println("In the model delete")
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
