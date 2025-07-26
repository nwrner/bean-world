package db

import (
	"database/sql"
	"fmt"
	"log"
)

func InitGoodsTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS goods (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create goods table: %v", err)
	}
}

func SeedGoods(db *sql.DB) {
	goods := []string{
		"Wheat", "Electronics", "Steel", "Cocoa", "Spices",
		"Textiles", "Lumber", "Plastics", "Medicines", "Tea",
		"Beans", "Coffee", "Oil", "Gold", "Software",
	}

	for _, name := range goods {
		_, err := db.Exec(`INSERT INTO goods (name) VALUES ($1) ON CONFLICT (name) DO NOTHING`, name)
		if err != nil {
			log.Printf("Error inserting good '%s': %v", name, err)
		}
		fmt.Println("🎉 Added good:", name)
	}
	fmt.Println("📦 Goods seeded!")
}
