// 代码生成时间: 2025-08-07 07:52:36
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// CartItem represents an item in the shopping cart
type CartItem struct {
    gorm.Model
    ProductID uint
    Quantity  int
}

// Cart represents a shopping cart
type Cart struct {
    gorm.Model
    Items []CartItem `gorm:"foreignKey:CartID"`
}

// Product represents a product that can be added to the cart
type Product struct {
    gorm.Model
    Name  string
    Price float64
}

// Database connection settings
const dbDSN = "sqlite3:shopping.db"

func main() {
    db, err := gorm.Open(sqlite.Open(dbDSN), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }

    // Migrate the schema
    db.AutoMigrate(&Cart{}, &CartItem{}, &Product{})

    // Create a new product
    product := Product{Name: "Laptop", Price: 999.99}
    db.Create(&product)

    // Create a new cart
    cart := Cart{}
    db.Create(&cart)

    // Add an item to the cart
    addItemToCart(db, cart, product.ID, 1)

    // Retrieve cart items
    items, err := getCartItems(db, cart.ID)
    if err != nil {
        fmt.Println("Error retrieving cart items: ", err)
        return
    }
    for _, item := range items {
        fmt.Printf("Product ID: %d, Quantity: %d
", item.ProductID, item.Quantity)
    }
}

// addItemToCart adds an item to the cart
func addItemToCart(db *gorm.DB, cart Cart, productID uint, quantity int) {
    var item CartItem
    item.CartID = cart.ID
    item.ProductID = productID
    item.Quantity = quantity
    db.Create(&item)
}

// getCartItems retrieves the items in a cart
func getCartItems(db *gorm.DB, cartID uint) ([]CartItem, error) {
    var items []CartItem
    if err := db.Where(&CartItem{CartID: cartID}).Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}
