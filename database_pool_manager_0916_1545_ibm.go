// 代码生成时间: 2025-09-16 15:45:22
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// DBConfig contains the configuration for the database connection
type DBConfig struct {
    Dsn string
}

// Database is a wrapper around GORM DB connection
type Database struct {
    *gorm.DB
}

// NewDatabase creates a new Database instance with the provided configuration
func NewDatabase(config *DBConfig) (*Database, error) {
    db, err := gorm.Open(sqlite.Open(config.Dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    // Migrate the schema
    err = db.AutoMigrate()
    if err != nil {
        return nil, fmt.Errorf("failed to migrate database: %w", err)
    }

    return &Database{db}, nil
}

// Close closes the database connection
func (db *Database) Close() error {
    if db.DB != nil {
        return db.DB.Close()
    }
    return nil
}

func main() {
    // Database configuration
    config := &DBConfig{Dsn: "sqlite.db"}

    // Create a new database connection
    db, err := NewDatabase(config)
    if err != nil {
        fmt.Printf("Error creating database connection: %s
", err)
        return
    }
    defer db.Close()

    // Use the database connection as needed...
    // For example, you can use db.DB to perform queries
}
