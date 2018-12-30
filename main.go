package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/rafischer1/react_inbox_db/handlers"
)

// db var references sql.DB
var db *sql.DB

// env variable declarations
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
	port := GetPort()
	r := mux.NewRouter()
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	srv := &http.Server{
		Addr:         dbport,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	r.HandleFunc("/messages", handlers.GetAll).Methods("GET")
	r.HandleFunc("/messages/{id}", handlers.GetOne).Methods("GET")
	r.HandleFunc("/messages", handlers.PostMessage).Methods("POST", "OPTIONS")
	r.HandleFunc("/messages/{id}", handlers.EditMessage).Methods("PUT", "OPTIONS")
	r.HandleFunc("/messages/{id}", handlers.DeleteMessage).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/files", handlers.GetAllFiles).Methods("GET")
	r.HandleFunc("/files/{id}", handlers.GetOneFile).Methods("GET")
	r.HandleFunc("/files", handlers.PostFile).Methods("POST")
	r.HandleFunc("/files/{id}", handlers.EditFile).Methods("PUT")
	r.HandleFunc("/files/{id}", handlers.DeleteFile).Methods("DELETE")

	// serve static files
	r.Handle("/", http.FileServer(http.Dir("static/")))

	// set router
	go func() {
		log.Println("Listening...3003")
		http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server stopped")
}

func initDb() {
	// grab .env variables using gotenv package
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
	// fmt.Println("db init info:", psqlInfo)

	// open and run the db
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

// basic config setup and error handling for db env variables
func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	// fmt.Println("host:", host)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	// fmt.Println("port:", port)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	// fmt.Println("user:", user)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	// fmt.Println("dbname:", name)
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

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5432"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
