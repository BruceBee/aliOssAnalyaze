/*
@Author : Bruce Bee
@Date   : 2019/12/17 10:17
@Email  : mzpy_1119@126.com
*/
package main

import (
    "./core"
    "fmt"
)

func main() {
     // core.InitDB()
     a := core.InitOSS()
     c := a.ReturnSize()
     fmt.Println(c)
     a.ListFile()
}
