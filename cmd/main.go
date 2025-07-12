package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      uint
	Title   string
	Content string
}

func main() {
	dsn := flag.String("dsn", "host=PLACEHOLDER user=gopherbase password=PLACEHOLDER dbname=gopherbasedb port=5432 sslmode=disable TimeZone=Asia/Shanghai", "Connection string FLAG. Please edit the placeholders.")
	// to-do: improve this, cause it's lame.
	flag.Parse()

	db, err := dbConnect(*dsn)
	if err != nil {
		log.Fatalf("Can't open database due to: %v", err)
	}
	db.AutoMigrate(&Article{})

	id := uint(1)
	title := "Bamboo shit"
	content := "Content"
	test := Article{
		ID:      id,
		Title:   title,
		Content: content,
	}

	db.Create(&test)
	// afterwards call db.Delete(&test) to clean up

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheck)
	r.HandleFunc("/articles", ReadArticle).Methods("GET")
	r.HandleFunc("/articles", WriteArticle).Methods("POST")

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Server failed to start due to: %v", err)
	}
}
