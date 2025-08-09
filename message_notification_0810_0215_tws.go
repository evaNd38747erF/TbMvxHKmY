// 代码生成时间: 2025-08-10 02:15:17
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Message struct represents a notification message
type Message struct {
    gorm.Model
    Content string `gorm:"type:text"`
}

// DBClient represents a database client
type DBClient struct {
    *gorm.DB
}

// NewDBClient initializes a new database client
func NewDBClient() (*DBClient, error) {
    // Initialize SQLite database connection
    db, err := gorm.Open(sqlite.Open("notification.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema
    db.AutoMigrate(&Message{})

    return &DBClient{db}, nil
}

// SendMessage adds a new message to the database
func (c *DBClient) SendMessage(content string) error {
    message := Message{Content: content}
    if err := c.Create(&message).Error; err != nil {
        return err
    }
    return nil
}

// GetMessages retrieves all messages from the database
func (c *DBClient) GetMessages() ([]Message, error) {
    var messages []Message
    if err := c.Find(&messages).Error; err != nil {
        return nil, err
    }
    return messages, nil
}

func main() {
    // Initialize the database client
    dbClient, err := NewDBClient()
    if err != nil {
        fmt.Printf("Failed to connect to database: %v
", err)
        return
    }
    defer dbClient.DB.Close()

    // Send a new message
    if err := dbClient.SendMessage("Hello, this is a test message!"); err != nil {
        fmt.Printf("Failed to send message: %v
", err)
        return
    }
    fmt.Println("Message sent successfully!")

    // Retrieve messages
    messages, err := dbClient.GetMessages()
    if err != nil {
        fmt.Printf("Failed to retrieve messages: %v
", err)
        return
    }
    fmt.Println("Retrieved messages:")
    for _, message := range messages {
        fmt.Printf("ID: %d, Content: %s
", message.ID, message.Content)
    }
}
