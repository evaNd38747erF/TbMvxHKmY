// 代码生成时间: 2025-08-09 19:01:49
 * It includes error handling, comments, and follows GoLang best practices for maintainability and scalability.
 */

package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "fmt"
)

// EncryptionKey is the key used for AES encryption and decryption.
// It should be a 32-byte key for AES-256.
var EncryptionKey = []byte("YouShouldUseA32ByteLongKey")

// Encrypt encrypts the provided plaintext using AES-256-GCM.
func Encrypt(plaintext []byte) (string, error) {
    if len(plaintext)%aes.BlockSize != 0 {
        return "", errors.New("plaintext is not a multiple of the block size")
    }

    iv := make([]byte, aes.BlockSize)
    if _, err := rand.Read(iv); err != nil {
        return "", err
    }

    block, err := aes.NewCipher(EncryptionKey)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    encrypted := gcm.Seal(iv, iv, plaintext, nil)
    return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt decrypts the provided ciphertext using AES-256-GCM.
func Decrypt(ciphertext string) (string, error) {
    data, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }

    iv := data[:aes.BlockSize]
    ciphertext = data[aes.BlockSize:]

    block, err := aes.NewCipher(EncryptionKey)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
    if err != nil {
        return "", err
    }

    return string(plaintext), nil
}

func main() {
    plaintext := "Hello, World!"
    fmt.Println("Plaintext: ", plaintext)

    encrypted, err := Encrypt([]byte(plaintext))
    if err != nil {
        fmt.Println("Error encrypting: ", err)
        return
    }
    fmt.Println("Encrypted: ", encrypted)

    decrypted, err := Decrypt(encrypted)
    if err != nil {
        fmt.Println("Error decrypting: ", err)
        return
    }
    fmt.Println("Decrypted: ", decrypted)
}