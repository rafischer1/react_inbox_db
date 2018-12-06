package models

import (
	"database/sql"
	"fmt"
	"log"

	d "github.com/rafischer1/react_inbox_db/db"
)

// File the psql table files
type File struct {
	ID          int64  `json:"id"`
	CoolNotCool bool   `json:"coolNotCot"`
	Body        string `json:"body"`
}

// GetAllFilesMDL function
func GetAllFilesMDL() []File {
	db, err := sql.Open("postgres", d.ConnStr)
	fmt.Println("connstr:", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM files;")
	fmt.Printf("rows:%+v", rows)
	defer rows.Close()

	var files []File
	for rows.Next() {
		file := File{}
		// gotta get all the fields!
		rows.Scan(&file.ID, &file.CoolNotCool, &file.Body)
		files = append(files, file)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("files getAll:", files)
	return files
}

// GetOneMessage Selects by ID fom db
func GetOneFileMDL(id string) File {
	fmt.Println("In the get one model", id)
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query(`SELECT * FROM files where id =` + id)

	var entry []File
	for row.Next() {
		file := File{}
		// gotta get all the fields!
		row.Scan(&file.ID, &file.CoolNotCool, &file.Body)
		entry = append(entry, file)
	}
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}
	return entry[0]
}

// PostFileMDL function
func PostFileMDL(body File) []File {
	fmt.Println("in POSTmessages:", body)
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	var file []File
	rows, err := db.Query(
		`INSERT INTO messages(ID, Read, Starred, Selected, Subject, Body, Labels, CreatedAt, UpdatedAt) VALUES`,
		body,
	)
	defer rows.Close()
	return file
}

// EditMessage function
func EditFileMDL() bool {
	fmt.Println("In the model edit")
	return true
}

// DeleteFileMDL  function
func DeleteFileMDL(id string) string {
	fmt.Println("In the delete file model", id)
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row, err := db.Query(`Delete FROM files where id =` + id)
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}
	return id
}
