/*
@Author : Bruce Bee
@Date   : 2019/12/17 10:17
@Email  : mzpy_1119@126.com
*/

// Package utils

package utils

import (
    "fmt"
    "strconv"
)

// FormatSize is oss file size Format
func FormatSize(size string) (s string) {

    intSize, _ := strconv.ParseFloat(size, 64)
    switch {

    case intSize > 1024*1024*1024 :
        s = fmt.Sprintf("%.2f %s", intSize / (1024 * 1024 * 1024), "GB")
    case intSize > 1024*1024 :
        s = fmt.Sprintf("%.2f %s", intSize / (1024 * 1024), "MB")
    case intSize > 1024 :
        s = fmt.Sprintf("%.2f %s", intSize / 1024 , "KB")
    default:
        s = fmt.Sprintf("%.f %s", intSize, "B")
    }
    return
}
