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

// BannerNfo ...
type BannerInfo struct {
	GrpID int64 `db:group_id`
	PictureURL string `db:picture_url`
}

type banner interface {
	Query() string

}

// Query ...
func (b *BannerInfo) Query() string {

	d, _ := InitDB()
	dd , _:= QueryBannerByID(d,28)
	fmt.Println(dd)
	return "121212"
}

// QueryBannerByID ...
func QueryBannerByID(DB *sql.DB, id int) ([]BannerInfo, error) {

	var banns []BannerInfo
	rows, err := DB.Query("SELECT picture_url FROM jdk_banner_info WHERE group_id= ?", id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var bann BannerInfo
		rows.Scan(&bann.PictureURL)
		banns = append(banns, bann)
	}
	return banns, nil
}