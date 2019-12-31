/*
@Author : Bruce Bee
@Date   : 2019/12/30 15:49
@Email  : mzpy_1119@126.com
*/

// package core is ...
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

// QueryColumnAnswerRemark is get a list of basic data types
func QueryColumnAnswerRemark(groupID int64) (Q []base.BaseInfo) {
	db, _ := db.InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	url , err:= QueryColumnAnswerRemarkURL(db, groupID)
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
			default:
				fmt.Println("err: no type")
			}
			Q = append(Q, b)
		}
	}
	return
}

// QueryColumnAnswerRemarkURL for Get the image URL list data through the database query
func QueryColumnAnswerRemarkURL(DB *sql.DB, id int64) (banns map[string][]string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","column_answer_remark")
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
		vo []string
	)

	for rows.Next() {
		var (
			pic,
			voice string
		)
		err :=rows.Scan(&pic, &voice)
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

	banns["pic"] = pp
	banns["voice"] = vo

	return
}

