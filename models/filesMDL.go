package models

import (
	"database/sql"
	"fmt"
	"log"

	d "github.com/rafischer1/react_inbox_db/db"
)

// File the psql table files
type FileWithId struct {
	*File
	ID int `json:"id"`
}

// File is cool
type File struct {
	CoolNotCool bool   `json:"coolNotCool"`
	Body        string `json:"body"`
}

// GetAllFilesMDL function
func GetAllFilesMDL() []FileWithId {
	db, err := sql.Open("postgres", d.ConnStr)
	fmt.Println("connstr:", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM files;")
	fmt.Printf("rows:%+v", rows)
	defer rows.Close()

	var files []FileWithId
	for rows.Next() {
		file := FileWithId{}
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
func GetOneFileMDL(id string) FileWithId {
	fmt.Println("In the get one model", id)
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query(`SELECT * FROM files where id =` + id)

	var entry []FileWithId
	for row.Next() {
		file := FileWithId{}
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
func PostFileMDL(body File) []FileWithId {
	fmt.Println("in POST File:", body)
	db, err := sql.Open("postgres", d.ConnStr)
	// fmt.Println("connstr:", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var postFile []FileWithId
	var CoolField = "false"
	if body.CoolNotCool == true {
		CoolField = "true"
	}

	// queryStr := fmt.Sprintf("INSERT INTO files (ID, CoolNotCool, Body) VALUES (%[1]v%[2]v %[3]v", 0, CoolField, body.Body)

	// queryStr := "INSERT INTO files (ID, CoolNotCool, Body) VALUES (5, true, 'hi')"
	queryStr := fmt.Sprintf("INSERT INTO files (CoolNotCool, Body) VALUES (%v, '%v')", CoolField, body.Body)

	fmt.Println("querStr:", queryStr)

	_, err = db.Exec(queryStr)
	if err != nil {
		panic(err)
	}
	// defer rows.Close()

	return postFile
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
