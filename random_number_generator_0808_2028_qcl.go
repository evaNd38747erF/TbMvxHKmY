// 代码生成时间: 2025-08-08 20:28:37
package main

import (
    "crypto/rand"
    "fmt"
    "math/big"
)

// RandomNumberGenerator defines the structure for generating random numbers.
type RandomNumberGenerator struct {
    // Min is the minimum value of the random number.
    Min int
    // Max is the maximum value of the random number.
    Max int
}

// NewRandomNumberGenerator creates a new instance of RandomNumberGenerator.
func NewRandomNumberGenerator(min, max int) *RandomNumberGenerator {
    return &RandomNumberGenerator{Min: min, Max: max}
}

// GenerateRandomNumber generates a random number between Min and Max.
func (r *RandomNumberGenerator) GenerateRandomNumber() (int, error) {
    // Ensure the max is greater than min for a valid range.
    if r.Max <= r.Min {
        return 0, fmt.Errorf("max must be greater than min")
    }
    
    // Generate a random number using crypto/rand for cryptographic security.
    randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(r.Max-r.Min+1)))
    if err != nil {
        return 0, err
    }
    
    // Add the min value to the generated random number to adjust the range.
    return int(randomNumber.Int64()) + r.Min, nil
}

func main() {
    // Create a new random number generator with a range of 1 to 100.
    rng := NewRandomNumberGenerator(1, 100)
    
    // Generate and print 5 random numbers.
    for i := 0; i < 5; i++ {
        randomNumber, err := rng.GenerateRandomNumber()
        if err != nil {
            fmt.Println("Error generating random number: ", err)
        } else {
            fmt.Printf("Random Number %d: %d
", i+1, randomNumber)
        }
    }
}