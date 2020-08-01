package main

import (
	// "database/sql"
	// "fmt"

	// _ "github.com/go-sql-driver/mysql"
	"ecommerce/cmd"
)

func main() {
	// fmt.Println("My ecommerce project.")
	// db, err := sql.Open("mysql", "root:password9@tcp(127.0.0.1:3306)/testdb")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer db.Close()

	// fmt.Println("Successfully connected to mysql database")
	cmd.Execute()
}
