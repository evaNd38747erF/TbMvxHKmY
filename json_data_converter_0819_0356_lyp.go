// 代码生成时间: 2025-08-19 03:56:07
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Converter is the struct that holds the database connection
type Converter struct {
    db *gorm.DB
}

// NewConverter creates a new instance of the Converter with a database connection
func NewConverter() (*Converter, error) {
    db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema
    db.AutoMigrate(&DataConversion{})

    return &Converter{db: db}, nil
}

// DataConversion is the model struct that represents the data conversion table
type DataConversion struct {
    ID        uint   `gorm:"primaryKey"`
    InputJSON string `gorm:"size:255"`
    OutputJSON string `gorm:"size:255"`
}

// ConvertJSON takes a JSON string as input and stores it in the database along with its converted output
func (c *Converter) ConvertJSON(inputJSON string) (string, error) {
    // Check if inputJSON is valid JSON
    var result map[string]interface{}
    if err := json.Unmarshal([]byte(inputJSON), &result); err != nil {
        return "", err
    }

    // Convert the JSON to a string
    outputJSON, err := json.Marshal(result)
    if err != nil {
        return "", err
    }

    // Save the input and output JSON to the database
    if err := c.db.Create(&DataConversion{InputJSON: inputJSON, OutputJSON: string(outputJSON)}).Error; err != nil {
        return "", err
    }

    return string(outputJSON), nil
}

func main() {
    converter, err := NewConverter()
    if err != nil {
        log.Fatal("Failed to create converter: ", err)
    }

    // Example usage
    inputJSON := `{"name": "John", "age": 30}`
    output, err := converter.ConvertJSON(inputJSON)
    if err != nil {
        log.Fatal("Failed to convert JSON: ", err)
    }

    fmt.Println("Converted JSON: ", output)
}