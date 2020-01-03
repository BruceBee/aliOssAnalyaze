/*
@Author : Bruce Bee
@Date   : 2019/12/30 16:21
@Email  : mzpy_1119@126.com
*/

// package core is ...
package module

import (
	"../../utils"
	"../base"
	"../db"
	"database/sql"
	"fmt"
	"runtime"
	"strings"
)

// QueryColumnChapter is get a list of basic data types
func QueryColumnChapter(groupID int64) (Q []base.BaseInfo) {

	sql, picBucket, picPrefix, picUrl,voiceBucket,voicePrefix,voiceUrl,videoBucket,videoPrefix,videoUrl,docBucket,docPrefix,docUrl := base.LoadConf("column_chapter")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	url , err:= QueryColumnChapterURL(mysqlConn, sql, groupID, picUrl+picPrefix, voiceUrl+voicePrefix, videoUrl+videoPrefix,docUrl+docPrefix)
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
			case "doc":
				b.DocURL = x
				b.DocBucket = docBucket
				b.DocPrefix = docPrefix
			default:
				fmt.Println("err: no type")
			}
			Q = append(Q, b)
		}
	}

	return
}

// QueryColumnChapterURL for the image URL list data through the database query
func QueryColumnChapterURL(DB *sql.DB, sql string, id int64, picPref, voicePref,videoPref, docPref string) (urls map[string][]string, err error) {

	fileRegexp := utils.FileRegexp()

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	urls = make(map[string][]string)
	var (
		pp ,
		vi ,
		vo ,
		doc []string
	)

	for rows.Next() {
		var contentStr string

		err := rows.Scan(&contentStr)
		if err != nil {
			fmt.Println(err)
		}else {
			c := fileRegexp.FindAllString(contentStr,-1)

			for _, x := range c {

				st1 := strings.HasPrefix(x, picPref)
				if st1 {
					u := strings.Replace(x, picPref, "", -1)
					pp = append(pp, u)
				}

				st2 := strings.HasPrefix(x, videoPref)
				if st2 {
					u := strings.Replace(x, videoPref, "", -1)
					vi = append(vi, u)
				}

				st3 := strings.HasPrefix(x, voicePref)
				if st3 {
					u := strings.Replace(x, voicePref, "", -1)
					vo = append(vo, u)
				}

				st4 := strings.HasPrefix(x, docPref)
				if st4 {
					u := strings.Replace(x, docPref, "", -1)
					doc = append(doc, u)
				}
			}

		}
	}

	urls["pic"] = pp
	urls["video"] = vi
	urls["voice"] = vo
	urls["doc"] = doc

	return
}


