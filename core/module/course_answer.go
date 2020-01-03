/*
@Author : Bruce Bee
@Date   : 2019/12/31 10:48
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

// QueryCourseAnswer for a list of basic data types
func QueryCourseAnswer(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, picPrefix, _,voiceBucket,voicePrefix,_,videoBucket,videoPrefix,_,_,_,_ := base.LoadConf("course_answer")
	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	url , err:= QueryCourseAnswerURL(mysqlConn, sql, groupID)
	if nil != err {
		fmt.Println("error")
	}

	for k, u := range url {
		b := base.BaseInfo{
			GrpID: groupID,
			TableName: filename,
		}

		for _, x := range u {
			switch k {
			case "pic":
				b.PicBucket = picBucket
				b.PicPrefix = picPrefix
				b.PicURL = x
			case "video":
				b.VideoURL = x
				b.VideoBucket = voiceBucket
				b.VideoPrefix = voicePrefix
			case "voice":
				b.VoiceURL = x
				b.VoiceBucket =videoBucket
				b.VoicePrefix = videoPrefix
			default:
				fmt.Println("err: no type")
			}
			Q = append(Q, b)
		}
	}
	return
}

// QueryCourseAnswerURL for the image URL list data through the database query
func QueryCourseAnswerURL(DB *sql.DB, sql string, id int64) (urls map[string][]string, err error) {

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	urls = make(map[string][]string)
	var (
		pp ,
		vi ,
		vo []string
	)

	for rows.Next() {
		var (
			pic,
			video,
			voice string
		)
		err :=rows.Scan(&pic, &video, &voice)
		if err != nil {
			fmt.Println(err)
		}else {
			if (pic != ""){
				p := strings.Split(pic, "|")
				for _, x := range p {
					if (x != ""){
						pp = append(pp, x)
					}

				}
			}

			if (video != ""){
				v := strings.Split(video, "|")
				for _, x := range v {
					if (x != ""){
						vi = append(vi, x)
					}

				}
			}

			if (voice != ""){
				v := strings.Split(voice, "|")
				for _, x := range v {
					if (x != ""){
						vo = append(vo, x)
					}

				}
			}
		}
	}

	urls["pic"] = pp
	urls["video"] = vi
	urls["voice"] = vo

	return
}

