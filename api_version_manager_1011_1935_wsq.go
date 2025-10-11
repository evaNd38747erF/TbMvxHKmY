// 代码生成时间: 2025-10-11 19:35:33
package main

import (
    "fmt"
    "log"
    "net/http"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// VersionRecord represents the structure of the API versions in the database
type VersionRecord struct {
    gorm.Model
    Version string `gorm:"primaryKey"`
}

// APIVersionManager is the struct that represents the API version manager
type APIVersionManager struct {
    DB *gorm.DB
}

// NewAPIVersionManager creates a new instance of the API version manager
func NewAPIVersionManager() *APIVersionManager {
    // Initialize the database connection
    db, err := gorm.Open(sqlite.Open("api_versions.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&VersionRecord{})

    return &APIVersionManager{DB: db}
}

// AddVersion adds a new API version to the database
func (m *APIVersionManager) AddVersion(version string) error {
    // Check if the version already exists
    var existingRecord VersionRecord
    if result := m.DB.Where("version = ?", version).First(&existingRecord); result.Error == nil {
        return fmt.Errorf("version already exists")
    }

    // Add a new version record
    newRecord := VersionRecord{Version: version}
    if result := m.DB.Create(&newRecord); result.Error != nil {
        return fmt.Errorf("failed to add version: %w", result.Error)
    }
    return nil
}

// GetVersions retrieves all API versions from the database
func (m *APIVersionManager) GetVersions() ([]VersionRecord, error) {
    var versions []VersionRecord
    if result := m.DB.Find(&versions); result.Error != nil {
        return nil, fmt.Errorf("failed to retrieve versions: %w", result.Error)
    }
    return versions, nil
}

// StartServer starts the HTTP server for managing API versions
func (m *APIVersionManager) StartServer(port string) {
    http.HandleFunc("/add-version", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "only POST method is allowed", http.StatusMethodNotAllowed)
            return
        }

        var version string
        if err := r.ParseForm(); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if version = r.FormValue("version"); version == "" {
            http.Error(w, "version parameter is required", http.StatusBadRequest)
            return
        }

        if err := m.AddVersion(version); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Version added successfully")
    })

    http.HandleFunc("/versions", func(w http.ResponseWriter, r *http.Request) {
        versions, err := m.GetVersions()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(versions); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    })

    log.Printf("API version manager server is running on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal("failed to start server: ", err)
    }
}

func main() {
    manager := NewAPIVersionManager()
    manager.StartServer("8080")
}