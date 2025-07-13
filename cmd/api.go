package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func read(w http.ResponseWriter, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Couldn't read API due to: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status code error: %v %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Couldn't read the response due to: %v", err)
	}

	fmt.Fprint(w, string(data))
}
