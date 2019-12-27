/*
@Author : Bruce Bee
@Date   : 2019/12/23 17:14
@Email  : mzpy_1119@126.com
*/

// Custom method, mainly through the database to get the URL list
package core

import (
	"database/sql"
	"fmt"
)

// QueryBanner, Gets a list of basic data types
func QueryBanner(groupID int64) (Q []BaseInfo) {
	db, _ := InitDB()
	b := BaseInfo{
		GrpID: groupID,
		PicBucket: "jdk3t-qiye",
		PicPrefix: "backend_pic/dst/poster/",
		TableName: "jdk_banner_info",
	}

	url , err:= QueryBannerURL(db, b.GrpID)
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

// QueryBannerURL, Get the image URL list data through the database query
func QueryBannerURL(DB *sql.DB, id int64) (banns []string, err error) {

	rows, err := DB.Query("SELECT picture_url FROM jdk_banner_info WHERE group_id= ? GROUP BY picture_url;", id)
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

