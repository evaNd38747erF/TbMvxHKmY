// 代码生成时间: 2025-08-28 01:18:29
package main

import (
    "archive/zip"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// Decompress unzips a zip file to a destination directory.
func Decompress(src, dest string) error {
    // Open the zip file
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()

    // Create the destination directory if it doesn't exist
    if _, err := os.Stat(dest); os.IsNotExist(err) {
        err := os.MkdirAll(dest, 0755)
        if err != nil {
            return err
        }
    }

    // Iterate through the files in the zip
    for _, f := range r.File {
        // Create the subdirectory structure
        fpath := filepath.Join(dest, f.Name)
        if f.FileInfo().IsDir() {
            // Create directory
            if err := os.MkdirAll(fpath, 0755); err != nil {
                return err
            }
            continue
        }

        // Create the directory structure needed to store the file
        if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
            return err
        }

        // Open the file in the zip for reading
        fr, err := f.Open()
        if err != nil {
            return err
        }
        defer fr.Close()

        // Open the file in the destination directory for writing
        fw, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.FileInfo().Mode())
        if err != nil {
            return err
        }
        defer fw.Close()

        // Copy the contents from the zip file to the destination file
        _, err = io.Copy(fw, fr)
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    src := "example.zip" // The source zip file
    dest := "destination" // The destination directory

    // Decompress the zip file
    if err := Decompress(src, dest); err != nil {
        log.Fatalf("Failed to decompress: %s
", err)
    } else {
        log.Println("Decompression successful.")
    }
}
