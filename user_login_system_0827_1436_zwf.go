// 代码生成时间: 2025-08-27 14:36:29
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User represents a user with username, password, and other fields
type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
}

// LoginRequest represents the login request payload
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// UserLoginService is a service that handles user login
type UserLoginService struct {
    DB *gorm.DB
}

// NewUserLoginService creates a new instance of UserLoginService
func NewUserLoginService(db *gorm.DB) *UserLoginService {
    return &UserLoginService{DB: db}
}

// Login performs user authentication and returns a success message or an error
func (service *UserLoginService) Login(request LoginRequest) (string, error) {
    var user User
    // Attempt to find the user by username
    if result := service.DB.Where("username = ?", request.Username).First(&user); result.Error != nil {
        // Handle different error cases
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return "", fmt.Errorf("user not found")
        }
        return "", result.Error
    }
    // Check if the password is correct
    if user.Password != request.Password {
        return "", fmt.Errorf("invalid password")
    }
    return "Login successful", nil
}

func main() {
    // Initialize the database connection
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: ", err)
    }

    // Migrate the schema
    db.AutoMigrate(&User{})

    // Create a new user login service
    service := NewUserLoginService(db)

    // Set up HTTP server
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
            return
        }

        var request LoginRequest
        if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        defer r.Body.Close()

        response, err := service.Login(request)
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprintln(w, err.Error())
            return
        }
        fmt.Fprintln(w, response)
    })

    fmt.Println("Server is running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
