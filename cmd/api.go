package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Struct fields need to be adjusted based on needs and real Bamboo, not on the mock up as it is now
type BambooInfo struct {
	BuildKey    string `json:"buildKey"`
	BuildNumber int    `json:"buildNumber"`
	Status      string `json:"Status"`
	BuildState  string `json:"buildState"`
}

func read(url string) (BambooInfo, error) {
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

	// fmt.Fprint(w, string(data))

	var bambooInfo BambooInfo
	err = json.Unmarshal(data, &bambooInfo)
	if err != nil {
		log.Fatalf("Couldn't unmarshal data due to: %v", err)
		return bambooInfo, err
	}

	return bambooInfo, nil
}
