/*
@Author : Bruce Bee
@Date   : 2019/12/31 10:48
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

type resData struct{
	voiceHtml sql.NullString
	voice sql.NullString
	pic sql.NullString
	contentJson sql.NullString
	pcContent sql.NullString
	hybirdContent sql.NullString
}

// QueryCourseCalender is get a list of basic data types
func QueryCourseCalender(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, picPrefix, picUrl,voiceBucket,voicePrefix,voiceUrl,videoBucket,videoPrefix,videoUrl,docBucket,docPrefix,docUrl := base.LoadConf("course_calendar")
	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	url , err:= QueryCourseCalenderURL(mysqlConn, sql, groupID, picUrl+picPrefix, voiceUrl+voicePrefix, videoUrl+videoPrefix,docUrl+docPrefix)
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

// QueryCourseCalenderURL for the image URL list data through the database query
func QueryCourseCalenderURL(DB *sql.DB, sql string, id int64, picPref, voicePref,videoPref, docPref string) (urls map[string][]string, err error) {

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

	var res resData
	for rows.Next() {
		var (
			voiceHtml ,
			voice,
			pic,
			contentJson,
			pcContent,
			hybirdContent string
		)
		err := rows.Scan(&res.voiceHtml, &res.voice, &res.pic, &res.contentJson, &res.pcContent, &res.hybirdContent)

		if err != nil {
			fmt.Println(err)
		}else {
			if res.voiceHtml.Valid {
				voiceHtml = res.voiceHtml.String
			}

			if res.voice.Valid {
				voice = res.voice.String
			}

			if res.pic.Valid {
				pic = res.pic.String
			}

			if res.contentJson.Valid {
				contentJson = res.contentJson.String
			}

			if res.pcContent.Valid {
				pcContent = res.pcContent.String
			}

			if res.hybirdContent.Valid {
				hybirdContent = res.hybirdContent.String
			}

			for _, item := range []string{voiceHtml, voice, pic, contentJson, pcContent, hybirdContent}{
				if (item != ""){
					val := fileRegexp.FindAllString(item,-1)
					if (len(val) != 0){
						for _, v := range val {
							picHasPre := strings.HasPrefix(v, picPref)
							if picHasPre {
								p := strings.Replace(v, picPref, "", -1)
								pp = append(pp, p)
							}

							viHasPre := strings.HasPrefix(v, videoPref)
							if viHasPre {
								i := strings.Replace(v, videoPref, "", -1)
								vi = append(vi, i)
							}

							voHasPre := strings.HasPrefix(v, voicePref)
							if voHasPre {
								o := strings.Replace(v, voicePref, "", -1)
								vo = append(vo, o)
							}

							docHasPre := strings.HasPrefix(v, docPref)
							if docHasPre {
								d := strings.Replace(v, docPref, "", -1)
								doc = append(doc, d)
							}
						}
					}
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

