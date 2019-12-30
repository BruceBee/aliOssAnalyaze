/*
@Author : Bruce Bee
@Date   : 2019/12/30 16:21
@Email  : mzpy_1119@126.com
*/
// package core is ...
package core

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"database/sql"
	"github.com/Unknwon/goconfig"
)

// QueryColumnChapter is get a list of basic data types
func QueryColumnChapter(groupID int64) (Q []BaseInfo) {

	db, _ := InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := BaseInfo{
		GrpID: groupID,
		VoiceBucket: "jdk3t-voice",
		VoicePrefix: "backend_voice/",
		TableName: filename,
	}
	url , err:= QueryColumnChapterURL(db, b.GrpID)
	if nil != err {
		fmt.Println("error")
	}

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	voiceOss, _ := cfg.GetValue("oss-cdn-url","voice_oss")

	for _, u := range url {
		if (u != "") {
			b.VoiceURL = strings.Replace(u, voiceOss, "", -1)
			Q = append(Q, b)
		}
	}

	return
}

// QueryColumnChapterURL for the image URL list data through the database query
func QueryColumnChapterURL(DB *sql.DB, id int64) (url []string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","column_chapter")
	if err != nil {
		panic("panic")
	}

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var contentStr string

		err := rows.Scan(&contentStr)
		if err != nil {
			fmt.Println(err)
		}else {
			r := regexp.MustCompile("https://([^:]*?)\\.(mp3|mp4|png|docx|jpg|pptx|gif|doc|pdf)")

			c := r.FindAllString(contentStr,-1)

			for _, x := range c {
				fmt.Println(x)
			}

		}
	}

	return
}


