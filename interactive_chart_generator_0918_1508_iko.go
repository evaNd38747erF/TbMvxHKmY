// 代码生成时间: 2025-09-18 15:08:55
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql" // 导入MySQL驱动
)

// ChartData 表示图表数据的结构体
type ChartData struct {
    ID        uint      `gorm:"primary_key"`
    Category  string    `gorm:"size:255"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// DBConfig 数据库配置
type DBConfig struct {
    Username string
    Password string
    Protocol string
    Host     string
    Port     string
    DBName  string
}

// ChartService 负责图表的业务逻辑
type ChartService struct {
    db *gorm.DB
}

// NewChartService 创建一个新的图表服务实例
func NewChartService(cfg DBConfig) (*ChartService, error) {
    connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
    db, err := gorm.Open(
        gorm.Open(connStr),
        &gorm.Config{
            // 配置GORM的日志输出
        }).SetLogger(logger{}),
    )
    if err != nil {
        return nil, err
    }

    // 迁移数据库，确保ChartData表存在
    db.AutoMigrate(&ChartData{})

    return &ChartService{db: db}, nil
}

// logger 用于自定义日志输出
type logger struct{}

// Print 打印日志信息
func (l logger) Print(v ...interface{}) {
    log.Printf("[INFO] %v", v...)
}

// AddChartData 添加图表数据
func (s *ChartService) AddChartData(category string) (uint, error) {
    data := ChartData{
        Category: category,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    if err := s.db.Create(&data).Error; err != nil {
        return 0, err
    }
    return data.ID, nil
}

// StartServer 启动HTTP服务器
func StartServer(service *ChartService) {
    r := gin.Default()
    r.POST("/addChartData", func(c *gin.Context) {
        var req struct {
            Category string `json:"category" binding:"required"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        id, err := service.AddChartData(req.Category)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{
            "id":  id,
            "data": req.Category,
        })
    })

    r.Run(":8080\) // 监听并在 0.0.0.0:8080 上启动服务
}

func main() {
    cfg := DBConfig{
        Username: "your_username",
        Password: "your_password",
        Protocol: "mysql",
        Host:     "127.0.0.1",
        Port:     "3306",
        DBName:   "your_dbname",
    }

    service, err := NewChartService(cfg)
    if err != nil {
        log.Fatalf("Failed to initialize chart service: %v", err)
    }

    StartServer(service)
}