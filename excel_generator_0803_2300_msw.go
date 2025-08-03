// 代码生成时间: 2025-08-03 23:00:33
package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/tealeg/xlsx/v3"
    "github.com/jinzhu/gorm"
)

// ExcelGenerator 包含生成Excel所需的数据
type ExcelGenerator struct {
    Data [][]string
    SheetName string
    FileName string
}

// NewExcelGenerator 创建一个新的ExcelGenerator实例
func NewExcelGenerator(data [][]string, sheetName, fileName string) *ExcelGenerator {
    return &ExcelGenerator{
        Data: data,
        SheetName: sheetName,
        FileName: fileName,
    }
}

// GenerateExcel 生成Excel文件
func (e *ExcelGenerator) GenerateExcel() error {
    file := xlsx.NewFile()
    // 创建一个新的工作表
    sheet, err := file.AddSheet(e.SheetName)
    if err != nil {
        return err
    }

    for _, row := range e.Data {
        for i, cell := range row {
            sheet.AddRow()
            sheet.Rows[i].AddCell().Value = cell
        }
    }

    // 将文件写入到磁盘
    f, err := os.Create(e.FileName)
    if err != nil {
        return err
    }
    defer f.Close()

    if err := file.Save(f); err != nil {
        return err
    }

    return nil
}

func main() {
    // 示例数据
    data := [][]string{
        {"", "Product", "Quantity", "Price"},
        {"1", "Apple", "10", "0.99"},
        {"2", "Banana", "20", "0.39"},
        {"3", "Cherry", "30", "0.29"},
    }

    // 创建ExcelGenerator实例
    generator := NewExcelGenerator(data, "Product List", "products.xlsx")

    // 生成Excel文件
    if err := generator.GenerateExcel(); err != nil {
        log.Fatalf("Failed to generate Excel file: %s", err)
    } else {
        fmt.Println("Excel file generated successfully.")
    }
}
