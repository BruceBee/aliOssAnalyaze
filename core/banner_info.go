/*
@Author : Bruce Bee
@Date   : 2019/12/23 17:14
@Email  : mzpy_1119@126.com
*/

package core

import (
	"database/sql"
	"fmt"
)

// QueryBanner ...
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

// QueryBannerURL ...
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


