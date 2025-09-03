// 代码生成时间: 2025-09-03 19:03:31
 * and Go best practices are followed for maintainability and extensibility.
 */

package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/disintegration/imaging"
    "gorm.io/driver/sqlite" // Assumes SQLite for simplicity, but can be replaced with other drivers.
    "gorm.io/gorm"
)

// Image represents the image model.
type Image struct {
    gorm.Model
    Path     string
    Width    int
    Height   int
    Format   string
    Thumbnail string
}

// DB is a global variable for the database connection.
var DB *gorm.DB

func main() {
    // Connect to the database.
    var err error
    DB, err = gorm.Open(sqlite.Open("image_resizer.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    defer DB.Migrator().Close()

    // Migrate the schema.
    if err := DB.AutoMigrate(&Image{}); err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    // Scan the directory for images.
    images, err := scanDirectory("path/to/images")
    if err != nil {
        log.Fatalf("failed to scan directory: %v", err)
    }

    // Process images.
    for _, image := range images {
        if err := processImage(image); err != nil {
            log.Printf("failed to process image %q: %v", image.Path, err)
        }
    }
}

// scanDirectory scans the given directory for image files and returns a list of Image structs.
func scanDirectory(directory string) ([]Image, error) {
    var images []Image
    err := filepath.WalkDir(directory, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
