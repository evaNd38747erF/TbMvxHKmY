// 代码生成时间: 2025-08-04 18:31:33
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Book defines a model for books
type Book struct {
    gorm.Model
    Title  string
    Author string
}

// SetupDB initializes the database connection
func SetupDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }
    
    // Migrate the schema
    db.AutoMigrate(&Book{})
    return db
}

// CreateBook handles HTTP POST requests for creating a new book
func CreateBook(c *gin.Context) {
    var book Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    db := SetupDB()
    if err := db.Create(&book).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
        return
    }
    
    c.JSON(http.StatusOK, book)
}

// GetBooks handles HTTP GET requests for retrieving all books
func GetBooks(c *gin.Context) {
    db := SetupDB()
    var books []Book
    if err := db.Find(&books).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
        return
    }
    
    c.JSON(http.StatusOK, books)
}

// main function to start the server
func main() {
    router := gin.Default()
    
    router.POST("/books", CreateBook)
    router.GET("/books", GetBooks)
    
    router.Run(":8080") // listen and serve on 0.0.0.0:8080
}