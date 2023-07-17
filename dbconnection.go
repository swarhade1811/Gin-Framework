package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func setupDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/employee")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func createTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS feed_configurations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		feed_name VARCHAR(255),
		feed_uuid VARCHAR(255),
		file_source_name VARCHAR(255),
		feed_index_name VARCHAR(255),
		targets VARCHAR(255),
		call_minutes INT,
		tags VARCHAR(255)
	)`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
