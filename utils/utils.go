package utils

import (
    "fmt"
    "strconv"
)

func FormatSize(size string) (s string) {

    int_size, _ := strconv.ParseFloat(size, 64)
    switch {

    case int_size > 1024*1024*1024 :
        s = fmt.Sprintf("%.2f %s", int_size / (1024 * 1024 * 1024), "GB")
    case int_size > 1024*1024 :
        s = fmt.Sprintf("%.2f %s", int_size / (1024 * 1024), "MB")
    case int_size > 1024 :
        s = fmt.Sprintf("%.2f %s", int_size / 1024 , "KB")
    default:
        s = fmt.Sprintf("%.f %s", int_size, "B")
    }
    return
}
