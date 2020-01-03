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
	"runtime"
	"strings"
	"../base"
	"../db"
)

// QueryDiscoverCourse for a list of basic data types
func QueryDiscoverCourse(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, picPrefix, picUrl,_,_,_,_,_,_,_,_,_ := base.LoadConf("discover_course")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := base.BaseInfo{
		GrpID: groupID,
		PicBucket: picBucket,
		PicPrefix: picPrefix,
		TableName: filename,
	}

	url , err:= QueryDiscoverCourseURL(mysqlConn, sql, groupID, picUrl+picPrefix)
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

// QueryDiscoverCourseURL for the image URL list data through the database query
func QueryDiscoverCourseURL(DB *sql.DB, sql string, id int64, prefix string) (urls []string, err error) {
	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var url string
		rows.Scan(&url)
		if (url != ""){
			urls = append(urls, strings.Replace(url, prefix, "", -1))
		}
	}
	return
}