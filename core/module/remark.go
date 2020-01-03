/*
@Author : Bruce Bee
@Date   : 2020/1/2 10:30
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


type remarkData struct{
	pic sql.NullString
	voice sql.NullString
	video sql.NullString
}

// QueryReamrk for a list of basic data types
func QueryReamrk(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, picPrefix, _,voiceBucket,voicePrefix,_,videoBucket,videoPrefix,_,_,_,_ := base.LoadConf("remark")
	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	url , err:= QueryReamrkURL(mysqlConn, sql, groupID)
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

// QueryReamrkURL for the image URL list data through the database query
func QueryReamrkURL(DB *sql.DB, sql string, id int64) (urls map[string][]string, err error) {
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

	var re remarkData
	for rows.Next() {
		err :=rows.Scan(&re.pic, &re.voice, &re.video)
		if err != nil {
			fmt.Println(err)
		}else {

			if re.pic.Valid {
				p := strings.Split(re.pic.String, "|")
				for _, x := range p {
					if (x != ""){
						pp = append(pp, x)
					}
				}
			}

			if re.voice.Valid{
				v := strings.Split(re.voice.String, "|")
				for _, x := range v {
					if (x != ""){
						vo = append(vo, x)
					}
				}
			}

			if re.video.Valid{
				v := strings.Split(re.video.String, "|")
				for _, x := range v {
					if (x != ""){
						vi = append(vi, x)
					}
				}
			}
		}
	}

	urls["pic"] = pp
	urls["voice"] = vo
	urls["video"] = vi

	return
}


