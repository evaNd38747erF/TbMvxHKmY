// 代码生成时间: 2025-08-22 01:16:55
package main

import (
    "crypto/rand"
    "fmt"
    "math/big"
    "os"
)

// RandomNumberGenerator is a struct that holds parameters for generating random numbers.
type RandomNumberGenerator struct {
    // Min is the minimum value of the generated random number.
    Min int
    // Max is the maximum value of the generated random number.
    Max int
}

// NewRandomNumberGenerator creates a new instance of RandomNumberGenerator with the given parameters.
func NewRandomNumberGenerator(min, max int) *RandomNumberGenerator {
    return &RandomNumberGenerator{Min: min, Max: max}
}

// GenerateInt64 generates a random int64 number within the range [Min, Max].
func (r *RandomNumberGenerator) GenerateInt64() (int64, error) {
    if r.Max <= r.Min {
        return 0, fmt.Errorf("invalid range: max must be greater than min")
    }
    num, err := rand.Int(rand.Reader, big.NewInt(int64(r.Max-r.Min+1)))
    if err != nil {
        return 0, fmt.Errorf("failed to generate random number: %w", err)
    }
    return int64(num.Int64()) + int64(r.Min), nil
}

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run random_number_generator.go <min> <max>")
        os.Exit(1)
    }
    min, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Invalid minimum value: must be an integer")
        os.Exit(1)
    }
    max, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Invalid maximum value: must be an integer")
        os.Exit(1)
    }
    
    rng := NewRandomNumberGenerator(min, max)
    result, err := rng.GenerateInt64()
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    fmt.Printf("Generated random number: %d
", result)
}