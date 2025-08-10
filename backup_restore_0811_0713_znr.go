// 代码生成时间: 2025-08-11 07:13:18
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "os"
    "path/filepath"
    "time"
)

// BackupRestore 定义了备份和恢复操作的结构
type BackupRestore struct {
    DB    *gorm.DB
    Path  string
    Backup bool
}

// NewBackupRestore 创建BackupRestore实例
func NewBackupRestore(db *gorm.DB, path string) *BackupRestore {
    return &BackupRestore{DB: db, Path: path}
}

// Backup 执行数据库备份操作
func (br *BackupRestore) Backup() error {
    if br.Backup {
        return fmt.Errorf("backup is already in progress")
    }
    br.Backup = true
    defer func() { br.Backup = false }()

    backupPath := fmt.Sprintf("%s/backup_%s.db", br.Path, time.Now().Format("20060102150405"))
    file, err := os.Create(backupPath)
    if err != nil {
        return fmt.Errorf("failed to create backup file: %w", err)
    }
    defer file.Close()

    return br.DB.(*gorm.DB).DB().Backup(file, nil)
}

// Restore 执行数据库恢复操作
func (br *BackupRestore) Restore(backupPath string) error {
    file, err := os.Open(backupPath)
    if err != nil {
        return fmt.Errorf("failed to open backup file: %w", err)
    }
    defer file.Close()

    return br.DB.(*gorm.DB).DB().CopyFrom(file, nil)
}

func main() {
    // 连接数据库
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Printf("failed to connect database: %v", err)
        return
    }
    defer db.Close()

    // 自动迁移模式
    db.AutoMigrate(&User{})

    // 创建BackupRestore实例
    backupRestore := NewBackupRestore(db, "./backups")

    // 创建备份目录（如果不存在）
    if _, err := os.Stat(backupRestore.Path); os.IsNotExist(err) {
        if err := os.MkdirAll(backupRestore.Path, 0755); err != nil {
            fmt.Printf("failed to create backup directory: %v", err)
            return
        }
    }

    // 执行备份
    if err := backupRestore.Backup(); err != nil {
        fmt.Printf("backup failed: %v", err)
    } else {
        fmt.Println("backup succeeded")
    }

    // 执行恢复（这里假设我们已经有了一个备份文件）
    backupFile := filepath.Join(backupRestore.Path, "backup_20240302150045.db")
    if err := backupRestore.Restore(backupFile); err != nil {
        fmt.Printf("restore failed: %v", err)
    } else {
        fmt.Println("restore succeeded")
    }
}

// User 定义用户模型
type User struct {
    gorm.Model
    Name string
}
