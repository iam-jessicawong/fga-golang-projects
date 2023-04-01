package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "gorm_challenge"
	db       *gorm.DB
	err      error
)

func StartDB() {
	log.Println("Connectiong to database")

	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	if db, err = gorm.Open(postgres.Open(config), &gorm.Config{}); err != nil {
		log.Fatal("error connecting to database:", err)
	}

	log.Println("Successfully connected to database")
}

func GetDB() *gorm.DB {
	return db
}
