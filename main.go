package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/rafischer1/react_inbox_db/handlers"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main() {
	initDb()
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/messages", handlers.GetAll).Methods("GET")
	router.HandleFunc("/messages", handlers.PostMessage).Methods("POST")
	router.HandleFunc("/messages/:id", handlers.EditMessage).Methods("PUT")
	router.HandleFunc("/messages/:id", handlers.DeleteMessage).Methods("DELETE")

	// set router
	router.Handle("/", http.FileServer(http.Dir("./static/")))
	// set listen port
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func initDb() {
	// grab .env variables
	gotenv.Load()
	// call dbConfig function to set env variables
	config := dbConfig()
	var err error
	// Loaded database info
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])
	// print out database information for development
	fmt.Println("db init info:", psqlInfo)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func dbConfig() map[string]string {

	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

// type Specification struct {
//     Debug       bool
//     Port        int
//     User        string
//     Users       []string
//     Rate        float32
//     Timeout     time.Duration
//     ColorCodes  map[string]int
// }

// func main() {
//     var s Specification
//     err := envconfig.Process("myapp", &s)
//     if err != nil {
//         log.Fatal(err.Error())
//     }
//     format := "Debug: %v\nPort: %d\nUser: %s\nRate: %f\nTimeout: %s\n"
//     _, err = fmt.Printf(format, s.Debug, s.Port, s.User, s.Rate, s.Timeout)
//     if err != nil {
//         log.Fatal(err.Error())
//     }

//     fmt.Println("Users:")
//     for _, u := range s.Users {
//         fmt.Printf("  %s\n", u)
//     }

//     fmt.Println("Color codes:")
//     for k, v := range s.ColorCodes {
//         fmt.Printf("  %s: %d\n", k, v)
//     }
// }
