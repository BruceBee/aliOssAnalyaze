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
func QueryBanner(group_id int64) (*BaseInfo, []string) {
	db, _ := InitDB()
	b := BaseInfo{
		GrpID: group_id,
		PictureBucket: "jdk3t-qiye",
		PicturePrefix: "backend_pic/dst/poster/",
		TableName: "jdk_banner_info",
	}

	url , err:= QueryBannerURL(db, b.GrpID)
	if nil != err {
		fmt.Println("error")
	}
	return &b, url
}

// QueryBannerURL ...
func QueryBannerURL(DB *sql.DB, id int64) ([]string, error) {

	var banns []string
	rows, err := DB.Query("SELECT picture_url FROM jdk_banner_info WHERE group_id= ?", id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var bann string
		rows.Scan(&bann)
		banns = append(banns, bann)
	}
	return banns, nil
}


