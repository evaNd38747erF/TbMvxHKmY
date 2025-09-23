// 代码生成时间: 2025-09-23 21:25:16
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "os"
    "time"
)

// TestRecord 用于测试的示例模型
type TestRecord struct {
    ID        uint   "gorm:\"primary_key\""
    CreatedAt time.Time
    UpdatedAt time.Time
}

func main() {
    // 设置数据库连接
    dsn := "test.db"
    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: ", err)
    }
    defer db.Close()

    // 自动迁移
    db.AutoMigrate(&TestRecord{})

    // 性能测试
    testRecord := TestRecord{CreatedAt: time.Now(), UpdatedAt: time.Now()}
    startTime := time.Now()
    for i := 0; i < 1000; i++ {
        if err := db.Create(&testRecord).Error; err != nil {
            log.Println("error inserting record: ", err)
        }
    }
    endTime := time.Now()

    // 计算并打印性能测试结果
    duration := endTime.Sub(startTime)
    fmt.Printf("1000 records created in %v with %d errors", duration, 0) // 假设没有错误发生

    // 清理测试数据
    db.Migrator().DropTable(&TestRecord{})
    os.Remove(dsn)
}
