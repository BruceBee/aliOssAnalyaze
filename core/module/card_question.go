/*
@Author : Bruce Bee
@Date   : 2019/12/27 09:49
@Email  : mzpy_1119@126.com
*/

// Package core is a core custom method, mainly through the database to get the URL list
package module

import (
	"fmt"
	"runtime"
	"strings"
	"database/sql"
	"../base"
	"../db"
	"../../utils"
)

// QueryCardQuestion is get a list of basic data types
func QueryCardQuestion(groupID int64) (Q []base.BaseInfo) {
	sql, _, _, _,voiceBucket,voicePrefix,voiceUrl,_,_,_,_,_,_ := base.LoadConf("card_question")
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
	url , err:= QueryCardQuestionURL(mysqlConn, sql, groupID, voiceUrl+voicePrefix)
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

// QueryCardQuestionURL for the image URL list data through the database query
func QueryCardQuestionURL(DB *sql.DB, sql string, id int64, prefix string) (urls []string, err error) {

	fileRegexp := utils.FileRegexp()

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var (
			contentStr,
			itemStr string
		)

		err := rows.Scan(&contentStr, &itemStr)
		if err != nil {
			fmt.Println(err)
		}else {
			contx := fileRegexp.FindAllString(contentStr,-1)
			item := fileRegexp.FindAllString(itemStr,-1)

			for _, val := range [][]string{contx, item} {
				for _, v := range val {
					hasPre := strings.HasPrefix(v, prefix)
					if hasPre {
						u := strings.Replace(v, prefix, "", -1)
						urls = append(urls, u)
					}
				}
			}
		}
	}
	return
}


