package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/nih13/button/pkg/model"
)

func main() {
	db := Initdb()
	query := model.New(db)

	user, err := query.GetUser(context.Background(), 1)

	fmt.Print(user, err)

}

func Initdb() *sql.DB {
	connStr := "postgres://postgres:postgres@localhost:5432/lootbox?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Database connection failed")
	}
	return db
}
