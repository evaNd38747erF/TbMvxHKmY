// 代码生成时间: 2025-09-12 11:41:32
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "sort"
    "strings"
)

// 文件重命名结构体
type FileRename struct {
    Source string
    Target string
}

// 文件批量重命名函数
func bulkRename(files []FileRename) error {
    for _, file := range files {
        sourceFile := file.Source
        targetFile := file.Target

        // 检查源文件是否存在
        if _, err := os.Stat(sourceFile); os.IsNotExist(err) {
            return fmt.Errorf("source file %s does not exist", sourceFile)
        }

        // 检查目标文件是否已存在
        if _, err := os.Stat(targetFile); err == nil {
            return fmt.Errorf("target file %s already exists", targetFile)
        }

        // 重命名文件
        if err := os.Rename(sourceFile, targetFile); err != nil {
            return fmt.Errorf("failed to rename %s to %s: %v", sourceFile, targetFile, err)
        }
    }
    return nil
}

// 生成新文件名的函数
func generateNewFileNames(files []string) []FileRename {
    var renamedFiles []FileRename
    for i, file := range files {
        newFileName := fmt.Sprintf("%d_%s", i+1, filepath.Base(file))
        renamedFiles = append(renamedFiles, FileRename{Source: file, Target: filepath.Dir(file) + "/" + newFileName})
    }
    return renamedFiles
}

// 获取指定目录下所有文件的函数
func getAllFilesInDir(dir string) ([]string, error) {
    files := []string{}
    err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    // 按文件名排序
    sort.Strings(files)
    return files, nil
}

func main() {
    // 指定目录路径
    dirPath := "./example"

    // 获取目录下所有文件
    files, err := getAllFilesInDir(dirPath)
    if err != nil {
        log.Fatalf("failed to get files: %v", err)
    }

    // 生成新文件名
    renamedFiles := generateNewFileNames(files)

    // 执行批量重命名
    if err := bulkRename(renamedFiles); err != nil {
        log.Fatalf("failed to rename files: %v", err)
    }

    fmt.Println("Files have been renamed successfully.")
}
