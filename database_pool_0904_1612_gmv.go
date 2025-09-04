// 代码生成时间: 2025-09-04 16:12:30
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// DatabaseConfig 存储数据库配置信息
type DatabaseConfig struct {
    DSN string // 数据源名称
}

// Database 封装了gorm.DB实例
type Database struct {
    db *gorm.DB
}

// NewDatabase 创建一个新的Database实例
func NewDatabase(config *DatabaseConfig) (*Database, error) {
    // 使用sqlite驱动，连接到数据库
    db, err := gorm.Open(sqlite.Open(config.DSN), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // 自动迁移模式
    db.AutoMigrate(&User{}) // 假设有一个User模型，这里用于数据库迁移

    return &Database{db}, nil
}

// Close 关闭数据库连接
func (d *Database) Close() error {
    // 使用sql.DB的Close方法关闭连接
    return d.db.Close()
}

// User 定义一个简单的用户模型
type User struct {
    gorm.Model
    Name string
}

func main() {
    // 数据库配置
    config := &DatabaseConfig{
        DSN: "test.db",
    }

    // 创建数据库连接
    db, err := NewDatabase(config)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // 插入一条用户数据
    newUser := User{Name: "John Doe"}
    if err := db.db.Create(&newUser).Error; err != nil {
        log.Fatalf("Failed to create user: %v", err)
    }

    // 查询所有用户数据
    var users []User
    if err := db.db.Find(&users).Error; err != nil {
        log.Fatalf("Failed to find users: %v", err)
    }

    fmt.Println("Users:")
    for _, user := range users {
        fmt.Printf("ID: %d, Name: %s
", user.ID, user.Name)
    }
}
