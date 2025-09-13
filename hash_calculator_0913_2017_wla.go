// 代码生成时间: 2025-09-13 20:17:01
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "errors"
    "fmt"
    "os"
)

// HashCalculator 结构体用于封装哈希计算工具的方法
type HashCalculator struct {
    // 此处可以添加更多字段，以支持不同的哈希算法或配置
}

// NewHashCalculator 创建一个新的 HashCalculator 实例
func NewHashCalculator() *HashCalculator {
    return &HashCalculator{}
}

// CalculateSHA256 接收一个字符串并返回其SHA-256哈希值
func (h *HashCalculator) CalculateSHA256(input string) (string, error) {
    // 使用SHA-256算法计算哈希值
    hash := sha256.Sum256([]byte(input))
    // 将字节切片转换为十六进制字符串
    hexHash := hex.EncodeToString(hash[:])
    return hexHash, nil
}

// CalculateSHA256File 计算文件内容的SHA-256哈希值并返回
func (h *HashCalculator) CalculateSHA256File(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", errors.New("failed to open file: " + err.Error())
    }
    defer file.Close()

    hash := sha256.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", errors.New("failed to copy file content to hash: " + err.Error())
    }
    // 将哈希值转换为十六进制字符串
    hexHash := hex.EncodeToString(hash.Sum(nil))
    return hexHash, nil
}

func main() {
    // 创建哈希计算工具实例
    hashCalc := NewHashCalculator()

    // 计算字符串的哈希值
    stringHash, err := hashCalc.CalculateSHA256("Hello, World!")
    if err != nil {
        fmt.Println("Error calculating string hash: ", err)
    } else {
        fmt.Printf("String hash: %s
", stringHash)
    }

    // 计算文件的哈希值
    fileHash, err := hashCalc.CalculateSHA256File("example.txt")
    if err != nil {
        fmt.Println("Error calculating file hash: ", err)
    } else {
        fmt.Printf("File hash: %s
", fileHash)
    }
}
