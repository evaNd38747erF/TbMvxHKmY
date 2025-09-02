// 代码生成时间: 2025-09-02 23:53:33
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "testing"
)

// User 定义一个用户模型
type User struct {
    gorm.Model
    Name string
    Age  uint
}

// SetupTestDB 初始化测试数据库
func SetupTestDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    
    // 自动迁移模式
    db.AutoMigrate(&User{})
    return db
}

// TestUserCreate 测试创建用户
func TestUserCreate(t *testing.T) {
    db := SetupTestDB()
    defer db.Migrator().DropTable(&User{})
    
    user := User{Name: "John Doe", Age: 30}
    
    // 创建用户
    if err := db.Create(&user).Error; err != nil {
        t.Errorf("failed to create user: %v", err)
        return
    }
    
    // 验证用户是否被正确创建
    var result User
    if err := db.First(&result, user.ID).Error; err != nil {
        t.Errorf("failed to find user: %v", err)
        return
    }
    
    if result.Name != user.Name || result.Age != user.Age {
        t.Errorf("User name or age is incorrect, got: %v years old %v, want: %v years old %v", result.Name, result.Age, user.Name, user.Age)
    }
}

func main() {
    // 运行测试
    testing.Main()
}