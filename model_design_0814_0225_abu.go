// 代码生成时间: 2025-08-14 02:25:59
 * and are used to interact with the database using the GORM framework.
 */

package main

import (
    "github.com/jinzhu/gorm"
    "log"
)

// User represents the User model with fields that map to the database columns.
type User struct {
    gorm.Model
    Name string
    Email string `gorm:"type:varchar(100);uniqueIndex"`
    Age  uint
}

// Product represents the Product model with fields that map to the database columns.
type Product struct {
    gorm.Model
    Code  string `gorm:"type:varchar(100);uniqueIndex"`
    Price uint
    // You can add more fields as needed.
}

// Error represents a custom error type for the application.
type Error struct {
    Message string
}

// NewError creates a new Error with the given message.
func NewError(message string) error {
    return &Error{Message: message}
}

// InitializeModels initializes the database models with GORM.
func InitializeModels(db *gorm.DB) error {
    // AutoMigrate will automatically create the tables if they don't exist.
    if err := db.AutoMigrate(&User{}, &Product{}).Error; err != nil {
        log.Printf("Failed to migrate database: %s", err)
        return NewError("Failed to migrate database")
    }
    return nil
}

// AddUser adds a new User to the database.
func AddUser(db *gorm.DB, user *User) error {
    if result := db.Create(user).Error; result != nil {
        log.Printf("Failed to add user: %s", result)
        return NewError("Failed to add user")
    }
    return nil
}

// AddProduct adds a new Product to the database.
func AddProduct(db *gorm.DB, product *Product) error {
    if result := db.Create(product).Error; result != nil {
        log.Printf("Failed to add product: %s", result)
        return NewError("Failed to add product")
    }
    return nil
}
