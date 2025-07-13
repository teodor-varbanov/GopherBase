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
	// mock api
	read(w, "https://bambooo.free.beeceptor.com/api/latest/result/bbd-ddb/latest")
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {

}
