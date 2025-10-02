// 代码生成时间: 2025-10-02 21:47:56
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Block represents a block in the blockchain
type Block struct {
    gorm.Model
    Index        int
    Transactions []Transaction
    Nonce        string `gorm:"index"`
    Hash         string
    PreviousHash string
}

// Transaction represents a transaction in the blockchain
type Transaction struct {
    gorm.Model
    From  string
    To    string
    Amount float64
}

// Blockchain represents the blockchain
type Blockchain struct {
    Blocks []Block
}

// NewBlockchain creates a new blockchain
func NewBlockchain() *Blockchain {
    return &Blockchain{
        Blocks: make([]Block, 0),
    }
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) error {
    var lastBlock Block
    if len(bc.Blocks) > 0 {
        lastBlock = bc.Blocks[len(bc.Blocks)-1]
    }

    newBlock := Block{
        Index:        len(bc.Blocks) + 1,
        Transactions: []Transaction{{From: "Genesis", To: data, Amount: 1.0}},
        PreviousHash: lastBlock.Hash,
    }

    if err := bc.MineBlock(&newBlock); err != nil {
        return err
    }

    bc.Blocks = append(bc.Blocks, newBlock)
    return nil
}

// MineBlock adds a nonce to the block and calculates the hash
func (bc *Blockchain) MineBlock(block *Block) error {
    block.Nonce = ""
    block.Hash = CalculateHash(block)
    for !block.IsValid() {
        block.Nonce += "1"
        block.Hash = CalculateHash(block)
    }
    return nil
}

// IsValid checks if the block is valid
func (block *Block) IsValid() bool {
    calculatedHash := CalculateHash(block)
    return calculatedHash == block.Hash
}

// CalculateHash calculates the hash of the block
func CalculateHash(block *Block) string {
    // This is a simple hash function for demonstration purposes
    // In a real-world scenario, a more complex hash function would be used
    return fmt.Sprintf("%x", block.Index)
}

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database"))
    }
    defer db.Close()

    db.AutoMigrate(&Block{}, &Transaction{})

    bc := NewBlockchain()
    if err := bc.AddBlock("Alice"); err != nil {
        fmt.Println("Error adding block: ", err)
        return
    }
    if err := bc.AddBlock("Bob"); err != nil {
        fmt.Println("Error adding block: ", err)
        return
    }

    fmt.Println("Blockchain:")
    for _, block := range bc.Blocks {
        fmt.Printf("Block: %+v
", block)
    }
}
