// 代码生成时间: 2025-09-12 00:10:56
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// FolderOrganizer is a structure to hold configuration for folder organizing
type FolderOrganizer struct {
    SourceFolder string
    TargetFolder string
}

// NewFolderOrganizer creates a new FolderOrganizer instance
func NewFolderOrganizer(source, target string) *FolderOrganizer {
    return &FolderOrganizer{
        SourceFolder: source,
        TargetFolder: target,
    }
}

// OrganizeFolders goes through the source folder and organizes files into the target folder based on file extension
func (fo *FolderOrganizer) OrganizeFolders() error {
    // Check if source folder exists
    if _, err := os.Stat(fo.SourceFolder); os.IsNotExist(err) {
        return fmt.Errorf("source folder does not exist: %w", err)
    }

    // Iterate through the files in the source folder
    err := filepath.WalkDir(fo.SourceFolder, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }

        if d.IsDir() {
            return nil // Ignore directories
        }

        // Extract file extension
        extension := strings.TrimPrefix(filepath.Ext(path), ".")
        if extension == "" {
            return nil // Ignore files without extension
        }

        // Construct target path based on file extension
        targetPath := filepath.Join(fo.TargetFolder, extension)

        // Create target directory if it does not exist
        if _, err := os.Stat(targetPath); os.IsNotExist(err) {
            if err := os.MkdirAll(targetPath, 0755); err != nil {
                return fmt.Errorf("failed to create target directory: %w", err)
            }
        }

        // Move file to target path
        destination := filepath.Join(targetPath, d.Name())
        if err := os.Rename(path, destination); err != nil {
            return fmt.Errorf("failed to move file: %w", err)
        }

        fmt.Printf("Moved %s to %s
", path, destination)
        return nil
    })

    return err
}

func main() {
    source := "./source"
    target := "./target"

    fo := NewFolderOrganizer(source, target)
    err := fo.OrganizeFolders()
    if err != nil {
        log.Fatalf("Failed to organize folders: %s
", err)
    }
    fmt.Println("Folder organization completed successfully.")
}
