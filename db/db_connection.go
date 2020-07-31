package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Create connection to database (mysql)
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root:password9@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is successfully connected.")
	}

	// defer db.Close
	// ensure that connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db failed to connect.")
		fmt.Println(err.Error())
	}
	return db
}
