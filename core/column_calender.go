/*
@Author : Bruce Bee
@Date   : 2019/12/30 16:16
@Email  : mzpy_1119@126.com
*/
// package core is ...
package core

import (
	"fmt"
	"database/sql"
	"github.com/Unknwon/goconfig"
	"runtime"
	"strings"
)

// QueryColumnCalender , Gets a list of basic data types
func QueryColumnCalender(groupID int64) (Q []BaseInfo) {

	db, _ := InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := BaseInfo{
		GrpID: groupID,
		PicBucket: "jdk3t-qiye",
		PicPrefix: "backend_pic/dst/poster/",
		TableName: filename,
	}
	url , err:= QueryColumnCalenderURL(db, b.GrpID)
	if nil != err {
		fmt.Println("error")
	}

	for _, u := range url {
		if (u != "") {
			b.PicURL = u
			Q = append(Q, b)
		}
	}
	return
}

// QueryColumnCalenderURL for the image URL list data through the database query
func QueryColumnCalenderURL(DB *sql.DB, id int64) (banns []string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","column_calendar")
	if err != nil {
		panic("panic")
	}

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var bann string
		rows.Scan(&bann)
		banns = append(banns, bann)
	}
	return
}


