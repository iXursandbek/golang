package main

import (
	"fmt"
	"instat/handlers"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBStore struct {
	db *sqlx.DB
}

func main() {
	// fs := http.FileServer(http.Dir("assets"))
	// http.Handle("/assets/", http.StripPrefix("/assets", fs))

	db, err := sqlx.Connect("postgres", "dbname=instat user=postgres password=x3127106 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	http.HandleFunc("/", handlers.RunIndex)
	http.HandleFunc("/contact", handlers.Contact)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Successfully connected!")
}
