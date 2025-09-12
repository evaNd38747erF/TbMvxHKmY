// 代码生成时间: 2025-09-13 02:37:31
package main

import (
    "crypto/rand"
    "encoding/binary"
    "fmt"
    "log"
    "math/big"
    "time"
)

// RandomNumberGenerator defines the structure for random number generation.
type RandomNumberGenerator struct {
    // Seed is the initial seed for generating random numbers.
    Seed int64
}

// NewRandomNumberGenerator creates a new RandomNumberGenerator instance.
func NewRandomNumberGenerator(seed int64) *RandomNumberGenerator {
    return &RandomNumberGenerator{
        Seed: seed,
    }
}

// GenerateRandomInt64 generates a random int64 number.
func (r *RandomNumberGenerator) GenerateRandomInt64() (int64, error) {
    max := big.NewInt(1<<63 - 1) // Max int64 value
    n, err := rand.Int(rand.Reader, max)
    if err != nil {
        return 0, err
    }
    return n.Int64(), nil
}

// GenerateRandomFloat64 generates a random float64 number between 0 and 1.
func (r *RandomNumberGenerator) GenerateRandomFloat64() (float64, error) {
    bytes := make([]byte, 8)
    if _, err := rand.Read(bytes); err != nil {
        return 0, err
    }
    return float64(binary.LittleEndian.Uint64(bytes)) / (1<<64 - 1), nil
}

func main() {
    // Initialize the random number generator with a seed for reproducibility.
    rng := NewRandomNumberGenerator(time.Now().UnixNano())
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Recovered in GenerateRandomInt64: %v", r)
        }
    }()

    // Generate and print a random int64 number.
    if randomInt, err := rng.GenerateRandomInt64(); err != nil {
        log.Fatalf("Error generating random int64: %v", err)
    } else {
        fmt.Printf("Generated random int64: %d
", randomInt)
    }

    // Generate and print a random float64 number.
    if randomFloat, err := rng.GenerateRandomFloat64(); err != nil {
        log.Fatalf("Error generating random float64: %v", err)
    } else {
        fmt.Printf("Generated random float64: %f
", randomFloat)
    }
}
