package utils

import (
    "fmt"
    "os"
)

func ReadFile(path string) string {
    content, err := os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
    }
    return string(content)
}
