/*
@Author : Bruce Bee
@Date   : 2019/12/31 10:48
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

// QueryCourseActivity for a list of basic data types
func QueryCourseActivity(groupID int64) (Q []base.BaseInfo) {
	db, _ := db.InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := base.BaseInfo{
		GrpID: groupID,
		PicBucket: "jdk3t-qiye",
		PicPrefix: "backend_pic/dst/poster/",
		TableName: filename,
	}

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")

	url , err:= QueryCourseActivityURL(db, b.GrpID)
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

// QueryCourseActivityURL for the image URL list data through the database query
func QueryCourseActivityURL(DB *sql.DB, id int64) (banns []string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","course_activity")
	if err != nil {
		panic("panic")
	}


	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var (
			poster,
			bannerURL,
			callCover string
		)
		err := rows.Scan(&poster, &bannerURL, &callCover)

		if err != nil {
			fmt.Println(err)
		}else {
			if (poster != ""){
				res := strings.HasPrefix(poster, "https://")
				if res {
					banns = append(banns, poster)
				}
			}

			if (bannerURL != ""){
				res := strings.HasPrefix(bannerURL, "https://")
				if res {
					banns = append(banns, bannerURL)
				}
			}

			if (callCover != ""){
				res := strings.HasPrefix(callCover, "https://")
				if res {
					banns = append(banns, callCover)
				}
			}
		}
	}
	return
}
