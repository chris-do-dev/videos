package main

import (
	"database/sql"
	"fmt"
)

func main() {
	readDatabase()
}

func readDatabase() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("caught a panic")
		}
	}()

	conn, _ := sql.Open("libsql", "file:"+filename)

	// will panic because conn is nil
	res, err := conn.ExecContext(ctx, statement, args...)
}
