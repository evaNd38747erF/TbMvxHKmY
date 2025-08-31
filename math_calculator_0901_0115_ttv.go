// 代码生成时间: 2025-09-01 01:15:32
package main

import (
    "fmt"
    "math"
)

// MathCalculator 包含了执行数学计算的方法
type MathCalculator struct {
    // 这里可以添加属性，例如存储计算历史等
}

// NewMathCalculator 创建并返回一个新的 MathCalculator 实例
func NewMathCalculator() *MathCalculator {
    return &MathCalculator{}
}

// Add 执行加法运算
func (mc *MathCalculator) Add(a, b float64) (float64, error) {
    sum := a + b
    return sum, nil // 这里返回nil表示没有错误
}

// Subtract 执行减法运算
func (mc *MathCalculator) Subtract(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("math: subtract by zero")
    }
    difference := a - b
    return difference, nil
}

// Multiply 执行乘法运算
func (mc *MathCalculator) Multiply(a, b float64) (float64, error) {
    product := a * b
    return product, nil
}

// Divide 执行除法运算
func (mc *MathCalculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("math: divide by zero")
    }
    quotient := a / b
    return quotient, nil
}

// Main function demonstrates usage of MathCalculator
func main() {
    // 创建 MathCalculator 实例
    calculator := NewMathCalculator()

    // 进行加法运算
    sum, err := calculator.Add(10, 5)
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("Sum: %.2f
", sum)
    }

    // 进行减法运算
    difference, err := calculator.Subtract(10, 5)
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("Difference: %.2f
", difference)
    }

    // 进行乘法运算
    product, err := calculator.Multiply(10, 5)
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("Product: %.2f
", product)
    }

    // 进行除法运算
    quotient, err := calculator.Divide(10, 5)
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("Quotient: %.2f
", quotient)
    }

    // 执行除以零的除法运算，以展示错误处理
    _, err = calculator.Divide(10, 0)
    if err != nil {
        fmt.Println("Error: ", err)
    }
}
