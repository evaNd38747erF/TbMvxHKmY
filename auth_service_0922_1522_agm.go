// 代码生成时间: 2025-09-22 15:22:09
package main

import (
    "fmt"
    "log"
    "net/http"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "encoding/json"
)

// User represents a user entity with fields for authentication
type User struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"password"`
}

// AuthService provides methods for user authentication
type AuthService struct {
    DB *gorm.DB
}

// NewAuthService creates a new AuthService instance
func NewAuthService() *AuthService {
    db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Migrate the schema
    db.AutoMigrate(&User{})
    return &AuthService{DB: db}
}

// Authenticate checks if a user exists with the given username and password
func (as *AuthService) Authenticate(username, password string) (*User, error) {
    var user User
    // Find the user by username
    if err := as.DB.Where(&User{Username: username}).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("database error: %w", err)
    }
    // Check the password (for simplicity, passwords are not hashed in this example)
    if user.Password != password {
        return nil, fmt.Errorf("invalid password")
    }
    return &user, nil
}

// AuthHandler handles HTTP requests for user authentication
func AuthHandler(w http.ResponseWriter, r *http.Request) {
    var authData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&authData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    authService := NewAuthService()
    user, err := authService.Authenticate(authData.Username, authData.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    // If authentication is successful, return the user data
    response, err := json.Marshal(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%s", string(response))
}

func main() {
    http.HandleFunc("/authenticate", AuthHandler)
    log.Println("Server is running on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Server failed to start: ", err)
    }
}