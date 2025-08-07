// 代码生成时间: 2025-08-08 07:48:32
package main

import (
    "crypto/sha256"
    "fmt"
    "log"
    "os"
)

// HashCalculator 结构体用于哈希值计算工具
type HashCalculator struct {
    // no fields are needed for this tool
}

// NewHashCalculator 创建一个新的 HashCalculator 实例
func NewHashCalculator() *HashCalculator {
    return &HashCalculator{}
}

// CalculateSHA256 计算给定字符串的 SHA-256 哈希值
func (hc *HashCalculator) CalculateSHA256(input string) (string, error) {
    if input == "" {
        return "", fmt.Errorf("input string is empty")
    }
    h := sha256.New()
    if _, err := h.Write([]byte(input)); err != nil {
        return "", err
    }
    hash := h.Sum(nil)
    return fmt.Sprintf("%x", hash), nil
}

func main() {
    // 创建哈希计算器实例
    hc := NewHashCalculator()

    // 从命令行参数获取输入字符串
    if len(os.Args) != 2 {
        log.Fatalf("Usage: %s <input string>", os.Args[0])
    }
    input := os.Args[1]

    // 计算 SHA-256 哈希值
    sha256Hash, err := hc.CalculateSHA256(input)
    if err != nil {
        log.Fatalf("Error calculating SHA-256 hash: %s", err)
    }

    // 输出结果
    fmt.Printf("The SHA-256 hash of '%s' is: %s
", input, sha256Hash)
}
