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

type evalData struct{
	evalSrc sql.NullString
	evalTitle sql.NullString
	courseSrc sql.NullString
}

// QueryDiscoveryEvaluation is get a list of basic data types
func QueryDiscoveryEvaluation(groupID int64) (Q []base.BaseInfo) {

	db, _ := db.InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]

	url , err:= QueryDiscoveryEvaluationURL(db, groupID)
	if nil != err {
		fmt.Println("error")
	}

	b := base.BaseInfo{
		GrpID: groupID,
		PicBucket: "jdk3t-qiye",
		PicPrefix: "backend_pic/dst/poster/",
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
func QueryDiscoveryEvaluationURL(DB *sql.DB, id int64) (banns []string, err error) {

	fileRegexp := utils.FileRegexp()

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","discovery_evaluation")
	if err != nil {
		panic("panic")
	}

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	qiyeOss, _ := cfg.GetValue("oss-cdn-url","qiye_oss")

	var eval evalData

	for rows.Next() {
		err := rows.Scan(&eval.evalSrc, &eval.evalTitle, &eval.courseSrc)

		if err != nil {
			fmt.Println(err)
		}else {
			if eval.evalSrc.Valid  {
				st1 := strings.HasPrefix(eval.evalSrc.String, qiyeOss)
				if st1 {
					u := strings.Replace(eval.evalSrc.String, qiyeOss, "", -1)
					banns = append(banns, u)
				}
			}

			if eval.evalTitle.Valid {
				if (eval.evalTitle.String != ""){
					c := fileRegexp.FindAllString(eval.evalTitle.String,-1)
					if (len(c) != 0){
						for _, y := range c {
							st1 := strings.HasPrefix(y, qiyeOss)
							if st1 {
								u := strings.Replace(y, qiyeOss, "", -1)
								banns = append(banns, u)
							}
						}
					}
				}
			}

			if eval.courseSrc.Valid {
				st1 := strings.HasPrefix(eval.evalSrc.String, qiyeOss)
				if st1 {
					u := strings.Replace(eval.evalSrc.String, qiyeOss, "", -1)
					banns = append(banns, u)
				}
			}
		}
	}

	return
}


