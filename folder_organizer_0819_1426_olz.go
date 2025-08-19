// 代码生成时间: 2025-08-19 14:26:38
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// Folder represents a folder with its name, path, and list of subfolders
type Folder struct {
    Name     string
    Path     string
    SubFiles []string
}

func main() {
    // Define the root folder path
    rootPath := "/path/to/your/folder"
    
    // Create a Folder instance for root
    rootFolder := Folder{Name: "Root", Path: rootPath}
    
    // Organize the folder structure
    if err := OrganizeFolder(rootFolder); err != nil {
        log.Fatalf("Failed to organize folder: %v", err)
    }
    fmt.Println("Folder organization completed successfully.")
}

// OrganizeFolder recursively organizes the folder structure
func OrganizeFolder(folder Folder) error {
    // Read the directory
    files, err := os.ReadDir(folder.Path)
    if err != nil {
        return err
    }
    
    for _, file := range files {
        fullPath := filepath.Join(folder.Path, file.Name())
        if file.IsDir() {
            // If it's a directory, create a new Folder and call OrganizeFolder recursively
            subFolder := Folder{Name: file.Name(), Path: fullPath}
            folder.SubFiles = append(folder.SubFiles, subFolder.Name)
            if err := OrganizeFolder(subFolder); err != nil {
                return err
            }
        } else {
            // If it's a file, add it to the SubFiles list
            folder.SubFiles = append(folder.SubFiles, fullPath)
        }
    }
    return nil
}
