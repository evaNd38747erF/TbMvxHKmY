// 代码生成时间: 2025-08-26 05:11:01
package main

import (
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
    "net/http"
    "log"
)

// Define a struct to hold the form data for validation.
type FormData struct {
    Name    string `form:"name" json:"name" binding:"required,alphanum"`
    Email   string `form:"email" json:"email" binding:"required,email"`
    Age     int    `form:"age" json:"age" binding:"required,gte=0"`
    Address string `form:"address" json:"address" binding:"omitempty"`
}

func main() {
    r := gin.Default()
    r.POST("/form", func(c *gin.Context) {
        // Create an instance of FormData to hold the data.
        var formData FormData

        // Validate the data from the request.
        if err := c.ShouldBind(&formData); err != nil {
            // Handle validation error.
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }

        // If no error, respond with the received data.
        c.JSON(http.StatusOK, formData)
    })

    // Start the server.
    if err := r.Run(); err != nil {
        log.Fatal(err)
    }
}
