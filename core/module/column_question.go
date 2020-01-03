/*
@Author : Bruce Bee
@Date   : 2019/12/31 10:47
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

type columnQuData struct{
	qsContent sql.NullString
	items sql.NullString
}

// QueryColumnQuestion is get a list of basic data types
func QueryColumnQuestion(groupID int64) (Q []base.BaseInfo) {
	sql, picBucket, picPrefix, picUrl,voiceBucket,voicePrefix,voiceUrl,videoBucket,videoPrefix,videoUrl,docBucket,docPrefix,docUrl := base.LoadConf("column_question")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	url , err:= QueryColumnQuestionURL(mysqlConn, sql, groupID, picUrl+picPrefix, voiceUrl+voicePrefix, videoUrl+videoPrefix,docUrl+docPrefix)
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

// QueryColumnQuestionURL for the image URL list data through the database query
func QueryColumnQuestionURL(DB *sql.DB, sql string, id int64,picPref, voicePref,videoPref, docPref string) (urls map[string][]string, err error) {

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

	var cloqu columnQuData
	for rows.Next() {
		var (
			qsContent,
			items string
		)

		err := rows.Scan(&cloqu.qsContent, &cloqu.items)
		if err != nil {
			fmt.Println(err)
		}else {

			if cloqu.qsContent.Valid {
				qsContent = cloqu.qsContent.String
			}

			if cloqu.items.Valid {
				items = cloqu.items.String
			}

			for _, x := range []string{qsContent, items}{
				if (x != ""){
					c := fileRegexp.FindAllString(x,-1)
					if (len(c) != 0){
						for _, y := range c {
							st1 := strings.HasPrefix(y, picPref)
							if st1 {
								u := strings.Replace(y, picPref, "", -1)
								pp = append(pp, u)
							}

							st2 := strings.HasPrefix(y, voicePref)
							if st2 {
								u := strings.Replace(y, voicePref, "", -1)
								vi = append(vi, u)
							}

							st3 := strings.HasPrefix(y, videoPref)
							if st3 {
								u := strings.Replace(y, videoPref, "", -1)
								vo = append(vo, u)
							}

							st4 := strings.HasPrefix(y, docPref)
							if st4 {
								u := strings.Replace(y, docPref, "", -1)
								doc = append(doc, u)
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
