// 代码生成时间: 2025-07-31 07:50:38
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User holds the user data
type User struct {
    gorm.Model
    Name string
    Email string `gorm:"uniqueIndex"`
}

// DB is our struct to hold our db connection
type DB struct {
    *gorm.DB
}

// NewDB returns a new instance of DB
func NewDB() (*DB, error) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema
    db.AutoMigrate(&User{})
    return &DB{db}, nil
}

func (db *DB) CreateUser(name, email string) error {
    // Using GORM's built-in methods to prevent SQL injection
    user := User{Name: name, Email: email}
    if result := db.Create(&user); result.Error != nil {
        return result.Error
    }
    return nil
}

func (db *DB) GetUserByEmail(email string) (*User, error) {
    // Using GORM's built-in methods to prevent SQL injection
    var user User
    if result := db.Where("email = ?", email).First(&user); result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func main() {
    db, err := NewDB()
    if err != nil {
        fmt.Println("Failed to connect to database: &", err)
        return
    }
    defer db.Close()

    // Example of creating a user
    if err := db.CreateUser("John Doe", "john.doe@example.com"); err != nil {
        fmt.Println("Failed to create user: &", err)
    } else {
        fmt.Println("User created successfully")
    }

    // Example of retrieving a user by email
    user, err := db.GetUserByEmail("john.doe@example.com")
    if err != nil {
        fmt.Println("Failed to retrieve user: &", err)
    } else {
        fmt.Printf("Retrieved user: %+v
", user)
    }
}
