/*
@Author : Bruce Bee
@Date   : 2019/12/27 11:02
@Email  : mzpy_1119@126.com
*/
// Package core is a core custom method, mainly through the database to get the URL list
package module

import (
	"fmt"
	"database/sql"
	"github.com/Unknwon/goconfig"
	"runtime"
	"strings"
	"../base"
	"../db"
)

// QueryColumnAnswer is get a list of basic data types
func QueryColumnAnswer(groupID int64) (Q []base.BaseInfo) {
	db, _ := db.InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	url , err:= QueryColumnAnswerURL(db, groupID)
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
				b.PicBucket = "jdk3t-qiye"
				b.PicPrefix = "backend_pic/dst/poster/"
				b.PicURL = x
			case "voice":
				b.VoiceURL = x
				b.VoiceBucket ="jdk3t-voice"
				b.VoicePrefix = "backend_voice/"
			case "video":
				b.VideoURL = x
				b.VideoBucket = "jdk3t-video"
				b.VideoPrefix = "video/"
			default:
				fmt.Println("err: no type")
			}
			Q = append(Q, b)
		}
	}
	return
}

// QueryColumnAnswerURL for Get the image URL list data through the database query
func QueryColumnAnswerURL(DB *sql.DB, id int64) (banns map[string][]string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","column_answer")
	if err != nil {
		panic("panic")
	}

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	banns = make(map[string][]string)
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

	banns["pic"] = pp
	banns["voice"] = vo
	banns["video"] = vi

	return
}
