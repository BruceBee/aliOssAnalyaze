/*
@Author : Bruce Bee
@Date   : 2019/12/24 15:30
@Email  : mzpy_1119@126.com
*/

// Custom method, mainly through the database to get the URL list
package core

import (
	"database/sql"
	"fmt"
)

// QueryCard , Gets a list of basic data types
func QueryCard(groupID int64) (Q []BaseInfo) {

	db, _ := InitDB()
	b := BaseInfo{
		GrpID: groupID,
		VoiceBucket: "jdk3t-voice",
		VoicePrefix: "backend_voice/",
		TableName: "jdk_card_answer",
	}
	url , err:= QueryCardAnswerURL(db, b.GrpID)
	if nil != err {
		fmt.Println("error")
	}

	for _, u := range url {
		if (u != "") {
			b.VoiceURL = u
			Q = append(Q, b)
		}
	}
	return
}

// QueryCardAnswerURL, Get the image URL list data through the database query
func QueryCardAnswerURL(DB *sql.DB, id int64) (banns []string, err error) {

	rows, err := DB.Query("SELECT voices FROM jdk_card_answer WHERE group_id= ? GROUP BY voices;", id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var bann string
		rows.Scan(&bann)
		banns = append(banns, bann)
	}
	return
}

