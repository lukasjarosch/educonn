package main

import (
	"github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id        string         `db:"id"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt mysql.NullTime `db:"deleted_at" `
}
