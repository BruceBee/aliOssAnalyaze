/*
@Author : Bruce Bee
@Date   : 2019/12/30 15:49
@Email  : mzpy_1119@126.com
*/

// package core is ...
package module

import (
	"../base"
	"../db"
	"database/sql"
	"fmt"
	"runtime"
	"strings"
)

// QueryColumnAnswerRemark is get a list of basic data types
func QueryColumnAnswerRemark(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, picPrefix, _,voiceBucket,voicePrefix,_,_,_,_,_,_,_ := base.LoadConf("column_answer_remark")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	url , err:= QueryColumnAnswerRemarkURL(mysqlConn, sql, groupID)
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
			default:
				fmt.Println("err: no type")
			}
			Q = append(Q, b)
		}
	}
	return
}

// QueryColumnAnswerRemarkURL for Get the image URL list data through the database query
func QueryColumnAnswerRemarkURL(DB *sql.DB, sql string, id int64) (urls map[string][]string, err error) {
	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	urls = make(map[string][]string)
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

	urls["pic"] = pp
	urls["voice"] = vo

	return
}

