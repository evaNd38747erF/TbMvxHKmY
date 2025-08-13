// 代码生成时间: 2025-08-13 09:53:58
// url_validator.go

package main

import (
    "fmt"
    "net/http"
    "net/url"
    "time"
    "github.com/jinzhu/gorm"
    \_ "github.com/jinzhu/gorm/dialects/sqlite" // 导入SQLite驱动
)

// 定义URLValidator结构体
type URLValidator struct {
    // DB是*gorm.DB的引用，用于数据库操作
    DB *gorm.DB
}

// NewURLValidator初始化URLValidator
func NewURLValidator(db *gorm.DB) *URLValidator {
    return &URLValidator{DB: db}
}

// ValidateURL检查URL是否有效
func (uv *URLValidator) ValidateURL(urlStr string) (bool, error) {
    // 解析URL
    u, err := url.ParseRequestURI(urlStr)
    if err != nil {
        return false, fmt.Errorf("invalid URL format: %w", err)
    }

    // 检查协议是否有效
    if u.Scheme != "http" && u.Scheme != "https" {
        return false, fmt.Errorf("unsupported URL scheme: %s", u.Scheme)
    }

    // 发送HEAD请求检查URL是否可达
    resp, err := http.Head(u.String())
    if err != nil {
        return false, fmt.Errorf("failed to reach URL: %w", err)
    }
    defer resp.Body.Close()

    // 检查HTTP状态码是否表示有效
    if resp.StatusCode != http.StatusOK {
        return false, nil
    }

    // 如果所有检查都通过，则URL有效
    return true, nil
}

func main() {
    // 连接到SQLite数据库
    db, err := gorm.Open("sqlite3:validator.db", &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    // 自动迁移模式
    db.AutoMigrate(&URLValidator{})

    // 创建URLValidator实例
    validator := NewURLValidator(db)

    // 示例URL字符串
    urlStr := "https://www.example.com"

    // 验证URL是否有效
    valid, err := validator.ValidateURL(urlStr)
    if err != nil {
        fmt.Printf("Error validating URL: %v\
", err)
    } else {
        fmt.Printf("URL %s is %s.\
", urlStr, fmt.Sprintf("valid" if valid else "invalid"))
    }
}
