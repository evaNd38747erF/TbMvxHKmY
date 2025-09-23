// 代码生成时间: 2025-09-23 08:40:55
package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "time"
    "github.com/PuerkitoBio/goquery"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// WebCrawler 结构体定义，用于存储网页内容抓取的工具
type WebCrawler struct {
    URL string
    DB  *gorm.DB
}

// NewWebCrawler 创建一个新的 WebCrawler 实例
func NewWebCrawler(url string, db *gorm.DB) *WebCrawler {
    return &WebCrawler{URL: url, DB: db}
}

// FetchContent 从给定的 URL 获取网页内容
func (wc *WebCrawler) FetchContent() (string, error) {
    resp, err := http.Get(wc.URL)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return "", err
    }
    
    // 假设我们想要抓取的是网页的全部HTML内容
    return doc.Text(), nil
}

// SaveContent 将抓取的内容保存到数据库
func (wc *WebCrawler) SaveContent(content string) error {
    // 这里需要定义一个模型来存储网页内容
    // 例如，一个简单的 Content 模型
    type Content struct {
        ID        uint      "gorm:AUTO_INCREMENT"
        URL       string    "gorm:unique"
        HTML      string    "gorm:Type:text"
        CreatedAt time.Time
    }
    
    // 创建一个新的 Content 实例并保存
    err := wc.DB.Create(&Content{URL: wc.URL, HTML: content, CreatedAt: time.Now()}).Error
    if err != nil {
        return err
    }
    
    return nil
}

func main() {
    // 连接数据库
    dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }
    defer db.Close()
    
    // 使用数据库迁移
    db.AutoMigrate(&Content{})
    
    // 创建 WebCrawler 实例
    crawler := NewWebCrawler("http://example.com", db)
    
    // 抓取网页内容
    content, err := crawler.FetchContent()
    if err != nil {
        log.Printf("failed to fetch content: %s", err)
        return
    }
    
    // 保存网页内容到数据库
    if err := crawler.SaveContent(content); err != nil {
        log.Printf("failed to save content: %s", err)
        return
    }
    
    fmt.Println("Web content has been successfully crawled and saved.")
}