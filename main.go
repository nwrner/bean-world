package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nwrner/bean-world/db"

	_ "github.com/lib/pq"
)

// WipeAllTables drops all relevant tables (dev-only!)
func WipeAllTables(db *sql.DB) {
	tables := []string{"countries", "goods"}

	for _, table := range tables {
		query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", table)
		_, err := db.Exec(query)
		if err != nil {
			fmt.Printf("❌ Failed to drop table %s: %v\n", table, err)
		} else {
			fmt.Printf("🔥 Dropped table %s\n", table)
		}
	}
}

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

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("Ping failed: %v", err)
	}

	fmt.Println("🎉 Connected to the Beanbase!")

	WipeAllTables(dbConn)

	db.InitCountriesTable(dbConn)
	db.SeedCountries(dbConn)

	db.InitGoodsTable(dbConn)
	db.SeedGoods(dbConn)
}
