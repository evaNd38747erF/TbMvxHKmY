// 代码生成时间: 2025-09-13 14:24:00
package main

import (
    "fmt"
    "gorm.io/driver/sqlite" // 使用SQLite作为数据库
    "gorm.io/gorm"
    "log"
    "os"
)

// AuditLog 定义安全审计日志的结构
type AuditLog struct {
    ID        uint   `gorm:"primaryKey"` // 主键
    CreatedAt string `gorm:"type:datetime"` // 创建时间
    UserID    uint   // 用户ID
    Action    string // 操作行为
    Details   string // 操作详情
}

// AuditLogService 定义审计日志服务
type AuditLogService struct {
    db *gorm.DB
}

// NewAuditLogService 初始化审计日志服务
func NewAuditLogService() *AuditLogService {
    db, err := gorm.Open(sqlite.Open("audit.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database", err)
    }
    // AutoMigrate 自动迁移数据库
    db.AutoMigrate(&AuditLog{})
    return &AuditLogService{db: db}
}

// Log 记录审计日志
func (service *AuditLogService) Log(userID uint, action, details string) error {
    // 创建审计日志实体
    log := AuditLog{
        CreatedAt: time.Now().Format(`2006-01-02 15:04:05`),
        UserID:    userID,
        Action:    action,
        Details:   details,
    }
    // 保存到数据库
    result := service.db.Create(&log)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func main() {
    // 初始化审计日志服务
    service := NewAuditLogService()
    // 记录审计日志
    err := service.Log(1, "登录", "用户成功登录系统")
    if err != nil {
        fmt.Printf("Error logging audit log: %s", err)
        os.Exit(1)
    }
    fmt.Println("Audit log recorded successfully")
}