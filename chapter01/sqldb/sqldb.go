package main

import (
	"database/sql"
	"log"
	"os"
)

func main() {
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	db, err := sql.Open("postgres", "pgURL")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
