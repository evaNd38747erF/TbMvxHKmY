// 代码生成时间: 2025-09-01 18:40:12
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// 定义User模型
type User struct {
    gorm.Model
    Name string `json:"name"`
    Email string `json:"email"`
}

// 初始化数据库
func initDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database"))
    }

    // 自动迁移模式，自动创建数据库表
    db.AutoMigrate(&User{})
    return db
}

// 新增用户
func createUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 插入数据到数据库
    db := initDB()
    if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// 获取用户列表
func getUsers(c *gin.Context) {
    db := initDB()
    var users []User
    if err := db.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}

// 更新用户
func updateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := initDB()
    if err := db.Model(&User{}).Where("id = ?", c.Param("id")).Update(user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// 删除用户
func deleteUser(c *gin.Context) {
    db := initDB()
    if err := db.Delete(&User{}, c.Param("id\)).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func main() {
    // 初始化路由
    router := gin.Default()

    // 用户相关接口
    router.POST("/users", createUser)
    router.GET("/users", getUsers)
    router.PUT("/users/:id", updateUser)
    router.DELETE("/users/:id", deleteUser)

    // 启动服务
    router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
