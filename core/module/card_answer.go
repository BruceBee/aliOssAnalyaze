/*
@Author : Bruce Bee
@Date   : 2019/12/24 15:30
@Email  : mzpy_1119@126.com
*/

// Package core is a core custom method, mainly through the database to get the URL list
package module

import (
	"../base"
	"../db"
	"database/sql"
	"fmt"
	"runtime"
	"strings"
)

// QueryCard , Gets a list of basic data types
func QueryCard(groupID int64) (Q []base.BaseInfo) {

	sql, _, _, _,voiceBucket,voicePrefix,_,_,_,_,_,_,_ := base.LoadConf("card_answer")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()

	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	b := base.BaseInfo{
		GrpID: groupID,
		VoiceBucket: voiceBucket,
		VoicePrefix: voicePrefix,
		TableName: filename,
	}
	url , err:= QueryCardAnswerURL(mysqlConn, sql, groupID)
	if nil != err {
		fmt.Println("error")
	}

	for _, u := range url {
		if (u != "") {
			b.VoiceURL = u
			Q = append(Q, b)
		}
	}
	return
}

// QueryCardAnswerURL for the image URL list data through the database query
func QueryCardAnswerURL(DB *sql.DB, sql string, id int64) (urls []string, err error) {

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var url string
		rows.Scan(&url)
		urls = append(urls, url)
	}
	return
}

