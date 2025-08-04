// 代码生成时间: 2025-08-04 08:05:01
package main

import (
    "fmt"
    "testing"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User 模型定义
type User struct {
    gorm.Model
    Name string
    Email string `gorm:"type:varchar(100);uniqueIndex"`
}

// DB 用于连接和操作数据库
var DB *gorm.DB

// SetupTest 初始化测试数据库
func SetupTest() {
    // 使用 SQLite 内存数据库，适合测试
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }

    // 自动迁移
    db.AutoMigrate(&User{})

    DB = db
}

// TearDownTest 清理测试环境
func TearDownTest() {
    // 清理数据库
    DB.Migrator().DropTable(&User{})
}

// TestUsers 测试用户相关的功能
func TestUsers(t *testing.T) {
    SetupTest()
    defer TearDownTest()

    // 创建用户
    user := User{Name: "John Doe", Email: "john.doe@example.com"}
    result := DB.Create(&user)
    if result.Error != nil {
        t.Errorf("Failed to create user: %v", result.Error)
        return
    }

    // 查询用户
    var dbUser User
    DB.First(&dbUser, user.ID)
    if dbUser.ID != user.ID {
        t.Errorf("User not found, expected %v, got %v", user.ID, dbUser.ID)
    }

    // 更新用户
    dbUser.Name = "Jane Doe"
    result = DB.Save(&dbUser)
    if result.Error != nil {
        t.Errorf("Failed to update user: %v\, result.Error)
        return
    }

    // 验证更新结果
    var updatedUser User
    DB.First(&updatedUser, user.ID)
    if updatedUser.Name != "Jane Doe" {
        t.Errorf("User name not updated, expected %v, got %v", "Jane Doe", updatedUser.Name)
    }

    // 删除用户
    result = DB.Delete(&User{}, user.ID)
    if result.Error != nil {
        t.Errorf("Failed to delete user: %v\, result.Error)
        return
    }

    // 验证删除结果
    var deletedUser User
    DB.First(&deletedUser, user.ID)
    if deletedUser.ID != 0 {
        t.Errorf("User not deleted, expected 0, got %v", deletedUser.ID)
    }
}

func main() {
    // 仅在测试时运行
    _ = testing.Verbose()
    testing.Main(
        // 测试函数
        func(pat, opt string) (bool, error) {
            return true, nil
        },
    )
}
