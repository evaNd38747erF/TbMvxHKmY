// 代码生成时间: 2025-09-02 00:12:45
package main

import (
    "image"
    "image/jpeg"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// ImageResizer 结构体，用于存储图片尺寸调整的相关参数
type ImageResizer struct {
    OutputPath string // 输出图片的目录
    NewWidth   int    // 新图片的宽度
    NewHeight  int    // 新图片的高度
}

// NewImageResizer 初始化并返回一个新的ImageResizer实例
func NewImageResizer(outputPath string, newWidth, newHeight int) *ImageResizer {
    return &ImageResizer{
        OutputPath: outputPath,
        NewWidth:   newWidth,
        NewHeight:  newHeight,
    }
}

// ResizeAll 递归遍历目录，调整所有JPEG图片的尺寸
func (r *ImageResizer) ResizeAll(directory string) error {
    files, err := ioutil.ReadDir(directory)
    if err != nil {
        return err
    }
    for _, file := range files {
        filePath := filepath.Join(directory, file.Name())
        if file.IsDir() {
            if err := r.ResizeAll(filePath); err != nil {
                return err
            }
        } else if strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".jpeg") {
            if err := r.resize(filePath); err != nil {
                return err
            }
        }
    }
    return nil
}

// resize 调整单个图片文件的尺寸
func (r *ImageResizer) resize(filePath string) error {
    src, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer src.Close()

    img, _, err := image.Decode(src)
    if err != nil {
        return err
    }

    dst := image.NewRGBA(img.Bounds())
    imgRect := img.Bounds()
    width, height := float64(r.NewWidth), float64(r.NewHeight)

    // 计算缩放比例
    scaleWidth := width / float64(imgRect.Dx())
    scaleHeight := height / float64(imgRect.Dy())
    scale := scaleWidth
    if scaleHeight < scaleWidth {
        scale = scaleHeight
    }

    newWidth, newHeight := int(float64(imgRect.Dx())*scale), int(float64(imgRect.Dy())*scale)
    dstRect := image.Rect(0, 0, newWidth, newHeight)

    // 缩放图片
    ia := &image.RGBA{Rect: dstRect, Stride: 4 * newWidth, Pix: make([]uint8, 4*newWidth*newHeight)}
    g := &image.RGBA{Rect: imgRect, Stride: 4 * imgRect.Dx(), Pix: img.(*image.RGBA).Pix}
    draw.Draw(ia, dstRect, g, image.ZP, draw.Src)

    // 保存调整尺寸后的图片
    dstPath := filepath.Join(r.OutputPath, file.Name())
    dstFile, err := os.Create(dstPath)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    if err := jpeg.Encode(dstFile, ia, nil); err != nil {
        return err
    }

    log.Printf("Resized image: %s
", filePath)
    return nil
}

func main() {
    // 使用示例
    resizer := NewImageResizer("./output", 800, 600)
    if err := resizer.ResizeAll("./input"); err != nil {
        log.Fatal(err)
    }
}