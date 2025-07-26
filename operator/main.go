package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/nwrner/bean-world/db"
)

func main() {
	dbHost := "postgres"
	dbPort := "5432"
	dbUser := "postgres"
	dbPass := "supersecret"
	dbName := "beandata"

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	fmt.Println("🎩 Operator online!")
	db.ReadCountries(dbConn) // or whatever function you wrote
	log.Println("🫘 BLO is online and chilling. Waiting for world domination plans...")
	for {
		time.Sleep(time.Hour)
	}
}
