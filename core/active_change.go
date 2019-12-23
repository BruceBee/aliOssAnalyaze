/*
@Author : Bruce Bee
@Date   : 2019/12/23 15:17
@Email  : mzpy_1119@126.com
*/

package core

import (
	"database/sql"
	"fmt"
)

type Active_change struct {
	User int64 `db:user_id`
	Poster string `db:poster`
}

type Activer interface {
	Query() string

}

func (a *Active_change) Query() string {

	d, _ := InitDB()
	dd := QueryActiveChangeById(d,15)
	return dd.Poster
}


func QueryActiveChangeById(DB *sql.DB, id int) Active_change {

	var Active Active_change
	err := DB.QueryRow("SELECT poster FROM jdk_activity_change_log WHERE id = ?", id).Scan(&Active.Poster)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}
	return Active
}

