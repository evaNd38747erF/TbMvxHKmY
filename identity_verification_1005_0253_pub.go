// 代码生成时间: 2025-10-05 02:53:20
package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User 代表用户模型
type User struct {
    gorm.Model
    Username string `gorm:"type:varchar(100);uniqueIndex"`
    Password string `gorm:"type:varchar(100)"`
}

// Database 连接配置
var db *gorm.DB
var err error

func main() {
    // 初始化数据库连接
    db, err = gorm.Open(sqlite.Open("identity.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("数据库连接失败：", err)
    }

    // 迁移数据库模式
    db.AutoMigrate(&User{})

    // 创建示例用户
    user := User{Username: "example", Password: "password123"}
    db.Create(&user)
    if err != nil {
        log.Fatal("创建用户失败：", err)
    }

    // 验证用户身份
    if err := verifyIdentity("example", "password123"); err != nil {
        log.Fatal("身份验证失败：", err)
    } else {
        fmt.Println("身份验证成功！")
    }
}

// verifyIdentity 验证用户身份
func verifyIdentity(username, password string) error {
    // 查询数据库中的用户
    user := User{}
    result := db.Where(&User{Username: username}).First(&user)
    if result.Error != nil {
        return result.Error
    }
    // 验证密码
    if user.Password != password {
        return fmt.Errorf("无效的用户名或密码")
    }
    return nil
}
