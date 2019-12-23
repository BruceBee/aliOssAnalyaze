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

type Banner_nfo struct {
	GrpID int64 `db:group_id`
	PictureUrl string `db:picture_url`
}

type banner interface {
	Query() string

}

func (b *Banner_nfo) Query() string {

	d, _ := InitDB()
	dd , _:= QueryBannerById(d,28)
	fmt.Println(dd)
	return "121212"
}


func QueryBannerById(DB *sql.DB, id int) ([]Banner_nfo, error) {

	var banns []Banner_nfo
	rows, err := DB.Query("SELECT picture_url FROM jdk_banner_info WHERE group_id= ?", id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var bann Banner_nfo
		rows.Scan(&bann.PictureUrl)
		banns = append(banns, bann)
	}
	return banns, nil
}