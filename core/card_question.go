/*
@Author : Bruce Bee
@Date   : 2019/12/27 09:49
@Email  : mzpy_1119@126.com
*/
// Package core is a core custom method, mainly through the database to get the URL list
package core

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"database/sql"
	"encoding/json"
	"github.com/Unknwon/goconfig"
)

// CardQuestion...
type CardQuestion struct {
	Text string `json:"text"`
	Notes string `json:"notes"`
	EvalTime int `json:"evalTime"`
	EvalLimit int `json:"evalLimit"`
	Voice CardQuestionVoice `json:"voice"` 
} 

// CardQuestionVoice ...
type CardQuestionVoice struct {
	VoiceURL string `json:"voice_url"`
	VoiceName string `json:"voice_name"`
	Name string `json:"name"`
	VoiceAvater string `json:"voice_avater"`
}

// IsEmpty for check sturct is empty
func (c CardQuestion) IsEmpty() bool {
	return reflect.DeepEqual(c, CardQuestion{})
}


// QueryCardQuestion is get a list of basic data types
func QueryCardQuestion(groupID int64) (Q []BaseInfo) {

	db, _ := InitDB()
	_, file, _, _ := runtime.Caller(0)
	f := strings.Split(file, "/")
	filename :=strings.Split(f[len(f)-1], ".")[0]
	b := BaseInfo{
		GrpID: groupID,
		VoiceBucket: "jdk3t-voice",
		VoicePrefix: "backend_voice/",
		TableName: filename,
	}
	url , err:= QueryCardQuestionURL(db, b.GrpID)
	if nil != err {
		fmt.Println("error")
	}

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	voiceOss, _ := cfg.GetValue("oss-cdn-url","voice_oss")

	for _, u := range url {
		if (u != "") {
			b.VoiceURL = strings.Replace(u, voiceOss, "", -1)
			Q = append(Q, b)
		}
	}

	return
}

// QueryCardQuestionURL for the image URL list data through the database query
func QueryCardQuestionURL(DB *sql.DB, id int64) (url []string, err error) {

	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err := cfg.GetValue("sql","card_question")
	if err != nil {
		panic("panic")
	}

	rows, err := DB.Query(sql, id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var (
			contentStr,
			itemStr string
		)

		err := rows.Scan(&contentStr, &itemStr)
		if err != nil {
			fmt.Println(err)
		}else {
			var (
				question,
				item CardQuestion
			)
			err = json.Unmarshal([]byte(contentStr), &question)
			if err == nil{
				if !question.IsEmpty(){
					st := reflect.ValueOf(question)
					st2 := st.FieldByName("Voice").FieldByName("VoiceURL").String()
					if(st2 != ""){
						url = append(url, st2)
					}
				}
			}


			err = json.Unmarshal([]byte(itemStr), &item)
			if err == nil{
				if !item.IsEmpty(){
					st := reflect.ValueOf(item)
					st2 := st.FieldByName("Voice").FieldByName("VoiceURL").String()
					if(st2 != ""){
						url = append(url, st2)
					}
				}
			}
		}
	}

	return
}


