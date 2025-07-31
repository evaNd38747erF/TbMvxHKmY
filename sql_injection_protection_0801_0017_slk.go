// 代码生成时间: 2025-08-01 00:17:53
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// 防止SQL注入示例程序
func main() {
    // 初始化数据库连接
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println("数据库连接失败：", err)
        return
    }
    defer db.Close()
    
    // 自动迁移模式，确保数据库表存在
    db.AutoMigrate(&User{})

    // 插入一条记录，演示防止SQL注入
    user := User{Name: "Alice", Age: 30}
    if err := db.Create(&user).Error; err != nil {
        fmt.Println("创建用户失败：", err)
        return
    }
    fmt.Println("用户创建成功")

    // 查询记录，演示防止SQL注入
    var result User
    if err := db.Where("name = ?", "Alice").First(&result).Error; err != nil {
        fmt.Println("查询用户失败：", err)
        return
    }
    fmt.Printf("查询到的用户：%+v
", result)
}

// User 定义用户模型
type User struct {
    gorm.Model
    Name string
    Age  uint
}
