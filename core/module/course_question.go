/*
@Author : Bruce Bee
@Date   : 2020/1/2 10:29
@Email  : mzpy_1119@126.com
*/
// package core is ...
package module

import (
	"fmt"
	"runtime"
	"strings"
	"database/sql"
	"github.com/Unknwon/goconfig"
	"../base"
	"../db"
	"../../utils"
)

type quData struct{
	quContent sql.NullString
	item sql.NullString
	analy sql.NullString
}

// QueryCourseQuestion is get a list of basic data types
func QueryCourseQuestion(groupID int64) (Q []base.BaseInfo) {

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	url , err:= QueryCourseQuestionURL(mysqlConn, groupID)
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
				b.VideoBucket ="jdk3t-video"
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

// QueryCourseQuestionURL for the image URL list data through the database query
func QueryCourseQuestionURL(DB *sql.DB, id int64) (banns map[string][]string, err error) {

	fileRegexp := utils.FileRegexp()

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","course_question")
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
		vi ,
		vo ,
		doc []string
	)

	qiyeOss, _ := cfg.GetValue("oss-cdn-url","qiye_oss")
	videoOss, _ := cfg.GetValue("oss-cdn-url","video_oss")
	voiceOss, _ := cfg.GetValue("oss-cdn-url","voice_oss")
	docOss, _ := cfg.GetValue("oss-cdn-url","doc_oss")

	var qu quData
	for rows.Next() {

		var (
			quContent ,
			item,
			analy string
		)
		err := rows.Scan(&qu.quContent, &qu.item, &qu.analy)

		if err != nil {
			fmt.Println(err)
		}else {
			if qu.quContent.Valid {
				quContent = qu.quContent.String
			}

			if qu.item.Valid {
				item = qu.item.String
			}

			if qu.analy.Valid {
				analy = qu.analy.String
			}

			for _, x := range []string{quContent, item, analy}{
				if (x != ""){
					c := fileRegexp.FindAllString(x,-1)
					if (len(c) != 0){
						for _, y := range c {
							st1 := strings.HasPrefix(y, qiyeOss)
							if st1 {
								u := strings.Replace(y, qiyeOss, "", -1)
								pp = append(pp, u)
							}

							st2 := strings.HasPrefix(y, videoOss)
							if st2 {
								u := strings.Replace(y, videoOss, "", -1)
								vi = append(vi, u)
							}

							st3 := strings.HasPrefix(y, voiceOss)
							if st3 {
								u := strings.Replace(y, voiceOss, "", -1)
								vo = append(vo, u)
							}

							st4 := strings.HasPrefix(y, docOss)
							if st4 {
								u := strings.Replace(y, docOss, "", -1)
								doc = append(doc, u)
							}
						}
					}
				}
			}

		}
	}

	banns["pic"] = pp
	banns["video"] = vi
	banns["voice"] = vo
	banns["doc"] = doc

	return
}


