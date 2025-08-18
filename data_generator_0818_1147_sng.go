// 代码生成时间: 2025-08-18 11:47:42
// data_generator.go
// This file contains the implementation of a data generator using GORM in Go.

package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Define the model
type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"type:varchar(100);uniqueIndex"`
}

// Database connection configuration
const dbConfig = "data.db"

func main() {
    // Initialize the database connection
    db, err := gorm.Open(sqlite.Open(dbConfig), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }

    // Migrate the schema
    db.AutoMigrate(&User{})

    // Generate test data
    err = generateTestData(db, 100)
    if err != nil {
        fmt.Println("Error generating test data: ", err)
    } else {
        fmt.Println("Test data generated successfully")
    }
}

// Function to generate the specified number of users
func generateTestData(db *gorm.DB, count int) error {
    for i := 0; i < count; i++ {
        user := User{
            Name:  fmt.Sprintf("User%d", i),
            Email: fmt.Sprintf("user%d@example.com", i),
        }
        // Create a new user
        if err := db.Create(&user).Error; err != nil {
            return err
        }
    }
    return nil
}
