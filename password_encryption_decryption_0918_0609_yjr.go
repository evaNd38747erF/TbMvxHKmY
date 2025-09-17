// 代码生成时间: 2025-09-18 06:09:18
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "fmt"
    "io"
    "log"
)

// EncryptionKey represents the key for encryption and decryption.
// It should be kept secret and only known to the people who need to encrypt or decrypt messages.
var EncryptionKey = []byte("your-256-bit-encryption-key")

// Encrypt encrypts the given plaintext using AES-256-GCM.
func Encrypt(plaintext []byte) (string, error) {
    // Generate a random nonce of 12 bytes
    var nonce [12]byte
    if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
        return "", err
    }
    
    // Create a block, give it the key
    block, err := aes.NewCipher(EncryptionKey)
    if err != nil {
        return "", err
    }
    
    // Get the block size
    blockSize := block.BlockSize()
    
    // Padding
    origData := PKCS7Padding(plaintext, blockSize)
    
    // Encrypt the data
    if origData == nil {
        return "", errors.New("padding error")
    }
    
    // Create the cipher text
    cipherText := make([]byte, aes.BlockSize+len(origData))
    
    // Copy the nonce to the beginning of the cipher text
    copy(cipherText, nonce[:])
    
    // Copy the original data to the end of the cipher text
    copy(cipherText[aes.BlockSize:], origData)
    
    // Actually encrypt the data
    mode := cipher.NewGCM(block)
    if mode == nil {
        return "", errors.New("aes gcm new error")
    }
    
    // Must be unique for each encryption with the same key
    mode.Seal(cipherText[aes.BlockSize:], nonce[:], origData, nil)
    
    // Encode the result to base64
    return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts the given ciphertext using AES-256-GCM.
func Decrypt(ciphertext string) ([]byte, error) {
    // Decode the message from base64
    cipherText, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return nil, err
    }
    
    // Get the nonce size which is the first 12 bytes
    nonceSize := 12
    
    // Extract the nonce from the cipher text
    var nonce [12]byte
    copy(nonce[:], cipherText[:nonceSize])
    
    // Extract the cipher text without the nonce
    cipherText = cipherText[nonceSize:]
    
    // Create a block, give it the key
    block, err := aes.NewCipher(EncryptionKey)
    if err != nil {
        return nil, err
    }
    
    // Get the block size
    blockSize := block.BlockSize()
    
    // Decrypt the data
    mode := cipher.NewGCM(block)
    if mode == nil {
        return nil, errors.New("aes gcm new error")
    }
    
    // Decrypt and get the original data
    origData := cipherText
    if len(origData) < blockSize {
        return nil, errors.New("ciphertext too short")
    }
    
    // Decrypt the data
    mode.Open(origData, nonce[:], origData, nil)
    
    // Unpadding
    origData = PKCS7UnPadding(origData)
    
    return origData, nil
}

// PKCS7Padding is a function to pad the plaintext to be a multiple of the block size.
func PKCS7Padding(src []byte, blockSize int) []byte {
    padding := blockSize - len(src)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(src, padtext...)
}

// PKCS7UnPadding is a function to remove the padding from the plaintext.
func PKCS7UnPadding(src []byte) []byte {
    length := len(src)
    unpadding := int(src[length-1])
    return src[:(length - unpadding)]
}

func main() {
    // Example usage of Encrypt and Decrypt functions
    plaintext := []byte("Hello, World!")
    encrypted, err := Encrypt(plaintext)
    if err != nil {
        log.Fatalf("Encryption error: %v", err)
    }
    fmt.Println("Encrypted: ", encrypted)

    decrypted, err := Decrypt(encrypted)
    if err != nil {
        log.Fatalf("Decryption error: %v", err)
    }
    fmt.Println("Decrypted: ", string(decrypted))
}