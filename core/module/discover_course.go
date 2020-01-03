/*
@Author : Bruce Bee
@Date   : 2020/1/2 10:29
@Email  : mzpy_1119@126.com
*/
// Package module is a core custom method, mainly through the database to get the URL list
package module

import (
	"fmt"
	"database/sql"
	"github.com/Unknwon/goconfig"
	"runtime"
	"strings"
	"../base"
	"../db"
)

// QueryDiscoverCourse for a list of basic data types
func QueryDiscoverCourse(groupID int64) (Q []base.BaseInfo) {
	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := base.BaseInfo{
		GrpID: groupID,
		PicBucket: "jdk3t-qiye",
		PicPrefix: "backend_pic/dst/poster/",
		TableName: filename,
	}

	url , err:= QueryDiscoverCourseURL(mysqlConn, b.GrpID)
	if nil != err {
		fmt.Println("error")
	}

	qiyeOss, _ := cfg.GetValue("oss-cdn-url","qiye_oss")
	for _, u := range url {
		if (u != "") {
			b.PicURL = strings.Replace(u, qiyeOss, "", -1)
			Q = append(Q, b)
		}
	}
	return
}

// QueryDiscoverCourseURL for the image URL list data through the database query
func QueryDiscoverCourseURL(DB *sql.DB, id int64) (banns []string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","discover_course")
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
		if (bann != ""){
			banns = append(banns, bann)
		}
	}
	return
}