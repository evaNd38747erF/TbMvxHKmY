// 代码生成时间: 2025-09-20 14:53:57
package main

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Config represents a configuration setting
type Config struct {
    gorm.Model
    Key   string `gorm:"primaryKey"`
    Value string
}

// ConfigManager struct to manage configurations
type ConfigManager struct {
    db *gorm.DB
}

// NewConfigManager creates a new ConfigManager instance with a SQLite database
func NewConfigManager(dbPath string) (*ConfigManager, error) {
    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema
    db.AutoMigrate(&Config{})

    return &ConfigManager{db: db}, nil
}

// LoadConfig loads a configuration setting by key
func (m *ConfigManager) LoadConfig(key string) (string, error) {
    var config Config
    result := m.db.First(&config, Config{Key: key})
    if result.Error != nil {
        return "", result.Error
    }
    return config.Value, nil
}

// SaveConfig saves a configuration setting
func (m *ConfigManager) SaveConfig(key, value string) error {
    config := Config{Key: key}
    result := m.db.FirstOrCreate(&config, Config{Key: key})
    if result.Error != nil {
        return result.Error
    }
    config.Value = value
    result = m.db.Save(&config)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// UpdateConfig updates an existing configuration setting
func (m *ConfigManager) UpdateConfig(key, value string) error {
    var config Config
    result := m.db.First(&config, Config{Key: key})
    if result.Error != nil {
        return result.Error
    }
    config.Value = value
    result = m.db.Save(&config)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func main() {
    // Example usage of ConfigManager
    dbPath := "config.db"
    configManager, err := NewConfigManager(dbPath)
    if err != nil {
        log.Fatalf("Failed to create ConfigManager: %v", err)
    }

    // Save configuration settings
    if err := configManager.SaveConfig("theme", "dark"); err != nil {
        log.Fatalf("Failed to save configuration: %v", err)
    }

    // Load configuration settings
    theme, err := configManager.LoadConfig("theme")
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }
    fmt.Printf("Loaded theme: %s
", theme)

    // Update configuration settings
    if err := configManager.UpdateConfig("theme", "light"); err != nil {
        log.Fatalf("Failed to update configuration: %v", err)
    }
}
