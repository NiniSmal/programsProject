package main

import (
	"database/sql"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "postgres")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	storage := NewStorage(db)
	handler := NewHandler(storage)

	router := http.NewServeMux()

	router.HandleFunc("GET /programs", handler.ProgramsByProjectID)
	router.HandleFunc("GET /programs/id", handler.ProgramByID)
}
