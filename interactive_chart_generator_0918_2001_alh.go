// 代码生成时间: 2025-09-18 20:01:26
// 交互式图表生成器程序
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/go-gorm/gorm"
    "github.com/jesseduffield/gh-cli/api"
)

// ChartData 用于存储图表数据的结构体
type ChartData struct {
    g    int `gorm:\"column:g\"` // X轴数据
    v    float64 `gorm:\"column:v\"` // Y轴数据
    desc string `gorm:\"column:desc\"` // 数据描述
}

// ChartService 处理图表数据的服务
type ChartService struct {
    db *gorm.DB
}

// NewChartService 初始化一个 ChartService
func NewChartService(db *gorm.DB) *ChartService {
    return &ChartService{db: db}
}

// AddChartData 添加图表数据
func (cs *ChartService) AddChartData(g int, v float64, desc string) error {
    data := ChartData{g: g, v: v, desc: desc}
    if err := cs.db.Create(&data).Error; err != nil {
        return err
    }
    return nil
}

// GetAllChartData 获取所有图表数据
func (cs *ChartService) GetAllChartData() ([]ChartData, error) {
    var data []ChartData
    if err := cs.db.Find(&data).Error; err != nil {
        return nil, err
    }
    return data, nil
}

// ChartController 处理HTTP请求的控制器
type ChartController struct {
    service *ChartService
}

// NewChartController 构造函数
func NewChartController(service *ChartService) *ChartController {
    return &ChartController{service: service}
}

// AddChartDataHandler 添加图表数据的HTTP处理器
func (cc *ChartController) AddChartDataHandler(c *gin.Context) {
    g := c.PostForm(\"g\")
    v := c.PostForm(\"v\")
    desc := c.PostForm(\"desc\")
    if err := cc.service.AddChartData(g, v, desc); err != nil {
        c.JSON(400, gin.H{\"error\": err.Error()})
        return
    }
    c.JSON(200, gin.H{\"message\": \"Chart data added successfully\"})
}

// GetAllChartDataHandler 获取所有图表数据的HTTP处理器
func (cc *ChartController) GetAllChartDataHandler(c *gin.Context) {
    data, err := cc.service.GetAllChartData()
    if err != nil {
        c.JSON(500, gin.H{\"error\": err.Error()})
        return
    }
    c.JSON(200, gin.H{\"data\": data})
}

func main() {
    db, err := gorm.Open(sqlite.Open(\"chart.db\"), &gorm.Config{})
    if err != nil {
        log.Fatalf(\"failed to connect database: %v\", err)
    }

    // 迁移数据库，确保ChartData表存在
    db.AutoMigrate(&ChartData{})

    router := gin.Default()
    chartService := NewChartService(db)
    chartController := NewChartController(chartService)

    router.POST(\"/addChartData\", chartController.AddChartDataHandler)
    router.GET(\"/getChartData\", chartController.GetAllChartDataHandler)

    // 启动服务
    if err := router.Run(\":8080\"); err != nil {
        log.Fatalf(\"failed to start server: %v\", err)
    }
}
