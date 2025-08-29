// 代码生成时间: 2025-08-29 22:52:35
// restful_api_with_gorm.go

package main

import (
    "database/sql"
    "fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Product 定义模型
type Product struct {
    gorm.Model
    Code  string
    Price uint
}

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
    router := gin.Default()

    // 产品路由组
    productGroup := router.Group("/products")
    {
        productGroup.POST("", createProduct) // 创建产品
        productGroup.GET("", listProducts)   // 列出所有产品
        productGroup.GET="/:id", getProductByID) // 根据ID获取产品
        productGroup.PUT="/:id", updateProduct) // 更新产品
        productGroup.DELETE="/:id", deleteProduct) // 删除产品
    }

    return router
}

// createProduct 创建一个新的产品
func createProduct(c *gin.Context) {
    var newProduct Product
    if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    db, err := createConnection()
    if err != nil {
        c.JSON(500, gin.H{"error": "database connection failed"})
        return
    }
    result := db.Create(&newProduct)
    if result.Error != nil {
        c.JSON(500, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(201, newProduct)
}

// listProducts 列出所有产品
func listProducts(c *gin.Context) {
    db, err := createConnection()
    if err != nil {
        c.JSON(500, gin.H{"error": "database connection failed"})
        return
    }
    var products []Product
    if result := db.Find(&products); result.Error != nil {
        c.JSON(500, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(200, products)
}

// getProductByID 根据ID获取产品
func getProductByID(c *gin.Context) {
    var product Product
    id := c.Param("id")
    db, err := createConnection()
    if err != nil {
        c.JSON(500, gin.H{"error": "database connection failed"})
        return
    }
    if result := db.First(&product, id).Error; result != nil {
        c.JSON(500, gin.H{"error": result.Error()})
        return
    }
    c.JSON(200, product)
}

// updateProduct 更新一个产品
func updateProduct(c *gin.Context) {
    var productToUpdate Product
    id := c.Param("id")
    if err := c.ShouldBindJSON(&productToUpdate); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    db, err := createConnection()
    if err != nil {
        c.JSON(500, gin.H{"error": "database connection failed"})
        return
    }
    if result := db.Model(&Product{}).Where("id = ?", id).Updates(productToUpdate).Error; result != nil {
        c.JSON(500, gin.H{"error": result.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "product updated successfully"})
}

// deleteProduct 删除一个产品
func deleteProduct(c *gin.Context) {
    id := c.Param("id")
    db, err := createConnection()
    if err != nil {
        c.JSON(500, gin.H{"error": "database connection failed"})
        return
    }
    if result := db.Delete(&Product{}, id).Error; result != nil {
        c.JSON(500, gin.H{"error": result.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "product deleted successfully"})
}

// createConnection 创建数据库连接
func createConnection() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    // 迁移模式，确保数据库结构是最新的
    db.AutoMigrate(&Product{})
    return db, nil
}

func main() {
    router := SetupRouter()
    router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
