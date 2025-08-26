// 代码生成时间: 2025-08-26 12:05:43
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// User represents a user in our system
type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
}

// LoginRequest is the structure for incoming login requests
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// LoginResponse is the structure for the response after a login attempt
type LoginResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
}

// LoginService is a struct that holds the database connection
type LoginService struct {
    db *gorm.DB
}

// NewLoginService initializes a new login service with a database connection
func NewLoginService(db *gorm.DB) *LoginService {
    return &LoginService{db: db}
}

// Login attempts to authenticate a user and returns a response
func (s *LoginService) Login(req LoginRequest) (*LoginResponse, error) {
    var user User
    // Attempt to find the user by username
    if err := s.db.Where(&User{Username: req.Username}).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("error finding user: %w", err)
    }

    // Check if the user's password matches the one provided
    if user.Password != req.Password {
        return nil, fmt.Errorf("invalid password")
    }

    // Return a successful login response
    return &LoginResponse{Success: true, Message: "Login successful"}, nil
}

func main() {
    // Initialize the database connection
    db, err := gorm.Open(sqlite.Open("login.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database: ", err)
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&User{})

    // Create a new login service
    loginService := NewLoginService(db)

    // Sample login request
    req := LoginRequest{Username: "admin", Password: "admin123"}

    // Perform the login and handle the response
    res, err := loginService.Login(req)
    if err != nil {
        fmt.Println("Login failed: ", err)
        return
    }

    // Print the login response
    fmt.Printf("Login response: %+v
", res)
}