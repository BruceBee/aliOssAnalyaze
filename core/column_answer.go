/*
@Author : Bruce Bee
@Date   : 2019/12/27 11:02
@Email  : mzpy_1119@126.com
*/
// package core is ...
package core

import (
	"fmt"
	"database/sql"
	"github.com/Unknwon/goconfig"
	"runtime"
	"strings"
)

// QueryColumnAnswer, Gets a list of basic data types
func QueryColumnAnswer(groupID int64) (Q []BaseInfo) {
	db, _ := InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := BaseInfo{
		GrpID: groupID,
		PicBucket: "jdk3t-qiye",
		PicPrefix: "backend_pic/dst/poster/",
		VoiceBucket: "jdk3t-voice",
		VoicePrefix: "backend_voice/",
		VideoBucket: "jdk3t-video",
		VideoPrefix: "video/",
		TableName: filename,
	}

	url , err:= QueryColumnAnswerURL(db, b.GrpID)
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

// QueryColumnAnswerURL, Get the image URL list data through the database query
func QueryColumnAnswerURL(DB *sql.DB, id int64) (banns []string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","banner_info")
	if err != nil {
		panic("panic")
	}

	rows, err := DB.Query(sql, id)
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
