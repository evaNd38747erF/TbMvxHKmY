// 代码生成时间: 2025-09-07 23:30:57
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "os/exec"
    "runtime"
    "strings"
    "time"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql" // 导入对应的数据库驱动
)

// 定义系统性能监控工具的结构体
type SystemPerformanceMonitor struct {
    DB *sql.DB
}

// NewSystemPerformanceMonitor 初始化系统性能监控工具
func NewSystemPerformanceMonitor(dataSourceName string) (*SystemPerformanceMonitor, error) {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }
    return &SystemPerformanceMonitor{DB: db}, nil
}

// MonitorSystemPerformance 监控系统性能
func (spm *SystemPerformanceMonitor) MonitorSystemPerformance() error {
    // 检查数据库连接
    if err := spm.DB.Ping(); err != nil {
        return fmt.Errorf("database ping failed: %w", err)
    }

    // 收集系统信息
    systemInfo, err := collectSystemInfo()
    if err != nil {
        return fmt.Errorf("failed to collect system info: %w", err)
    }

    // 将系统信息写入数据库
    if err := writeSystemInfoToDatabase(spm.DB, systemInfo); err != nil {
        return fmt.Errorf("failed to write system info to database: %w", err)
    }

    return nil
}

// collectSystemInfo 收集系统性能信息
func collectSystemInfo() (map[string]string, error) {
    systemInfo := make(map[string]string)

    // 收集CPU信息
    cpuInfo, err := exec.Command("top", "-b", "-n1", "-H").Output()
    if err != nil {
        return nil, err
    }
    systemInfo["CPU"] = strings.TrimSpace(string(cpuInfo))

    // 收集内存信息
    memInfo, err := exec.Command("free", "-m").Output()
    if err != nil {
        return nil, err
    }
    systemInfo["Memory"] = strings.TrimSpace(string(memInfo))

    // 收集磁盘信息
    diskInfo, err := exec.Command("df", "-h").Output()
    if err != nil {
        return nil, err
    }
    systemInfo["Disk"] = strings.TrimSpace(string(diskInfo))

    return systemInfo, nil
}

// writeSystemInfoToDatabase 将系统信息写入数据库
func writeSystemInfoToDatabase(db *sql.DB, systemInfo map[string]string) error {
    // 创建表
    createTableSQL := `CREATE TABLE IF NOT EXISTS system_performance (
        id INT AUTO_INCREMENT PRIMARY KEY,
        cpu TEXT,
        memory TEXT,
        disk TEXT,
        timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`
    _, err := db.Exec(createTableSQL)
    if err != nil {
        return err
    }

    // 插入数据
    insertSQL := `INSERT INTO system_performance (cpu, memory, disk) VALUES (?, ?, ?)`
    _, err = db.Exec(insertSQL, systemInfo["CPU"], systemInfo["Memory"], systemInfo["Disk"])
    return err
}

func main() {
    dataSourceName := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local"
    spm, err := NewSystemPerformanceMonitor(dataSourceName)
    if err != nil {
        log.Fatalf("failed to create system performance monitor: %v", err)
    }
    defer spm.DB.Close()

    if err := spm.MonitorSystemPerformance(); err != nil {
        log.Printf("failed to monitor system performance: %v", err)
    } else {
        fmt.Println("System performance monitoring completed successfully")
    }
}