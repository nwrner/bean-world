package db

import (
	"database/sql"
	"fmt"
	"log"
)

func ReadCountries(db *sql.DB) {
	rows, err := db.Query("SELECT name FROM countries")
	if err != nil {
		log.Fatalf("Failed to query countries: %v", err)
	}
	defer rows.Close()

	fmt.Println("🌍 Countries in the Beanbase:")
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Printf("🇧🇪 %s\n", name)
	}
}
