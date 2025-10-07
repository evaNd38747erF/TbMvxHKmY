// 代码生成时间: 2025-10-07 18:16:50
package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

//定义物流信息结构
type Logistics struct {
    gorm.Model
    TrackingNumber string `gorm:"primaryKey"`
    CurrentLocation string
    Status          string
}

// 初始化数据库连接
var db *gorm.DB
var err error

func init() {
    db, err = gorm.Open(sqlite.Open("logistics.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    // 自动迁移模式
    db.AutoMigrate(&Logistics{})
}

// 添加物流信息
func AddLogistics(logistics *Logistics) error {
    result := db.Create(&logistics)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// 更新物流信息
func UpdateLogistics(logistics *Logistics) error {
    result := db.Save(&logistics)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// 获取物流信息
func GetLogistics(trackingNumber string) (*Logistics, error) {
    var logistics Logistics
    result := db.First(&logistics, TrackingNumber)
    if result.Error != nil {
        return nil, result.Error
    }
    return &logistics, nil
}

// 删除物流信息
func DeleteLogistics(trackingNumber string) error {
    result := db.Delete(&Logistics{}, TrackingNumber)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func main() {
    // 添加物流信息示例
    logistics := Logistics{TrackingNumber: "123456789", CurrentLocation: "Shanghai", Status: "In Transit"}
    if err := AddLogistics(&logistics); err != nil {
        fmt.Println("Error adding logistics: ", err)
    }

    // 更新物流信息示例
    logistics.Status = "Delivered"
    if err := UpdateLogistics(&logistics); err != nil {
        fmt.Println("Error updating logistics: ", err)
    }

    // 获取物流信息示例
    logisticsInfo, err := GetLogistics("123456789")
    if err != nil {
        fmt.Println("Error getting logistics: ", err)
    } else {
        fmt.Printf("Logistics Info: Tracking Number: %s, Location: %s, Status: %s
"
            , logisticsInfo.TrackingNumber, logisticsInfo.CurrentLocation, logisticsInfo.Status)
    }

    // 删除物流信息示例
    if err := DeleteLogistics("123456789"); err != nil {
        fmt.Println("Error deleting logistics: ", err)
    }
}
