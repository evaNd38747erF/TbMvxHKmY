// 代码生成时间: 2025-09-20 22:10:26
package main

import (
    "fmt"
    "log"
    "net/http"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Chart represents the data structure for a chart
type Chart struct {
    gorm.Model
    Title string `gorm:"type:varchar(100);uniqueIndex"`
    Data  string `gorm:"type:text"` // JSON string of chart data
}

// SetupDB initializes the database connection
func SetupDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("chart.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    // Migrate the schema
    db.AutoMigrate(&Chart{})
    return db
}

// ChartHandler handles HTTP requests for creating and retrieving charts
func ChartHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    switch r.Method {
    case "GET":
        // Retrieve all charts
        var charts []Chart
        if err := db.Find(&charts).Error; err != nil {
            http.Error(w, "Error retrieving charts: "+err.Error(), http.StatusInternalServerError)
            return
        }
        // Assuming we encode the charts to JSON and send back to the client
        // For simplicity, not implemented here
        fmt.Fprintf(w, "{"charts": %+v}", charts)
    case "POST":
        // Create a new chart
        var chart Chart
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
            return
        }
        chart.Title = r.FormValue("title")
        chart.Data = r.FormValue("data")
        if err := db.Create(&chart).Error; err != nil {
            http.Error(w, "Error creating chart: "+err.Error(), http.StatusInternalServerError)
            return
        }
        // Send back the created chart as JSON
        fmt.Fprintf(w, "{"chart": %+v}", chart)
    default:
        http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
    }
}

func main() {
    db := SetupDB()
    defer db.Close()

    http.HandleFunc("/chart", func(w http.ResponseWriter, r *http.Request) {
        ChartHandler(w, r, db)
    })

    fmt.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
