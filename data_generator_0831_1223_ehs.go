// 代码生成时间: 2025-08-31 12:23:04
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// User represents a user model
type User struct {
    ID       uint   "json:"id" gorm:"primaryKey""
data
    Name     string "json:"name" gorm:"type:varchar(100)"
    Email    string "json:"email" gorm:"type:varchar(100)"
    Age      int    "json:"age"
}

// GenerateData generates a specified number of user records
func GenerateData(count int) ([]User, error) {
    users := make([]User, count)
    rand.Seed(time.Now().UnixNano())

    for i := range users {
        users[i].Name = fmt.Sprintf("User%d", i+1)
        users[i].Email = fmt.Sprintf("user%d@example.com", i+1)
        users[i].Age = rand.Intn(100)
    }

    return users, nil
}

func main() {
    count := 10 // Define the number of users to generate
    users, err := GenerateData(count)
    if err != nil {
        fmt.Printf("Error generating data: %v
", err)
        return
    }

    // Print out the generated user data
    for _, user := range users {
        fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d
", user.ID, user.Name, user.Email, user.Age)
    }
}
