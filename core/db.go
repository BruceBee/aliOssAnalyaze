/*
@Author : Bruce Bee
@Date   : 2019/12/17 10:17
@Email  : mzpy_1119@126.com
*/
package core

import (
    "database/sql"
    "fmt"
    "github.com/Unknwon/goconfig"
    _ "github.com/go-sql-driver/mysql"
)

//var DB *sql.DB

func InitDB() (*sql.DB, error){
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
        return nil, err
    }
    fmt.Println("connnect success")
    return DB, nil
}


