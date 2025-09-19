// 代码生成时间: 2025-09-20 06:25:24
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User 定义用户数据模型
type User struct {
    gorm.Model
    Name     string
    Email    string `gorm:"type:varchar(100);uniqueIndex"`
    Age     int
}

// 初始化数据库连接
func initDB() *gorm.DB {
    conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    return conn
}

// Migrate 迁移数据库，创建数据模型
func migrate(db *gorm.DB) {
    // 自动迁移模型到数据库，即自动创建数据库表
    db.AutoMigrate(&User{})
}

func main() {
    // 初始化数据库连接
    db := initDB()
    defer db.Close()
    
    // 迁移数据库，创建数据模型
    migrate(db)
    
    // 创建新用户
    user := User{Name: "John Doe", Email: "johndoe@example.com", Age: 30}
    // 错误处理
    if err := db.Create(&user).Error; err != nil {
        fmt.Println("Error creating user: \{err}")
        return
    }
    
    // 查询用户
    var userResult User
    if err := db.First(&userResult, 1).Error; err != nil {
        fmt.Println("Error finding user: \{err}")
        return
    }
    
    // 打印用户信息
    fmt.Printf("Found user: %+v\
", userResult)
}
