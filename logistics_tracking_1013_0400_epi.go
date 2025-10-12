// 代码生成时间: 2025-10-13 04:00:26
package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// TrackingRecord represents a single logistics tracking record
type TrackingRecord struct {
    gorm.Model
    PackageID  string `gorm:"primaryKey"` // Unique identifier for the package
    Status     string `gorm:"index"`      // Current status of the package
    LastUpdate string // Timestamp of the last update
}

// TrackingService is responsible for managing tracking records
type TrackingService struct {
    db *gorm.DB
}

// NewTrackingService initializes a new TrackingService with a database connection
func NewTrackingService(db *gorm.DB) *TrackingService {
    return &TrackingService{db: db}
}

// AddRecord adds a new tracking record to the database
func (s *TrackingService) AddRecord(record *TrackingRecord) error {
    if err := s.db.Create(record).Error; err != nil {
        return err // Return error if record cannot be created
    }
    return nil
}

// UpdateRecord updates an existing tracking record
func (s *TrackingService) UpdateRecord(id uint, status string) error {
    var record TrackingRecord
    if err := s.db.First(&record, id).Error; err != nil {
        return err // Return error if record not found
    }
    record.Status = status
    if err := s.db.Save(&record).Error; err != nil {
        return err // Return error if record cannot be saved
    }
    return nil
}

// GetRecord retrieves a tracking record by its ID
func (s *TrackingService) GetRecord(id uint) (TrackingRecord, error) {
    var record TrackingRecord
    if err := s.db.First(&record, id).Error; err != nil {
        return record, err // Return record and error if not found
    }
    return record, nil
}

func main() {
    // Initialize the database connection
    db, err := gorm.Open(sqlite.Open("logistics.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    defer db.Close()

    // Migrate the schema
    if err := db.AutoMigrate(&TrackingRecord{}); err != nil {
        log.Fatal("Failed to migrate database: ", err)
    }

    // Initialize the tracking service
    trackingService := NewTrackingService(db)

    // Sample usage of the service
    if err := trackingService.AddRecord(&TrackingRecord{
        PackageID:  "PKG123",
        Status:     "In Transit",
        LastUpdate: "2023-04-01T12:00:00Z",
    }); err != nil {
        log.Fatal("Failed to add tracking record: ", err)
    }

    // Update tracking record
    if err := trackingService.UpdateRecord(1, "Delivered"); err != nil {
        log.Fatal("Failed to update tracking record: ", err)
    }

    // Retrieve tracking record
    record, err := trackingService.GetRecord(1)
    if err != nil {
        log.Fatal("Failed to retrieve tracking record: ", err)
    }
    fmt.Printf("Tracking Record: %+v
", record)
}
