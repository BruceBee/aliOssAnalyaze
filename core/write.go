/*
@Author : Bruce Bee
@Date   : 2019/12/24 11:17
@Email  : mzpy_1119@126.com
*/

package core

import (
	"bufio"
	"os"
)

const filePath = "./data/"

var (
	newFile *os.File
	err error
)

// CreateFile ...
func CreateFile(filename string, text string)  {
	file, _ := os.OpenFile(filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0664)

	defer file.Close()

	// 获取bufio.Writer实例
	writer := bufio.NewWriter(file)

	// 写入字符串
	_, err = writer.WriteString(text)
	// 清空缓存 确保写入磁盘
	writer.Flush()
}

