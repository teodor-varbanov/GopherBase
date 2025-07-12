package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dbConnect(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Can't connect to DB due to error: %v", err)
		return nil, err
	}
	return db, nil
}
