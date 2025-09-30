// 代码生成时间: 2025-10-01 03:38:23
@author Your Name
@version 1.0
*/

package main

import (
    "fmt"
    "os"
    "path/filepath"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// FileMetadata defines the structure for file metadata
type FileMetadata struct {
    ID        uint   "gorm::"primaryKey""
    FilePath  string
    FileName  string
    FileSize  int64
    CreatedAt time.Time
}

// DB is a global variable for database connection
var DB *gorm.DB

func main() {
    // Connect to the database
    db, err := gorm.Open(sqlite.Open("file_metadata.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to the database: %v", err)
    }
    DB = db
    defer DB.Migrator().DropTable("file_metadata")

    // Initialize the database schema
    err = DB.AutoMigrate(&FileMetadata{})
    if err != nil {
        log.Fatalf("failed to auto migrate: %v", err)
    }

    // Extract metadata from a file
    metadata, err := ExtractMetadata("example.txt")
    if err != nil {
        log.Fatalf("failed to extract metadata: %v", err)
    }

    // Save metadata to the database
    err = SaveMetadata(metadata)
    if err != nil {
        log.Fatalf("failed to save metadata: %v", err)
    }

    fmt.Printf("Metadata extracted and saved for file: %s
", metadata.FileName)
}

// ExtractMetadata extracts metadata from a file
func ExtractMetadata(filePath string) (*FileMetadata, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    info, err := file.Stat()
    if err != nil {
        return nil, err
    }

    metadata := &FileMetadata{
        FilePath:  filePath,
        FileName:  info.Name(),
        FileSize:  info.Size(),
        CreatedAt: info.ModTime(),
    }

    return metadata, nil
}

// SaveMetadata saves metadata to the database
func SaveMetadata(metadata *FileMetadata) error {
    result := DB.Create(metadata)
    if result.Error != nil {
        return result.Error
    }
    return nil
}
