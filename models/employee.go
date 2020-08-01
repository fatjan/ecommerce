package models

import (
	"database/sql"
	_ "database/sql"
	"ecommerce/db"
	"fmt"
)

type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}

var con *sql.DB

// get employees from database
func GetEmployee() Employees {
	con := db.CreateCon()
	sqlStatement := "SELECT * FROM employee ORDER BY id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}

		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary,
			&employee.Age, &employee.Gender, &employee.Status)

		if err2 != nil {
			fmt.Print(err2)
		}
		result.Employees = append(result.Employees, employee)
	}
	return result
}
