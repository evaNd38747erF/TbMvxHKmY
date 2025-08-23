// 代码生成时间: 2025-08-24 03:19:29
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "reflect"
)

// JSONDataConverter provides functionality to convert JSON data to and from structs.
type JSONDataConverter struct{}

// ConvertToStruct takes JSON data as a string and a pointer to struct to convert into.
// It returns an error if the conversion fails.
func (c *JSONDataConverter) ConvertToStruct(jsonStr string, dest interface{}) error {
    err := json.Unmarshal([]byte(jsonStr), dest)
    if err != nil {
        return fmt.Errorf("error unmarshalling JSON: %w", err)
    }
    return nil
}

// ConvertFromStruct takes a struct and converts it into a JSON string.
// It returns an error if the conversion fails.
func (c *JSONDataConverter) ConvertFromStruct(src interface{}) (string, error) {
    jsonData, err := json.MarshalIndent(src, "", "    ")
    if err != nil {
        return "", fmt.Errorf("error marshalling to JSON: %w", err)
    }
    return string(jsonData), nil
}

// Example usage of JSONDataConverter.
func main() {
    converter := JSONDataConverter{}

    // Example struct to hold JSON data.
    type ExampleStruct struct {
        Name    string `json:"name"`
        Age     int    `json:"age"`
        IsAdmin bool   `json:"isAdmin"`
    }

    // JSON string to convert to struct.
    jsonStrToStruct := `{
        "name": "John",
        "age": 30,
        "isAdmin": true
    }`

    // Struct to convert to JSON.
    exampleStruct := ExampleStruct{
        Name: "Jane",
        Age:  25,
    }

    // Convert JSON to struct.
    var exampleStructFromJSON ExampleStruct
    if err := converter.ConvertToStruct(jsonStrToStruct, &exampleStructFromJSON); err != nil {
        log.Fatalf("Failed to convert JSON to struct: %s", err)
    }
    fmt.Printf("Converted struct: %+v
", exampleStructFromJSON)

    // Convert struct to JSON.
    jsonStrFromStruct, err := converter.ConvertFromStruct(exampleStruct)
    if err != nil {
        log.Fatalf("Failed to convert struct to JSON: %s", err)
    }
    fmt.Printf("Converted JSON: %s
", jsonStrFromStruct)
}
