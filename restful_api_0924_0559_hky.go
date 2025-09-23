// 代码生成时间: 2025-09-24 05:59:55
package main

import (
    "encoding/json"
    "net/http"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Item represents the data model for an item in the database.
type Item struct {
    gorm.Model
    Name  string
    Price uint
}

// SetupDatabase initializes the database connection.
func SetupDatabase() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&Item{})
    return db
}

// ItemHandler handles the API requests for the 'Item' resource.
func ItemHandler(w http.ResponseWriter, r *http.Request) {
    db := SetupDatabase()
    var item Item
    switch r.Method {
    case http.MethodGet:
        // Fetch all items
        items := []Item{}
        if err := db.Find(&items).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(items)
    case http.MethodPost:
        // Create a new item
        if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := db.Create(&item).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json\)
        json.NewEncoder(w).Encode(item)
    default:
        http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
    }
}

// main function to start the HTTP server.
func main() {
    http.HandleFunc("/items", ItemHandler)
    http.ListenAndServe(":8080", nil)
}