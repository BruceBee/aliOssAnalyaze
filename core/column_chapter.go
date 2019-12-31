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

	url , err:= QueryColumnChapterURL(db, groupID)
	if nil != err {
		fmt.Println("error")
	}

	for k, u := range url {
		b := BaseInfo{
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

// QueryColumnChapterURL for the image URL list data through the database query
func QueryColumnChapterURL(DB *sql.DB, id int64) (banns map[string][]string, err error) {

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

	for rows.Next() {
		var contentStr string

		err := rows.Scan(&contentStr)
		if err != nil {
			fmt.Println(err)
		}else {
			r := regexp.MustCompile("https://([^:]*?)\\.(mp3|mp4|png|docx|jpg|pptx|gif|doc|pdf)")

			c := r.FindAllString(contentStr,-1)

			for _, x := range c {

				st1 := strings.HasPrefix(x, qiyeOss)
				if st1 {
					u := strings.Replace(x, qiyeOss, "", -1)
					pp = append(pp, u)
				}

				st2 := strings.HasPrefix(x, videoOss)
				if st2 {
					u := strings.Replace(x, videoOss, "", -1)
					vi = append(vi, u)
				}

				st3 := strings.HasPrefix(x, voiceOss)
				if st3 {
					u := strings.Replace(x, voiceOss, "", -1)
					vo = append(vo, u)
				}

				st4 := strings.HasPrefix(x, docOss)
				if st4 {
					u := strings.Replace(x, docOss, "", -1)
					doc = append(doc, u)
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


