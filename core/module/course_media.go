/*
@Author : Bruce Bee
@Date   : 2020/1/2 10:29
@Email  : mzpy_1119@126.com
*/

// Package module is a core custom method, mainly through the database to get the URL list
package module

import (
	"../base"
	"../db"
	"database/sql"
	"fmt"
	"runtime"
	"strings"
)

// QueryCourseMedia for a list of basic data types
func QueryCourseMedia(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, picPrefix, _,_,_,_,_,_,_,_,_,_ := base.LoadConf("course_media")

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

	url , err:= QueryCourseMediaURL(mysqlConn, sql, groupID)
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

// QueryCourseMediaURL for the image URL list data through the database query
func QueryCourseMediaURL(DB *sql.DB, sql string, id int64) (urls []string, err error) {

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var (
			miniPic ,
			pic string
		)
		rows.Scan(&miniPic, &pic)

		if(miniPic != ""){
			urls = append(urls, miniPic)
		}

		if(pic != ""){
			urls = append(urls, pic)
		}
	}
	return
}


