// 代码生成时间: 2025-08-25 23:32:34
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Order represents the model for an order
type Order struct {
    gorm.Model
    OrderID   uint   `gorm:"primaryKey"`
    Username string
    Amount   float64
    Status   string // Possible values: 'pending', 'completed', 'cancelled'
}

// DB is the global variable for our database connection
var DB *gorm.DB

// SetupDB sets up the database connection
func SetupDB() error {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return err
    }

    // Migrate the schema
    if err = DB.AutoMigrate(&Order{}); err != nil {
        return err
    }

    return nil
}

// CreateOrder creates a new order
func CreateOrder(username string, amount float64) (*Order, error) {
    newOrder := Order{Username: username, Amount: amount, Status: "pending"}
    result := DB.Create(&newOrder)
    if result.Error != nil {
        return nil, result.Error
    }
    return &newOrder, nil
}

// CompleteOrder marks an order as completed
func CompleteOrder(orderID uint) error {
    var order Order
    if result := DB.First(&order, orderID); result.Error != nil {
        return result.Error
    }
    order.Status = "completed"
    result := DB.Save(&order)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// CancelOrder marks an order as cancelled
func CancelOrder(orderID uint) error {
    var order Order
    if result := DB.First(&order, orderID); result.Error != nil {
        return result.Error
    }
    order.Status = "cancelled"
    result := DB.Save(&order)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func main() {
    err := SetupDB()
    if err != nil {
        fmt.Println("Failed to setup database:", err)
        return
    }

    newOrder, err := CreateOrder("john_doe", 99.99)
    if err != nil {
        fmt.Println("Failed to create order:", err)
        return
    }
    fmt.Printf("Created order %+v
", newOrder)

    err = CompleteOrder(newOrder.OrderID)
    if err != nil {
        fmt.Println("Failed to complete order:", err)
        return
    }
    fmt.Println("Order completed successfully")

    // Uncomment to test cancellation
    // err = CancelOrder(newOrder.OrderID)
    // if err != nil {
    //     fmt.Println("Failed to cancel order:", err)
    //     return
    // }
    // fmt.Println("Order cancelled successfully")
}
