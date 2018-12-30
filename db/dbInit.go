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
func DBInit() string {
	gotenv.Load()
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	ConnStr := fmt.Sprintf("user=%[1]v "+
		"dbname=%[2]v sslmode=disable", dbuser, dbname)
	return ConnStr
}

// func DBInit() string {
// 	gotenv.Load()
// 	var ConnStr string
// 	dbname := os.Getenv("DBNAME")
// 	dbuser := os.Getenv("DBUSER")
// 	url := os.Getenv("DATABASE_URL")

// 	if url == "" {
// 		ConnStr := fmt.Sprintf("user=%[1]v "+
// 			"dbname=%[2]v sslmode=disable", dbuser, dbname)
// 		fmt.Println("Connstr for dev:", ConnStr)
// 	} else if url != "" {
// 		ConnStr := url
// 		fmt.Println("Connstr for prod:", ConnStr)
// 	}
// 	return ConnStr
// }
