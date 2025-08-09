// 代码生成时间: 2025-08-10 07:53:44
package main

import (
    "context"
    "fmt"
    "io/fs"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "time"

    "github.com/disintegration/imaging"
)

// ImageResizer is the main application struct
type ImageResizer struct {
    // DirectoryPath is the path to the directory containing images to resize
    DirectoryPath string
    // OutputPath is the path to save resized images
    OutputPath string
    // Width is the desired width of the resized images
    Width int
    // Height is the desired height of the resized images
    Height int
}

// NewImageResizer creates a new instance of ImageResizer
func NewImageResizer(directoryPath, outputPath string, width, height int) *ImageResizer {
    return &ImageResizer{
        DirectoryPath: directoryPath,
        OutputPath: outputPath,
        Width: width,
        Height: height,
    }
}

// ResizeImage resizes an image and saves it to the output path
func (r *ImageResizer) ResizeImage(ctx context.Context, imagePath string) error {
    file, err := os.Open(imagePath)
    if err != nil {
        return fmt.Errorf("failed to open image file: %w", err)
    }
    defer file.Close()

    img, err := imaging.Decode(file)
    if err != nil {
        return fmt.Errorf("failed to decode image: %w", err)
    }

    resizedImg := imaging.Resize(img, r.Width, r.Height, imaging.Lanczos)
    newImagePath := filepath.Join(r.OutputPath, filepath.Base(imagePath))

    err = os.MkdirAll(r.OutputPath, os.ModePerm)
    if err != nil {
        return fmt.Errorf("failed to create output directory: %w", err)
    }

    outFile, err := os.Create(newImagePath)
    if err != nil {
        return fmt.Errorf("failed to create new image file: %w", err)
    }
    defer outFile.Close()

    if err := imaging.EncodeJPEG(outFile, resizedImg, imaging.JPEGQuality(50)); err != nil {
        return fmt.Errorf("failed to encode and save resized image: %w", err)
    }

    return nil
}

// ResizeAllImages resizes all images in the directory and saves them to the output path
func (r *ImageResizer) ResizeAllImages(ctx context.Context) error {
    err := filepath.WalkDir(r.DirectoryPath, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() && (filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".png") {
            if err := r.ResizeImage(ctx, path); err != nil {
                log.Printf("failed to resize image %s: %v", path, err)
            }
        }
        return nil
    })
    return err
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel()

    resizer := NewImageResizer("/path/to/input/images", "/path/to/output/images", 800, 600)
    if err := resizer.ResizeAllImages(ctx); err != nil {
        log.Fatalf("failed to resize images: %v", err)
    }

    fmt.Println("Image resizing completed.")
}