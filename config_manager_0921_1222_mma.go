// 代码生成时间: 2025-09-21 12:22:23
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Config represents the configuration structure
type Config struct {
    ID       uint   `gorm:"primaryKey"`
    Key      string
    Value    string
    Modified string // Assuming the modified time is stored as a string
}

// ConfigManager handles operations related to configuration management
type ConfigManager struct {
    db *gorm.DB
}

// NewConfigManager creates a new instance of ConfigManager
func NewConfigManager(dbPath string) (*ConfigManager, error) {
    // Connect to the SQLite database
    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema
    db.AutoMigrate(&Config{})

    return &ConfigManager{db: db}, nil
}

// GetConfig retrieves a configuration value by key
func (cm *ConfigManager) GetConfig(key string) (string, error) {
    var config Config
    // Attempt to find the configuration item by key
    if err := cm.db.First(&config, Config{Key: key}).Error; err != nil {
        return "", err
    }
    return config.Value, nil
}

// SetConfig sets or updates a configuration value
func (cm *ConfigManager) SetConfig(key, value string) error {
    var config Config
    // Check if the configuration already exists
    if err := cm.db.First(&config, Config{Key: key}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return err
    }

    if errors.Is(err, gorm.ErrRecordNotFound) {
        // Create new configuration item
        config = Config{Key: key, Value: value}
        if err := cm.db.Create(&config).Error; err != nil {
            return err
        }
    } else {
        // Update existing configuration item
        config.Value = value
        if err := cm.db.Save(&config).Error; err != nil {
            return err
        }
    }
    return nil
}

func main() {
    // Define the database path
    dbPath := filepath.Join(os.Getenv("HOME"), "config.db")

    // Create a new config manager
    cm, err := NewConfigManager(dbPath)
    if err != nil {
        log.Fatalf("Failed to create config manager: %v", err)
    }

    // Example usage: Set a configuration value
    if err := cm.SetConfig("example_key", "example_value"); err != nil {
        log.Printf("Failed to set config: %v", err)
    }

    // Example usage: Get a configuration value
    configValue, err := cm.GetConfig("example_key")
    if err != nil {
        log.Printf("Failed to get config: %v", err)
    } else {
        fmt.Println("Config value: ", configValue)
    }
}
