// 代码生成时间: 2025-09-21 06:11:51
Features:
1. Clear code structure for easy understanding.
2. Includes error handling.
3. Contains necessary comments and documentation.
4. Follows GoLang best practices.
5. Ensures code maintainability and extensibility.
*/

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// Folder represents a folder with its name and children
type Folder struct {
    Name string
    Children []*Folder
}

// NewFolder creates a new Folder instance
func NewFolder(name string) *Folder {
    return &Folder{Name: name, Children: []*Folder{}}
}

// AddChild adds a child folder to the current folder
func (f *Folder) AddChild(child *Folder) {
    f.Children = append(f.Children, child)
}

// OrganizeFolderStructure recursively organizes the folder structure
func (f *Folder) OrganizeFolderStructure(rootPath string) error {
    files, err := os.ReadDir(rootPath)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }
    for _, file := range files {
        // Skip directories
        if file.IsDir() {
            child := NewFolder(file.Name())
            f.AddChild(child)
            if err := child.OrganizeFolderStructure(filepath.Join(rootPath, file.Name())); err != nil {
                return fmt.Errorf("failed to organize subdirectory: %w", err)
            }
        }
    }
    return nil
}

// PrintStructure prints the folder structure in a readable format
func (f *Folder) PrintStructure(indent string) {
    if f.Name != "" {
        fmt.Println(indent + f.Name)
    }
    for _, child := range f.Children {
        child.PrintStructure(indent + "  ")
    }
}

func main() {
    rootFolder := NewFolder("Root")
    if err := rootFolder.OrganizeFolderStructure("."); err != nil {
        log.Fatalf("failed to organize folder structure: %s", err)
    }
    rootFolder.PrintStructure("")
}