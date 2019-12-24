/*
@Author : Bruce Bee
@Date   : 2019/12/17 10:17
@Email  : mzpy_1119@126.com
*/
package main

import (
    "fmt"
    "flag"

    "./core"
)

func main() {

    var groupID int
    flag.IntVar(&groupID, "g", 0,"组织ID")
    flag.Parse()

    if (groupID == 0){
        fmt.Printf("Usage of ./aliOssAnalyaze:\n  -g int\n        组织ID\n")
        return
    }

     // core.InitDB()
     a := core.InitOSS()
     c := a.ReturnSize()
     fmt.Println(c)
     a.ListFile()
}
