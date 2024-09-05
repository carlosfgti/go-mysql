package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "username:userpass@tcp(127.0.0.1:3306)/goproject"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Print("Successfully connected!")

	selectRows(db)
}

func selectRows(db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}
}
