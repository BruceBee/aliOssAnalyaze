/*
@Author : Bruce Bee
@Date   : 2019/12/24 15:30
@Email  : mzpy_1119@126.com
*/

package core

import (
	"database/sql"
	"fmt"
)

// QueryCard ...
func QueryCard(groupID int64) []BaseInfo {

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
	var Q []BaseInfo

	for _, u := range url {
		if (u != "") {
			b.VoiceURL = u
			Q = append(Q, b)
		}

	}
	return Q
}

// QueryCardAnswerURL ...
func QueryCardAnswerURL(DB *sql.DB, id int64) ([]string, error) {

	var banns []string
	rows, err := DB.Query("SELECT voices FROM jdk_card_answer WHERE group_id= ?", id)
	if nil != err {
		fmt.Println("QueryRow Error", err)
	}

	for rows.Next() {
		var bann string
		rows.Scan(&bann)
		banns = append(banns, bann)
	}
	return banns, nil
}

