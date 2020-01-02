/*
@Author : Bruce Bee
@Date   : 2020/1/2 10:30
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

// QuerySubmit is get a list of basic data types
func QuerySubmit(groupID int64) (Q []base.BaseInfo) {
	db, _ := db.InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	url , err:= QuerySubmitURL(db, groupID)
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
			case "doc":
				b.DocURL = x
				b.DocBucket ="jdk3t-doc"
				b.DocPrefix = "document/"
			default:
				fmt.Println("err: no type")
			}
			Q = append(Q, b)
		}
	}
	return
}

// QuerySubmitURL for Get the image URL list data through the database query
func QuerySubmitURL(DB *sql.DB, id int64) (banns map[string][]string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","submit")
	if err != nil {
		panic("panic")
	}

	var name string
	tableName := DB.QueryRow("SELECT submit_table_name from jdk_submit_relation WHERE group_id = ?", id)

	tableName.Scan(&name)

	sql = fmt.Sprintf("%s %s WHERE group_id = ?", sql, name)

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	banns = make(map[string][]string)
	var (
		pp ,
		vo ,
		vi ,
		dd []string
	)

	for rows.Next() {
		var (
			pic,
			voice,
			video,
			doc string
		)
		err :=rows.Scan(&pic, &voice, &video, &doc)
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
			if (doc != ""){
				v := strings.Split(doc, ";")
				for _, x := range v {
					dd = append(dd, x)
				}
			}
		}
	}

	banns["pic"] = pp
	banns["voice"] = vo
	banns["video"] = vi
	banns["doc"] = dd

	return
}
