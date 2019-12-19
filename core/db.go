package core

import (
    "fmt"
    "database/sql"
    "encoding/json"
    _ "github.com/go-sql-driver/mysql"
    "github.com/pkg/errors"
    "github.com/Unknwon/goconfig"
)

var DB *sql.DB

type City struct {
    id 	int64
    city string
    url string
}

func InitDB() {
    cfg, err := goconfig.LoadConfigFile("conf/app.ini")
    if err != nil {
        panic("panic")
    }

    host, err := cfg.GetValue("mysql","host")
    user, err := cfg.GetValue("mysql","user")
    pwd, err := cfg.GetValue("mysql","pwd")
    port, err := cfg.Int("mysql","port")
    db, err := cfg.GetValue("mysql","db")

    url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		user,
		pwd,
		host,
		port,
		db,
		"utf8")

    DB, _ := sql.Open("mysql", url)
    if err := DB.Ping(); err != nil {
        fmt.Println("open database fail")
        return
    }
    fmt.Println("connnect success")
    fmt.Printf("hello world")
    return
}

func Query() {
    var city City
    rows , err := DB.Query("SELECT * FROM city")
    if err == nil {
        errors.New("QUERY INCUR ERROR")
    }
    for rows.Next() {
        e := rows.Scan(city.id, city.city, city.url)
        if e != nil {
            fmt.Println(json.Marshal(city))
        }
    }
    rows.Close()
    defer DB.Close()
}

