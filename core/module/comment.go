/*
@Author : Bruce Bee
@Date   : 2019/12/31 10:47
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

// QueryComment for a list of basic data types
func QueryComment(groupID int64) (Q []base.BaseInfo) {

	sql, _, _, _,voiceBucket,voicePrefix,_,_,_,_,_,_,_ := base.LoadConf("comment")

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

	url , err:= QueryCommentURL(mysqlConn, sql, groupID)
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

// QueryCommentURL for the image URL list data through the database query
func QueryCommentURL(DB *sql.DB, sql string, id int64) (urls []string, err error) {

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var bann string
		err := rows.Scan(&bann)

		if err != nil {
			fmt.Println(err)
		}else {
			if (bann != ""){
				b := strings.Split(bann, "|")
				for _, x := range b {
					urls = append(urls, x)
				}
			}
		}
	}
	return
}

