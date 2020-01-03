/*
@Author : Bruce Bee
@Date   : 2020/1/2 10:30
@Email  : mzpy_1119@126.com
*/

// Package module is a core custom method, mainly through the database to get the URL list
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

type sourceData struct{
	url sql.NullString
	kind string
	pcContent sql.NullString
}

// QuerySource for a list of basic data types
func QuerySource(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, picPrefix, picUrl,voiceBucket,voicePrefix,voiceUrl,videoBucket,videoPrefix,videoUrl,docBucket,docPrefix,docUrl := base.LoadConf("source")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	url , err:= QuerySourceURL(mysqlConn, sql, groupID, picUrl+picPrefix, voiceUrl+voicePrefix, videoUrl+videoPrefix,docUrl+docPrefix)
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

// QuerySourceURL for the image URL list data through the database query
func QuerySourceURL(DB *sql.DB, sql string, id int64, picPref, voicePref,videoPref, docPref string) (urls map[string][]string, err error) {
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

	var fileType = map[string][]string{"1":pp, "2":vo, "3":vi, "4":doc}

	var source sourceData
	for rows.Next() {
		var (
			url,
			pcContent string
		)
		err := rows.Scan(&source.url, &source.kind, &source.pcContent)

		if err != nil {
			fmt.Println(err)
		}else {

			if source.url.Valid {
				url = strings.Replace(source.url.String, "document/", "", -1)
				url = strings.Replace(url, "backend_pic/dst/poster/", "", -1)
				url = strings.Replace(url, "video/", "", -1)
				if (url != ""){
					fileType[source.kind] = append(fileType[source.kind], url)
				}
			}

			if source.pcContent.Valid {
				pcContent = source.pcContent.String
				pc := fileRegexp.FindAllString(pcContent,-1)
				if (len(pc) != 0){
					for _, p := range pc {
						picHasPre := strings.HasPrefix(p, picPref)
						if picHasPre {
							u := strings.Replace(p, picPref, "", -1)
							pp = append(pp, u)
						}

						viHasPre := strings.HasPrefix(p, videoPref)
						if viHasPre {
							i := strings.Replace(p, videoPref, "", -1)
							vi = append(vi, i)
						}

						voHasPre := strings.HasPrefix(p, voicePref)
						if voHasPre {
							o := strings.Replace(p, voicePref, "", -1)
							vo = append(vo, o)
						}

						docHasPre := strings.HasPrefix(p, docPref)
						if docHasPre {
							d := strings.Replace(p, docPref, "", -1)
							doc = append(doc, d)
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


