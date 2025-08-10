// 代码生成时间: 2025-08-11 02:03:47
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// User represents the user model
type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
}

// AuthService handles user authentication
type AuthService struct {
    db *gorm.DB
}

// NewAuthService initializes a new AuthService with a database connection
func NewAuthService(db *gorm.DB) *AuthService {
    return &AuthService{db: db}
}

// HashPassword hashes a plain password for storage in the database
func HashPassword(password string) string {
    // Hash the password using SHA-256
    hashed := sha256.Sum256([]byte(password))
    return hex.EncodeToString(hashed[:])
}

// VerifyPassword checks if a given password matches the stored hash
func VerifyPassword(hashedPassword, password string) bool {
    return hashedPassword == HashPassword(password)
}

// RegisterUser registers a new user with a given username and password
func (a *AuthService) RegisterUser(username, password string) error {
    // Hash the password before saving it to the database
    hashedPassword := HashPassword(password)
    
    user := User{Username: username, Password: hashedPassword}
    
    // Attempt to save the new user
    result := a.db.Create(&user)
    if result.Error != nil {
        // Handle any errors that occur during the save operation
        return result.Error
    }
    return nil
}

// AuthenticateUser authenticates a user with a given username and password
func (a *AuthService) AuthenticateUser(username, password string) (bool, error) {
    var user User
    // Find the user by username
    result := a.db.Where(&User{Username: username}).First(&user)
    if result.Error != nil {
        // If the user is not found, return false and no error
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return false, nil
        }
        // Handle other database errors
        return false, result.Error
    }
    // Verify the password
    if !VerifyPassword(user.Password, password) {
        return false, nil
    }
    return true, nil
}

func main() {
    // Initialize the database connection
    db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    defer db.Close()
    
    // Migrate the schema
    db.AutoMigrate(&User{})
    
    authService := NewAuthService(db)
    
    // Register a new user
    err = authService.RegisterUser("testuser", "testpassword")
    if err != nil {
        log.Println("Registration failed: ", err)
    } else {
        log.Println("User registered successfully")
    }
    
    // Authenticate the user
    authenticated, err := authService.AuthenticateUser("testuser", "testpassword")
    if err != nil {
        log.Println("Authentication failed: ", err)
    } else if authenticated {
        log.Println("User authenticated successfully")
    } else {
        log.Println("Authentication failed: Incorrect credentials")
    }
}