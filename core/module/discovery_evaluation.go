/*
@Author : Bruce Bee
@Date   : 2020/1/2 10:29
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

type evalData struct{
	evalSrc sql.NullString
	evalTitle sql.NullString
	courseSrc sql.NullString
}

// QueryDiscoveryEvaluation is get a list of basic data types
func QueryDiscoveryEvaluation(groupID int64) (Q []base.BaseInfo) {

	sql, picBucket, picPrefix, picUrl,_,_,_,_,_,_,_,_,_ := base.LoadConf("discovery_evaluation")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	url , err:= QueryDiscoveryEvaluationURL(mysqlConn, sql, groupID, picUrl+picPrefix)
	if nil != err {
		fmt.Println("error")
	}

	b := base.BaseInfo{
		GrpID: groupID,
		PicBucket: picBucket,
		PicPrefix: picPrefix,
		TableName: filename,
	}

	for _, u := range url {
		if (u != "") {
			b.PicURL = u
			Q = append(Q, b)
		}
	}

	return
}

// QueryDiscoveryEvaluationURL for the image URL list data through the database query
func QueryDiscoveryEvaluationURL(DB *sql.DB, sql string, id int64, prefix string) (urls []string, err error) {

	fileRegexp := utils.FileRegexp()

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	var eval evalData

	for rows.Next() {
		err := rows.Scan(&eval.evalSrc, &eval.evalTitle, &eval.courseSrc)

		if err != nil {
			fmt.Println(err)
		}else {
			if eval.evalSrc.Valid  {
				picHasPre := strings.HasPrefix(eval.evalSrc.String, prefix)
				if picHasPre {
					es := strings.Replace(eval.evalSrc.String, prefix, "", -1)
					urls = append(urls, es)
				}
			}

			if eval.evalTitle.Valid {
				if (eval.evalTitle.String != ""){
					et := fileRegexp.FindAllString(eval.evalTitle.String,-1)
					if (len(et) != 0){
						for _, v := range et {
							picHasPre := strings.HasPrefix(v, prefix)
							if picHasPre {
								u := strings.Replace(v, prefix, "", -1)
								urls = append(urls, u)
							}
						}
					}
				}
			}

			if eval.courseSrc.Valid {
				picHasPre := strings.HasPrefix(eval.evalSrc.String, prefix)
				if picHasPre {
					es := strings.Replace(eval.evalSrc.String, prefix, "", -1)
					urls = append(urls, es)
				}
			}
		}
	}

	return
}


