package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	const (
		host     = ""
		port     = ""
		user     = ""
		password = ""
		dbname   = ""
	)

	fmt.Println("[STATUS] Establishing connection to database.")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("[STATUS] Successfully connected!")
	fmt.Println("[STATUS] Creating club table.")

	createClubTable := `CREATE TABLE IF NOT EXISTS clubs(id int PRIMARY KEY, name text)`

	ctx, cancellation := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancellation()

	_, err = db.ExecContext(ctx, createClubTable)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
	}

	fmt.Println("[STATUS] Club table created.")
	fmt.Println("[STATUS] Closing database connection")
	err = db.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("[STATUS] Database connection closed.")
}
