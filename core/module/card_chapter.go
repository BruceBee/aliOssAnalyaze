/*
@Author : Bruce Bee
@Date   : 2019/12/26 15:27
@Email  : mzpy_1119@126.com
*/

package module

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"database/sql"
	"encoding/json"
	"github.com/Unknwon/goconfig"
	"../base"
	"../db"
)

// CardChapter is ...
type CardChapter struct {
	Type string `json:"type"`
	Key string `json:"key"`
	Content []CardChapterContent `json:"content"`
}

// CardChapterContent is ...
type CardChapterContent struct {
	PicURL string `json:"picture_url"`
	PicName string `json:"picture_name"`
	PicWidth string `json:"picture_width"`
	PicHeight string `json:"picture_height"`
	PicPosition string `json:"picture_position"`
}

// IsEmpty for check sturct is empty
func (c CardChapter) IsEmpty() bool {
	return reflect.DeepEqual(c, CardChapter{})
}

// QueryCardChapter is get a list of basic data types
func QueryCardChapter(groupID int64) (Q []base.BaseInfo) {

	mysqlConn, _ := db.InitDB()
	defer mysqlConn.Close()

	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := base.BaseInfo{
		GrpID: groupID,
		PicBucket: "jdk3t-qiye",
		PicPrefix: "backend_pic/dst/poster/",
		TableName: filename,
	}
	url , err:= QueryCardChapterURL(mysqlConn, b.GrpID)
	if nil != err {
		fmt.Println("error")
	}

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	qiyeOss, _ := cfg.GetValue("oss-cdn-url","qiye_oss")

	for _, u := range url {
		if (u != "") {
			b.PicURL = strings.Replace(u, qiyeOss, "", -1)
			Q = append(Q, b)
		}
	}

	return
}

// QueryCardChapterURL for the image URL list data through the database query
func QueryCardChapterURL(DB *sql.DB, id int64) (url []string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","card_chapter")
	if err != nil {
		panic("panic")
	}

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
			var card []CardChapter
			err = json.Unmarshal([]byte(cardStr), &card)
			if err != nil{
				fmt.Println(err)
			}else {
				if (len(card) != 0){
					st := reflect.ValueOf(card[0])
					st2 := st.FieldByName("Content").Index(0).FieldByName("PicURL").String()
					if(st2 != ""){
						url = append(url, st2)
					}
				}
			}
		}
	}

	return
}

