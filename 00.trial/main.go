package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	log.SetFlags(0)
}

func main() {
	// ------------------------------
	// Open
	//
	var (
		db  *sql.DB
		err error
	)

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db0")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	// ------------------------------
	// Fetch rows
	//
	var (
		rows *sql.Rows
	)

	rows, err = db.Query("SELECT * FROM tableA")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			name string
		)

		err = rows.Scan(&id, &name)
		if err != nil {
			log.Panic(err)
		}

		log.Printf("%d:%s", id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Panic(err)
	}

	// ------------------------------
	// Get current max id
	//
	var (
		maxId int
		row   *sql.Row
	)

	row = db.QueryRow("SELECT MAX(id) FROM tableA")
	err = row.Scan(&maxId)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("MAX(id)=%d", maxId)

	// ------------------------------
	// Insert next id
	//
	var (
		tx     *sql.Tx
		nextId = maxId + 1
	)

	tx, err = db.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO tableA (id, name) values (?, ?)", nextId, fmt.Sprintf("gitpod%d", nextId))
	if err != nil {
		log.Panic(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}
