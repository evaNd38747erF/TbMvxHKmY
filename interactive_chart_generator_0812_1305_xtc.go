// 代码生成时间: 2025-08-12 13:05:07
 * interactive_chart_generator.go
 * This program is an interactive chart generator using GORM for database operations.
 */

package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// ChartData represents the data structure for chart data.
type ChartData struct {
    ID    uint   `gorm:""`
    Label string `gorm:"type:varchar(255);"`
    Value int    `gorm:"type:int"`
}

// Chart represents the data structure for a chart.
type Chart struct {
    gorm.Model
    Name  string `gorm:"type:varchar(255);"`
    // One-to-many relationship with ChartData
    Data []ChartData `gorm:"foreignKey:ChartID"`
}

func main() {
    // Initialize the database connection.
    db, err := gorm.Open(sqlite.Open("chart.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    // Migrate the schema.
    db.AutoMigrate(&Chart{}, &ChartData{})

    // Simulate interactive chart generation.
    // In a real-world scenario, this would be replaced with user input.
    fmt.Println("Enter chart name: ")
    var chartName string
    fmt.Scanln(&chartName)
    fmt.Println("Enter chart data (label and value, separated by a space): ")
    var labels []string
    var values []int
    for {
        var label string
        var value int
        fmt.Scanln(&label, &value)
        labels = append(labels, label)
        values = append(values, value)
        fmt.Println("Enter another data point (label value) or type 'done' to finish: ")
        var input string
        fmt.Scanln(&input)
        if input == "done" {
            break
        }
    }

    // Create a new chart.
    chart := Chart{Name: chartName}
    if result := db.Create(&chart); result.Error != nil {
        panic(result.Error)
    }

    // Create chart data and associate with the chart.
    for i := range labels {
        chartData := ChartData{Label: labels[i], Value: values[i], ChartID: chart.ID}
        if result := db.Create(&chartData); result.Error != nil {
            panic(result.Error)
        }
    }

    fmt.Printf("Chart '%s' created successfully with the following data points:
", chartName)
    for _, data := range chart.Data {
        fmt.Printf("Label: %s, Value: %d
", data.Label, data.Value)
    }
}
