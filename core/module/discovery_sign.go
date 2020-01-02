/*
@Author : Bruce Bee
@Date   : 2020/1/2 10:30
@Email  : mzpy_1119@126.com
*/

// Package module is a core custom method, mainly through the database to get the URL list
package module

import (
	"fmt"
	"database/sql"
	"github.com/Unknwon/goconfig"
	"runtime"
	"strings"
	"../base"
	"../db"
	"../../utils"
)

type discovSignData struct{
	pcContent sql.NullString
	qrImg sql.NullString
}

// QueryDiscoverySign for a list of basic data types
func QueryDiscoverySign(groupID int64) (Q []base.BaseInfo) {
	db, _ := db.InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	url , err:= QueryDiscoverySignURL(db, groupID)
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
			case "qr":
				b.PicBucket = "jdk3t-qiye"
				b.PicPrefix = ""
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

// QueryDiscoverySignURL for the image URL list data through the database query
func QueryDiscoverySignURL(DB *sql.DB, id int64) (banns map[string][]string, err error) {
	fileRegexp := utils.FileRegexp()

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","discovery_sign")
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
		qr ,
		vi ,
		vo ,
		doc []string
	)

	qiyeOss, _ := cfg.GetValue("oss-cdn-url","qiye_oss")
	videoOss, _ := cfg.GetValue("oss-cdn-url","video_oss")
	voiceOss, _ := cfg.GetValue("oss-cdn-url","voice_oss")
	docOss, _ := cfg.GetValue("oss-cdn-url","doc_oss")


	var disCovSign discovSignData
	for rows.Next() {
		var (
			pcContent,
			qrImg string
		)
		err := rows.Scan(&disCovSign.pcContent, &disCovSign.qrImg)

		if err != nil {
			fmt.Println(err)
		}else {
			if disCovSign.qrImg.Valid {
				if (disCovSign.qrImg.String != ""){
					fmt.Println(disCovSign.qrImg.String)
					qrImg = disCovSign.qrImg.String[1:]
					qr = append(qr, qrImg)
				}
			}

			if disCovSign.pcContent.Valid {
				pcContent = disCovSign.pcContent.String
				pc := fileRegexp.FindAllString(pcContent,-1)
				if (len(pc) != 0){
					for _, p := range pc {
						st1 := strings.HasPrefix(p, qiyeOss)
						if st1 {
							u := strings.Replace(p, qiyeOss, "", -1)
							pp = append(pp, u)
						}

						st2 := strings.HasPrefix(p, videoOss)
						if st2 {
							u := strings.Replace(p, videoOss, "", -1)
							vi = append(vi, u)
						}

						st3 := strings.HasPrefix(p, voiceOss)
						if st3 {
							u := strings.Replace(p, voiceOss, "", -1)
							vo = append(vo, u)
						}

						st4 := strings.HasPrefix(p, docOss)
						if st4 {
							u := strings.Replace(p, docOss, "", -1)
							doc = append(doc, u)
						}
					}
				}
			}

		}
	}

	banns["pic"] = pp
	banns["qr"] = qr
	banns["video"] = vi
	banns["voice"] = vo
	banns["doc"] = doc

	return
}



