package main

import (
    "fmt"
    "./core"
)

func main() {
     // core.InitDB()
     a := core.InitOSS()
     c := a.ReturnSize()
     fmt.Println(c)
}
