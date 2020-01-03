/*
@Author : Bruce Bee
@Date   : 2019/12/26 15:27
@Email  : mzpy_1119@126.com
*/

package module

import (
	"fmt"
	"runtime"
	"strings"
	"database/sql"
	"../base"
	"../db"
	"../../utils"
)

// QueryCardChapter is get a list of basic data types
func QueryCardChapter(groupID int64) (Q []base.BaseInfo) {

	sql, picBucket, picPrefix, picUrl,_,_,_,_,_,_,_,_,_ := base.LoadConf("card_chapter")

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()

	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := base.BaseInfo{
		GrpID: groupID,
		PicBucket: picBucket,
		PicPrefix: picPrefix,
		TableName: filename,
	}
	url , err:= QueryCardChapterURL(mysqlConn, sql, groupID,  picUrl+picPrefix )
	if nil != err {
		fmt.Println("error")
	}

	for _, u := range url {
		if (u != "") {
			b.PicURL = u
			Q = append(Q, b)
		}
	}

	return
}

// QueryCardChapterURL for the image URL list data through the database query
func QueryCardChapterURL(DB *sql.DB, sql string, id int64, prefix string) (urls []string, err error) {

	fileRegexp := utils.FileRegexp()

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var cardStr string
		err := rows.Scan(&cardStr)

		if err != nil {
			fmt.Println(err)
		}else {
			card := fileRegexp.FindAllString(cardStr,-1)
			for _, cc := range card {
				hasPre := strings.HasPrefix(cc, prefix)
				if hasPre {
					u := strings.Replace(cc, prefix, "", -1)
					urls = append(urls, u)
				}
			}
		}
	}

	return
}

