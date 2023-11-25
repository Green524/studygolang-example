package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // enable support for MySQL
	_ "github.com/lib/pq"              // enable support for Postgres
)

var (
	db, err = sql.Open("postgres", "dbname") // OK
	//db, err = sql.Open("mysql", "dbname")    // OK
	//db, err = sql.Open("sqlite3", "dbname")  // returns error: unknown driver "sqlite3"
)

func main() {
}
