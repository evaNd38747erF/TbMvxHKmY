// 代码生成时间: 2025-09-14 01:32:25
package main

import (
    "fmt"
    "net/url"
    "strings"
)

// IsValidURL checks if the given string is a valid URL
func IsValidURL(toTest string) (bool, error) {
    // Parse the URL to determine if it is valid
    parsedURL, err := url.ParseRequestURI(toTest)
    if err != nil {
        return false, err
    }

    // Check if the scheme is valid (http, https)
    if !(strings.HasPrefix(parsedURL.Scheme, "http") && (len(parsedURL.Scheme) == 4 || len(parsedURL.Scheme) == 5)) {
        return false, fmt.Errorf("invalid scheme: %s", parsedURL.Scheme)
    }

    // Check if both host and path are present
    if parsedURL.Host == "" || parsedURL.Path == "" {
        return false, fmt.Errorf("URL must have both a host and a path")
    }

    return true, nil
}

func main() {
    // Test URLs
    testURLs := []string{
        "https://example.com",
        "http://example.com/path",
        "invalid-url",
        "ftp://example.com",
    }

    for _, testURL := range testURLs {
        isvalid, err := IsValidURL(testURL)
        if err != nil {
            fmt.Printf("Error validating URL '%s': %s
", testURL, err)
        } else if isvalid {
            fmt.Printf("URL '%s' is valid.
", testURL)
        } else {
            fmt.Printf("URL '%s' is invalid.
", testURL)
        }
    }
}