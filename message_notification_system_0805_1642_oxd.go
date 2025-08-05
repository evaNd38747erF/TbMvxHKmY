// 代码生成时间: 2025-08-05 16:42:29
package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Notification defines the structure for a notification message.
type Notification struct {
    gorm.Model
    Title       string  `gorm:"column:title;type:varchar(255);"`
    Content     string  `gorm:"column:content;type:text;"`
    ReceiverID uint    `gorm:"column:receiver_id;type:int;"`
    Status      string  `gorm:"column:status;type:varchar(50);"`
}

// User defines the structure for a user.
type User struct {
    gorm.Model
    Name       string  `gorm:"column:name;type:varchar(255);"`
    Email      string  `gorm:"column:email;type:varchar(255);"`
    Notifications []Notification `gorm:"foreignKey:ReceiverID;"`
}

func main() {
    // Initialize a new SQLite database.
    db, err := gorm.Open(sqlite.Open("notification.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    defer db.Close()
    
    // Migrate the schema.
    db.AutoMigrate(&Notification{}, &User{})

    // Create a new user.
    newUser := User{Name: "John Doe", Email: "john.doe@example.com"}
    if result := db.Create(&newUser); result.Error != nil {
        log.Fatal("Failed to create new user: ", result.Error)
    }

    // Create a new notification for the user.
    newNotification := Notification{
        Title:       "Welcome",
        Content:     "Thank you for registering with us!",
        ReceiverID: newUser.ID,
        Status:      "unread",
    }
    if result := db.Create(&newNotification); result.Error != nil {
        log.Fatal("Failed to create new notification: ", result.Error)
    }

    // Fetch notifications for the user.
    var notifications []Notification
    if result := db.Where("receiver_id = ? AND status = 'unread'", newUser.ID).Find(&notifications).Error; result != nil {
        log.Fatal("Failed to fetch notifications: ", result.Error)
    }
    fmt.Println("Notifications for user:", newUser.Name)
    for _, notification := range notifications {
        fmt.Printf("Title: %s, Content: %s
", notification.Title, notification.Content)
    }
}
