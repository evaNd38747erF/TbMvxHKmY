// 代码生成时间: 2025-08-27 19:06:58
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/xuri/excelize/v2"
)

// ExcelGenerator 结构体用于封装Excel生成器的配置和操作
type ExcelGenerator struct {
    File *excelize.File
}

// NewExcelGenerator 初始化ExcelGenerator结构体并打开一个新的Excel文件
func NewExcelGenerator() (*ExcelGenerator, error) {
    file, err := excelize.CreateFile()
    if err != nil {
        return nil, err
    }
    return &ExcelGenerator{File: file}, nil
}

// AddSheet 在Excel文件中添加一个新的工作表
func (e *ExcelGenerator) AddSheet(sheetName string) error {
    return e.File.NewSheet(sheetName)
}

// WriteData 向指定工作表写入数据
func (e *ExcelGenerator) WriteData(sheetName string, data [][]string) error {
    for i, row := range data {
        for j, value := range row {
            // 将数据写入单元格
            if err := e.File.SetCellValue(sheetName, fmt.Sprintf("A%d", i+1), value); err != nil {
                return err
    }
        }
    }
    return nil
}

// Save 保存Excel文件到指定路径
func (e *ExcelGenerator) Save(filePath string) error {
    if err := e.File.SaveAs(filePath); err != nil {
        return err
    }
    return nil
}

func main() {
    generator, err := NewExcelGenerator()
    if err != nil {
        log.Fatalf("Failed to create Excel generator: %v", err)
    }
    defer generator.File.Close()

    // 添加工作表
    if err := generator.AddSheet("Sheet1"); err != nil {
        log.Fatalf("Failed to add sheet: %v", err)
    }

    // 准备数据
    data := [][]string{
        {"Header1", "Header2", "Header3"},
        {"Data1", "Data2", "Data3"},
        // ... 更多数据行
    }

    // 写入数据
    if err := generator.WriteData("Sheet1", data); err != nil {
        log.Fatalf("Failed to write data: %v", err)
    }

    // 保存Excel文件
    currentDate := time.Now().Format("2006-01-02")
    filePath := fmt.Sprintf("./output/excel_%s.xlsx", currentDate)
    if err := generator.Save(filePath); err != nil {
        log.Fatalf("Failed to save Excel file: %v", err)
    }

    fmt.Printf("Excel file saved at: %s
", filePath)
}
