/*
@Author : Bruce Bee
@Date   : 2019/12/31 10:47
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

// QueryCourse for a list of basic data types
func QueryCourse(groupID int64) (Q []base.BaseInfo) {
	db, _ := db.InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := base.BaseInfo{
		GrpID: groupID,
		PicBucket: "jdk3t-qiye",
		PicPrefix: "",
		TableName: filename,
	}

	url , err:= QueryCourseURL(db, b.GrpID)
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
func QueryCourseURL(DB *sql.DB, id int64) (banns []string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","course")
	if err != nil {
		panic("panic")
	}

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
					banns = append(banns, userImg[1:])
				}else {
					banns = append(banns, userImg)
				}

			}

			if (bannerImg != ""){
				banns = append(banns, "backend_pic/dst/poster/"+bannerImg)
			}
		}
	}
	return
}


