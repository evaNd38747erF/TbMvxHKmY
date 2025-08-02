// 代码生成时间: 2025-08-02 21:51:44
// json_converter.go

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Define a struct to represent our JSON data
type JSONData struct {
    Name    string `json:"name"`
    Age    int    `json:"age"`
    Address string `json:"address"`
}

// Define a struct to represent our database model
type Person struct {
    gorm.Model
    Name    string
    Age     int
    Address string
}

func main() {
    // Connect to the SQLite database
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&Person{})

    // Example JSON data
    jsonData := `{"name": "John Doe", "age": 30, "address": "123 Main St"}`

    // Convert JSON to Go struct
    var data JSONData
    err = json.Unmarshal([]byte(jsonData), &data)
    if err != nil {
        log.Fatalf("failed to unmarshal JSON: %v", err)
    }

    // Create a new person record
    person := Person{Name: data.Name, Age: data.Age, Address: data.Address}

    // Save the person record to the database
    result := db.Create(&person)
    if result.Error != nil {
        log.Fatalf("failed to create person record: %v", result.Error)
    }

    // Read the person record from the database
    var retrievedPerson Person
    result = db.First(&retrievedPerson, person.ID)
    if result.Error != nil {
        log.Fatalf("failed to retrieve person record: %v", result.Error)
    }

    // Convert Go struct back to JSON
    personJSON, err := json.MarshalIndent(&retrievedPerson, "", "  ")
    if err != nil {
        log.Fatalf("failed to marshal to JSON: %v", err)
    }

    // Print the JSON data
    fmt.Println(string(personJSON))
}
