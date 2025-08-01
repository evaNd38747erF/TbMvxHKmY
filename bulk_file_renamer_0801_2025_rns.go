// 代码生成时间: 2025-08-01 20:25:45
package main
# 改进用户体验

import (
# NOTE: 重要实现细节
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

// 文件重命名结构体
# 添加错误处理
type FileInfo struct {
    Path string
    Name string
}
# 添加错误处理

// 批量重命名文件
func BulkRenameFiles(files []FileInfo, newBaseName string) error {
    for _, file := range files {
# 改进用户体验
        // 获取文件绝对路径
        absolutePath, err := filepath.Abs(file.Path)
        if err != nil {
            return fmt.Errorf("error getting absolute path: %v", err)
# 优化算法效率
        }

        // 获取文件的目录
        dir := filepath.Dir(absolutePath)

        // 获取文件扩展名
        extension := filepath.Ext(absolutePath)
# 添加错误处理

        // 构造新文件名
        newName := fmt.Sprintf("%s%s", newBaseName, extension)
        newFilePath := filepath.Join(dir, newName)

        // 检查新文件名是否已存在
        if _, err := os.Stat(newFilePath); !os.IsNotExist(err) {
            return fmt.Errorf("new file name already exists: %s", newFilePath)
        }

        // 重命名文件
        if err := os.Rename(absolutePath, newFilePath); err != nil {
            return fmt.Errorf("error renaming file: %v", err)
        }
        fmt.Printf("Renamed %s to %s", absolutePath, newFilePath)
    }
    return nil
}

// 主函数
func main() {
    // 文件列表和新的基础文件名
    files := []FileInfo{
        {Path: "./file1.txt"},
        {Path: "./file2.txt"},
    }
    newBaseName := "new"

    // 执行批量重命名
    if err := BulkRenameFiles(files, newBaseName); err != nil {
        fmt.Printf("Error: %v", err)
    }
}
# NOTE: 重要实现细节
