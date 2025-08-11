// 代码生成时间: 2025-08-12 05:44:34
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

// JSONDataConverter is the structure that holds the data to be converted.
type JSONDataConverter struct {
    // Data holds the input JSON data as a byte slice.
    Data []byte
}

// NewJSONDataConverter creates a new instance of JSONDataConverter with the provided data.
func NewJSONDataConverter(data []byte) *JSONDataConverter {
    return &JSONDataConverter{
        Data: data,
    }
}

// ConvertJSON converts the JSON data to a pretty-printed JSON string.
func (converter *JSONDataConverter) ConvertJSON() (string, error) {
    // Create a temporary variable to hold the pretty-printed JSON data.
    var prettyJSON []byte
    var err error

    // Attempt to marshal the data into a pretty-printed JSON format.
    prettyJSON, err = json.MarshalIndent(converter.Data, "", "    ")
    if err != nil {
        return "", err
    }

    // Convert the byte slice to a string and return it.
    return string(prettyJSON), nil
}

func main() {
    // Example JSON data to be converted.
    jsonData := []byte(`{"name": "John", "age": 30, "is_employee": true}`)

    // Create a new JSONDataConverter instance.
    converter := NewJSONDataConverter(jsonData)

    // Convert the JSON data to a pretty-printed format.
    prettyJSON, err := converter.ConvertJSON()
    if err != nil {
        log.Fatalf("Error converting JSON: %v", err)
    }

    // Output the pretty-printed JSON to the console.
    fmt.Println("Pretty-printed JSON: 
", prettyJSON)
}
