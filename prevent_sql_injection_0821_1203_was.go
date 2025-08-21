// 代码生成时间: 2025-08-21 12:03:53
// 防止SQL注入的示例程序
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User 代表用户模型
type User struct {
    gorm.Model
    Name string
}

// DB 是 *gorm.DB 的别名，用于数据库操作
var DB *gorm.DB

func initDB() error {
    // 连接数据库
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return err
    }

    // 自动迁移模式
    DB.AutoMigrate(&User{})
    return nil
}

func main() {
    if err := initDB(); err != nil {
        fmt.Printf("初始化数据库失败: %v
", err)
        return
    }

    // 使用事务防止SQL注入
    if err := DB.Transaction(func(tx *gorm.DB) error {
        // 假设我们需要查询用户，防止SQL注入的方式是使用参数化查询
        // 而不是直接拼接SQL语句
        user := User{Name: "John Doe"}
        if err := tx.Create(&user).Error; err != nil {
            return err
        }

        // 使用Find方法查询用户，而不是直接传入SQL语句
        // 这样可以避免SQL注入的风险
        var foundUser User
        if err := tx.Where(&User{Name: "John Doe"}).First(&foundUser).Error; err != nil {
            return err
        }

        fmt.Printf("Found user: %+v
", foundUser)
        return nil
    }); err != nil {
        fmt.Printf("事务处理失败: %v
", err)
    }
}