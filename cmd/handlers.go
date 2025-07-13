package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "HealthCheck status: OK")
}

func WriteArticle(w http.ResponseWriter, r *http.Request) {

}

func ReadEndpoint(w http.ResponseWriter, r *http.Request) {

	url := flag.String("url", "https://bambooo.free.beeceptor.com/api/latest/result/bbd-ddb/latest", "URL of API resource to GET info from")
	flag.Parse()
	info, err := read(*url)
	if err != nil {
		log.Fatalf("Unable to read API endpoint")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {

}
