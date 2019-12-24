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

// ActiveChange ...
type ActiveChange struct {
	User int64 `db:user_id`
	Poster string `db:poster`
}

// Activer ...
type Activer interface {
	Query() string

}

// Query ...
func (a *ActiveChange) Query() string {

	d, _ := InitDB()
	dd := QueryActiveChangeByID(d,15)
	return dd.Poster
}

// QueryActiveChangeByID ...
func QueryActiveChangeByID(DB *sql.DB, id int) ActiveChange {

	var Active ActiveChange
	err := DB.QueryRow("SELECT poster FROM jdk_activity_change_log WHERE id = ?", id).Scan(&Active.Poster)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}
	return Active
}

