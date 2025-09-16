// 代码生成时间: 2025-09-17 00:08:51
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User represents a user with access control
type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
    Role     string
}

// Database connection
var db *gorm.DB
var err error

func initDB() {
    // Connect to SQLite database
    db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&User{})
}

// CheckAccess checks if a user has access based on their role
func CheckAccess(username, password, role string) error {
    var user User
    // Find user by username and password
    if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
        return fmt.Errorf("authentication failed: %w", err)
    }

    // Check if user role matches the required role
    if user.Role != role {
        return fmt.Errorf("access denied: user role does not match required role")
    }

    return nil
}

func main() {
    initDB()
    defer db.Close()

    // Example usage
    username := "admin"
    password := "password123"
    requiredRole := "admin"

    if err := CheckAccess(username, password, requiredRole); err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Println("Access granted")
    }
}
