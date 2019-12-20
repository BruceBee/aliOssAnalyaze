package main

import (
    "./core"
)

func main() {
     // core.InitDB()
     a := core.InitOSS()
     //c := a.ReturnSize()
     //fmt.Println(c)
     a.ListFile()
}
