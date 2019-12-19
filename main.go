package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    DB, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
    if err := DB.Ping(); err != nil {
        fmt.Println("open database fail")
        return
    }
    fmt.Println("connnect success")
    fmt.Printf("hello world")
}
