// 代码生成时间: 2025-09-10 23:12:18
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "os"
)

// HashCalculator 定义一个哈希值计算工具的结构体
type HashCalculator struct {}

// NewHashCalculator 创建一个新的哈希值计算工具实例
func NewHashCalculator() *HashCalculator {
    return &HashCalculator{}
}

// CalculateHash 计算给定字符串的SHA-256哈希值
func (h *HashCalculator) CalculateHash(input string) (string, error) {
    // 将输入字符串转换为字节切片
    inputBytes := []byte(input)

    // 使用SHA-256算法计算哈希值
    hash := sha256.Sum256(inputBytes)

    // 将哈希值的字节切片转换为十六进制字符串
    hashString := hex.EncodeToString(hash[:])

    return hashString, nil
}

func main() {
    // 创建哈希值计算工具实例
    calculator := NewHashCalculator()

    // 从标准输入读取字符串
    var input string
    fmt.Print("Enter a string to calculate its hash: ")
    fmt.Scanln(&input)

    // 计算哈希值
    hash, err := calculator.CalculateHash(input)
    if err != nil {
        fmt.Println("Error calculating hash: ", err)
        os.Exit(1)
    }

    // 输出哈希值
    fmt.Println("The hash of the input string is: ", hash)
}