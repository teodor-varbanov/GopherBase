package main

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "HealthCheck status: OK")
}

func WriteArticle(w http.ResponseWriter, r *http.Request) {

}

func ReadArticle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Response from ReadArticle")
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {

}
