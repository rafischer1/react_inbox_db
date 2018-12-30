package db

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
)

/************
*  initialize connection string
*  for db using .env
**************/
var ConnStr = DBInit()

// DBInit initializes the databasw using .env vars
//FOR DEVELEOPMENT ONLY !!!!
func DBInit() string {
	gotenv.Load()
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	ConnStr := fmt.Sprintf("user=%[1]v "+
		"dbname=%[2]v sslmode=disable", dbuser, dbname)
	return ConnStr
}

// FOR PRODUCTION BUILD ONLY!!!!
// func DBInit() string {
// 	gotenv.Load()
// 	url := os.Getenv("DATABASE_URL")
// 	ConnStr := url
// 	return ConnStr
// }
