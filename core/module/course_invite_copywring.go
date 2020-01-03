/*
@Author : Bruce Bee
@Date   : 2020/1/2 10:28
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

type inviteData struct{
	tripContent sql.NullString
	tripQrCode sql.NullString
}

// QueryCourseInviteCopywring is get a list of basic data types
func QueryCourseInviteCopywring(groupID int64) (Q []base.BaseInfo) {

	sql, picBucket, picPrefix, picUrl,voiceBucket,voicePrefix,voiceUrl,videoBucket,videoPrefix,videoUrl,docBucket,docPrefix,docUrl := base.LoadConf("course_invite_copywring")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	url , err:= QueryCourseInviteCopywringURL(mysqlConn, sql, groupID, picUrl+picPrefix, voiceUrl+voicePrefix, videoUrl+videoPrefix,docUrl+docPrefix)
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

// QueryCourseInviteCopywringURL for the image URL list data through the database query
func QueryCourseInviteCopywringURL(DB *sql.DB, sql string, id int64, picPref, voicePref,videoPref, docPref string) (urls map[string][]string, err error) {

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

	var inv inviteData
	for rows.Next() {

		var (
			tripContent ,
			tripQrCode string
		)
		err := rows.Scan(&inv.tripContent, &inv.tripQrCode)

		if err != nil {
			fmt.Println(err)
		}else {
			if inv.tripContent.Valid {
				tripContent = inv.tripContent.String
			}

			if inv.tripQrCode.Valid {
				tripQrCode = inv.tripQrCode.String
			}

			for _, item := range []string{tripContent, tripQrCode}{
				if (item != ""){
					val := fileRegexp.FindAllString(item,-1)
					if (len(val) != 0){
						for _, v := range val {
							st1 := strings.HasPrefix(v, picPref)
							if st1 {
								p := strings.Replace(v, picPref, "", -1)
								pp = append(pp, p)
							}

							st2 := strings.HasPrefix(v, videoPref)
							if st2 {
								i := strings.Replace(v, videoPref, "", -1)
								vi = append(vi, i)
							}

							st3 := strings.HasPrefix(v, voicePref)
							if st3 {
								o := strings.Replace(v, voicePref, "", -1)
								vo = append(vo, o)
							}

							st4 := strings.HasPrefix(v, docPref)
							if st4 {
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


