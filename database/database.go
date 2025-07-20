package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func NewDatabase() (*gorm.DB, error) {
	dbPath := "./mydb.sqlite"
	_, err := os.Create(dbPath)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Test basic query
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("failed to get DB handle: %w", err))
	}

	// Ping the DB (succeeds for SQLite unless there's a file/lock issue)
	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Errorf("ping failed: %w", err))
	}

	fmt.Println("✅ Connected to SQLite database successfully.")
	return db, err
}
