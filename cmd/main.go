package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	dsn := flag.String("dsn", "host=PLACEHOLDER user=gopherbase password=PLACEHOLDER dbname=gopherbasedb port=5432 sslmode=disable TimeZone=Asia/Shanghai", "Connection string FLAG. Please edit the placeholders.")
	// to-do: improve this, cause it's lame.
	flag.Parse()

	db, err := dbConnect(*dsn)
	if err != nil {
		log.Fatalf("Can't open database due to: %v", err)
	}
	db.AutoMigrate(&BambooInfo{})

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheck)
	r.HandleFunc("/articles", ReadEndpoint).Methods("GET")
	r.HandleFunc("/articles", WriteArticle).Methods("POST")

	log.Println("Starting server on :8000")
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Server failed to start due to: %v", err)
	}
}
