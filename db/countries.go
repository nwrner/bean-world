package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func InitCountriesTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS countries (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create countries table: %v", err)
	}
}

func SeedCountries(db *sql.DB) {
	rand.Seed(time.Now().UnixNano())
	prefixes := []string{"Black", "Kidney", "Bean", "Lentil", "Soy", "Refri", "Chick", "Mungo", "Pinto", "Cannar"}
	suffixes := []string{"mark", "land", "topia", "shire", "zuela", "rich", "stan"}

	for i := 0; i < 20; i++ {
		name := prefixes[rand.Intn(len(prefixes))] + suffixes[rand.Intn(len(suffixes))]
		_, err := db.Exec(`INSERT INTO countries (name) VALUES ($1) ON CONFLICT (name) DO NOTHING`, name)
		if err != nil {
			log.Printf("Error inserting country '%s': %v", name, err)
		}
		fmt.Println("🎉 Added country:", name)
	}
	fmt.Println("🌍 Countries seeded!")
}
