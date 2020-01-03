/*
@Author : Bruce Bee
@Date   : 2019/12/27 11:02
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

// QueryColumnAnswer is get a list of basic data types
func QueryColumnAnswer(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, picPrefix, _,voiceBucket,voicePrefix,_,videoBucket,videoPrefix,_,_,_,_ := base.LoadConf("column_answer")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()

	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	url , err:= QueryColumnAnswerURL(mysqlConn, sql, groupID)
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
			case "voice":
				b.VoiceURL = x
				b.VoiceBucket = voiceBucket
				b.VoicePrefix = voicePrefix
			case "video":
				b.VideoURL = x
				b.VideoBucket = videoBucket
				b.VideoPrefix = videoPrefix
			default:
				fmt.Println("err: no type")
			}
			Q = append(Q, b)
		}
	}
	return
}

// QueryColumnAnswerURL for Get the image URL list data through the database query
func QueryColumnAnswerURL(DB *sql.DB, sql string, id int64) (urls map[string][]string, err error) {

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	urls = make(map[string][]string)
	var (
		pp ,
		vo ,
		vi []string
	)

	for rows.Next() {
		var (
			pic,
			voice,
			video string
		)
		err :=rows.Scan(&pic, &voice, &video)
		if err != nil {
			fmt.Println(err)
		}else {

			if (pic != ""){
				p := strings.Split(pic, ";")
				for _, x := range p {
					pp = append(pp, x)
				}
			}

			if (voice != ""){
				v := strings.Split(voice, ";")
				for _, x := range v {
					vo = append(vo, x)
				}
			}

			if (video != ""){
				v := strings.Split(video, ";")
				for _, x := range v {
					vi = append(vi, x)
				}
			}
		}
	}

	urls["pic"] = pp
	urls["voice"] = vo
	urls["video"] = vi

	return
}
