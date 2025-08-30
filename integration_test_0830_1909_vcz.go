// 代码生成时间: 2025-08-30 19:09:16
package main

import (
	"fmt"
	"os"
	"testing"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 定义用户模型
type User struct {
	gorm.Model
	Name string
	Age  uint
}

// DB 初始化数据库连接
var DB *gorm.DB

// Setup 初始化测试环境
func Setup() {
	var err error
	// 使用内存数据库进行测试
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	n := DB.AutoMigrate(&User{})
	fmt.Println("Auto migration: ", n)
}

// TestMain 设置测试入口
func TestMain(m *testing.M) {
	Setup()
	os.Exit(m.Run())
}

// TestCreateUser 测试用户创建
func TestCreateUser(t *testing.T) {
	// 给定
	user := User{Name: "John Doe", Age: 30}

	// 当创建用户
	result := DB.Create(&user)

	// 则数据库中应包含该用户
	var dbUser User
	DB.First(&dbUser, user.ID)
	if result.Error != nil {
		t.Errorf("failed to create user: %v", result.Error)
	} else if dbUser.Name != user.Name || dbUser.Age != user.Age {
		t.Errorf("user data mismatch, expected: %v, got: %v", user, dbUser)
	}
}

// TestUpdateUser 测试用户更新
func TestUpdateUser(t *testing.T) {
	// 给定
	user := User{Name: "Jane Doe", Age: 25}
	DB.Create(&user)

	// 当更新用户
	user.Name = "Jane Doe Updated"
	result := DB.Save(&user)

	// 则数据库中应包含更新后的用户
	var dbUser User
	DB.First(&dbUser, user.ID)
	if result.Error != nil {
		t.Errorf("failed to update user: %v", result.Error)
	} else if dbUser.Name != user.Name || dbUser.Age != user.Age {
		t.Errorf("user data mismatch, expected: %v, got: %v", user, dbUser)
	}
}

// TestDeleteUser 测试用户删除
func TestDeleteUser(t *testing.T) {
	// 给定
	user := User{Name: "Delete Me", Age: 40}
	DB.Create(&user)

	// 当删除用户
	result := DB.Delete(&user, user.ID)

	// 则数据库中不应包含该用户
	var dbUser User
	if result.Error != nil {
		t.Errorf("failed to delete user: %v", result.Error)
	} else if DB.First(&dbUser, user.ID).Error == nil {
		t.Errorf("user should be deleted, but found in database")
	}
}
