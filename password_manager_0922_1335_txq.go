// 代码生成时间: 2025-09-22 13:35:28
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "log"
)

// PasswordManager 是一个密码加密解密工具的结构体
type PasswordManager struct {
    key []byte // 用于加密和解密的密钥
}

// NewPasswordManager 创建一个新的 PasswordManager 实例
func NewPasswordManager(key []byte) (*PasswordManager, error) {
    if len(key) != 32 {
        return nil, errors.New("密钥长度必须为32字节")
    }
    return &PasswordManager{key: key}, nil
}

// Encrypt 使用AES-256-GCM加密密码
func (pm *PasswordManager) Encrypt(plainText string) (string, error) {
    block, err := aes.NewCipher(pm.key)
    if err != nil {
        return "", err
    }

    nonce := make([]byte, 12) // AES-GCM需要12字节的nonce
    if _, err := rand.Read(nonce); err != nil {
        return "", err
    }

    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    encrypted := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
    return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt 使用AES-256-GCM解密密码
func (pm *PasswordManager) Decrypt(cipherText string) (string, error) {
    encryptedData, err := base64.StdEncoding.DecodeString(cipherText)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(pm.key)
    if err != nil {
        return "", err
    }

    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonceSize := aesGCM.NonceSize()
    if len(encryptedData) < nonceSize {
        return "", errors.New("密文长度不足")
    }

    nonce, cipherText := encryptedData[:nonceSize], encryptedData[nonceSize:]
    plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
    if err != nil {
        return "", err
    }
    return string(plainText), nil
}

func main() {
    key := []byte("your-32-byte-long-secret-key") // 密钥应该是32字节长
    manager, err := NewPasswordManager(key)
    if err != nil {
        log.Fatalf("创建密码管理器失败: %v", err)
    }

    plainText := "secretpassword"
    encrypted, err := manager.Encrypt(plainText)
    if err != nil {
        log.Fatalf("密码加密失败: %v", err)
    }
    log.Printf("加密后: %s", encrypted)

    decrypted, err := manager.Decrypt(encrypted)
    if err != nil {
        log.Fatalf("密码解密失败: %v", err)
    }
    log.Printf("解密后: %s", decrypted)
}
