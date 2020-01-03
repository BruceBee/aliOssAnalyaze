/*
@Author : Bruce Bee
@Date   : 2019/12/31 10:48
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

// QueryCourseActivity for a list of basic data types
func QueryCourseActivity(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, picPrefix, picUrl,_,_,_,_,_,_,_,_,_ := base.LoadConf("course_activity")
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

	url , err:= QueryCourseActivityURL(mysqlConn, sql, groupID, picUrl+picPrefix)
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

// QueryCourseActivityURL for the image URL list data through the database query
func QueryCourseActivityURL(DB *sql.DB, sql string, id int64, prefix string) (urls []string, err error) {

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
				res := strings.HasPrefix(poster, prefix)
				if res {
					urls = append(urls, strings.Replace(poster, prefix, "", -1))
				}
			}

			if (bannerURL != ""){
				res := strings.HasPrefix(bannerURL, prefix)
				if res {
					urls = append(urls, strings.Replace(bannerURL, prefix, "", -1))
				}
			}

			if (callCover != ""){
				res := strings.HasPrefix(callCover, prefix)
				if res {
					urls = append(urls, strings.Replace(callCover, prefix, "", -1))
				}
			}
		}
	}
	return
}
