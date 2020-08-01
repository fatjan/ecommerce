package entity

import (
	_ "github.com/go-sql-driver/mysql"
)

// User Model
type User struct {
	ID          int    `xorm:"INT(11) 'id'"`
	Name        string `xorm:"VARCHAR(64) 'name'"`
	Email       string `xorm:"VARCHAR(64) 'email'"`
	PhoneNumber string `xorm:"VARCHAR(256) 'phonenumber'"`
}
