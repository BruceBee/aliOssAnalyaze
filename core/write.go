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
	file, _ := os.OpenFile(filePath + filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0664)

	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(text)

	// clear cache and written to disk
	writer.Flush()
}

