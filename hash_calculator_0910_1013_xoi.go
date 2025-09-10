// 代码生成时间: 2025-09-10 10:13:56
package main

import (
    "crypto/sha256"
    "fmt"
    "os"
    "log"
)

// HashCalculator 结构体包含输入数据
type HashCalculator struct {
    input string
}

// NewHashCalculator 创建一个新的 HashCalculator 实例
func NewHashCalculator(input string) *HashCalculator {
    return &HashCalculator{input: input}
}

// CalculateHash 计算输入数据的 SHA-256 哈希值
func (hc *HashCalculator) CalculateHash() (string, error) {
    if hc.input == "" {
        return "", fmt.Errorf("input cannot be empty")
    }
    
    // 将输入数据转换为字节切片
    inputBytes := []byte(hc.input)
    
    // 创建一个新的SHA-256哈希器
    hash := sha256.New()
    
    // 写入输入数据到哈希器
    _, err := hash.Write(inputBytes)
    if err != nil {
        return "", err
    }
    
    // 计算最终的哈希值
    result := hash.Sum(nil)
    
    // 将哈希值转换为十六进制字符串
    return fmt.Sprintf("%x", result), nil
}

func main() {
    // 从命令行参数获取输入
    if len(os.Args) != 2 {
        log.Fatal("Usage: hash_calculator <input>")
    }
    inputData := os.Args[1]
    
    // 创建 HashCalculator 实例
    calculator := NewHashCalculator(inputData)
    
    // 计算哈希值
    hash, err := calculator.CalculateHash()
    if err != nil {
        log.Fatalf("Error calculating hash: %s", err)
    }
    
    // 打印哈希值
    fmt.Printf("The hash of '%s' is: %s
", inputData, hash)
}