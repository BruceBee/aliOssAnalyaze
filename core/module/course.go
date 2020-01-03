/*
@Author : Bruce Bee
@Date   : 2019/12/31 10:47
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

// QueryCourse for a list of basic data types
func QueryCourse(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, _, _,_,_,_,_,_,_,_,_,_ := base.LoadConf("course")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := base.BaseInfo{
		GrpID: groupID,
		PicBucket: picBucket,
		PicPrefix: "",
		TableName: filename,
	}

	url , err:= QueryCourseURL(mysqlConn, sql, groupID)
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

// QueryCourseURL for the image URL list data through the database query
func QueryCourseURL(DB *sql.DB, sql string, id int64) (urls []string, err error) {

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var (
			userImg,
			bannerImg string
		)
		err := rows.Scan(&userImg, &bannerImg)

		if err != nil {
			fmt.Println(err)
		}else {
			if (userImg != ""){
				res := strings.HasPrefix(userImg, "/")
				if res {
					urls = append(urls, userImg[1:])
				}else {
					urls = append(urls, userImg)
				}

			}

			if (bannerImg != ""){
				urls = append(urls, "backend_pic/dst/poster/"+bannerImg)
			}
		}
	}
	return
}


