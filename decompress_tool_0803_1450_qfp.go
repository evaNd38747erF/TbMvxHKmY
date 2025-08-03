// 代码生成时间: 2025-08-03 14:50:06
package main

import (
    "archive/zip"
    "io"
    "log"
    "os"
    "path/filepath"
)

// DecompressTool 定义了压缩文件解压工具的结构
type DecompressTool struct {
    // srcPath 是压缩文件的路径
    srcPath string
    // destPath 是解压后文件的目标路径
    destPath string
}

// NewDecompressTool 创建并返回一个新的 DecompressTool 实例
func NewDecompressTool(srcPath, destPath string) *DecompressTool {
    return &DecompressTool{
        srcPath: srcPath,
        destPath: destPath,
    }
}

// Decompress 解压指定的压缩文件到目标路径
func (dt *DecompressTool) Decompress() error {
    // 打开压缩文件
    srcFile, err := os.Open(dt.srcPath)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    // 解压缩文件
    destZip, err := zip.OpenReader(srcFile.Name())
    if err != nil {
        return err
    }
    defer destZip.Close()

    // 创建目标目录
    if err := os.MkdirAll(dt.destPath, 0755); err != nil {
        return err
    }

    // 循环解压文件
    for _, file := range destZip.File {
        // 创建目标文件的路径
        destFilePath := filepath.Join(dt.destPath, file.Name)

        // 确保目标路径是文件而非目录
        if file.FileInfo().IsDir() {
            // 创建目录
            if err := os.MkdirAll(destFilePath, file.Mode()); err != nil {
                return err
            }
            continue
        }

        // 打开压缩包中的文件
        srcFileInZip, err := file.Open()
        if err != nil {
            return err
        }
        defer srcFileInZip.Close()

        // 创建目标文件
        if err := os.MkdirAll(filepath.Dir(destFilePath), 0755); err != nil {
            return err
        }
        destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
        if err != nil {
            return err
        }
        defer destFile.Close()

        // 复制文件内容
        _, err = io.Copy(destFile, srcFileInZip)
        if err != nil {
            return err
        }
    }

    return nil
}

func main() {
    // 使用示例
    srcPath := "path/to/your/compressed/file.zip"
    destPath := "path/to/your/destination/folder"
    decompressTool := NewDecompressTool(srcPath, destPath)

    // 执行解压操作
    if err := decompressTool.Decompress(); err != nil {
        log.Fatalf("Failed to decompress: %v", err)
    }
    log.Println("Decompression completed successfully.")
}