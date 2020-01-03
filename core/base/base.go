/*
@Author : Bruce Bee
@Date   : 2019/12/24 15:35
@Email  : mzpy_1119@126.com
*/

package base

import "github.com/Unknwon/goconfig"

// BaseInfo ...
type BaseInfo struct {
	GrpID int64
	PicBucket string
	PicPrefix string
	PicURL string
	VoiceBucket string
	VoicePrefix string
	VoiceURL string
	VideoBucket string
	VideoPrefix string
	VideoURL string
	DocBucket string
	DocPrefix string
	DocURL string
	TableName string
}


// LoadConf
func LoadConf(sqlString string) (sql,
	picBucket,
	picPrefix,
	picUrl,
	voiceBucket,
	voicePrefix,
	voiceUrl,
	videoBucket,
	videoPrefix,
	videoUrl,
	docBucket,
	docPrefix,
	docUrl string ) {
	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic("panic")
	}

	sql, err = cfg.GetValue("sql",sqlString)
	if err != nil {
		panic("panic")
	}

	picBucket,  err = cfg.GetValue("oss-info","pic_bucket")
	if err != nil {
		panic("pic_bucket panic")
	}

	picPrefix, err = cfg.GetValue("oss-info","pic_prefix")
	if err != nil {
		panic("pic_prefix panic")
	}

	picUrl, err = cfg.GetValue("oss-info","pic_url")
	if err != nil {
		panic("pic_url panic")
	}

	voiceBucket, err = cfg.GetValue("oss-info","voice_bucket")
	if err != nil {
		panic("voice_bucket panic")
	}

	voicePrefix, err = cfg.GetValue("oss-info","voice_prefix")
	if err != nil {
		panic("voice_prefix panic")
	}

	voiceUrl, err = cfg.GetValue("oss-info","voice_url")
	if err != nil {
		panic("voice_url panic")
	}

	videoBucket, err = cfg.GetValue("oss-info","video_bucket")
	if err != nil {
		panic("video_bucket panic")
	}

	videoPrefix, err = cfg.GetValue("oss-info","video_prefix")
	if err != nil {
		panic("video_prefix panic")
	}

	videoUrl, err = cfg.GetValue("oss-info","video_url")
	if err != nil {
		panic("video_url panic")
	}

	docBucket, err = cfg.GetValue("oss-info","doc_bucket")
	if err != nil {
		panic("doc_bucket panic")
	}

	docPrefix, err = cfg.GetValue("oss-info","doc_prefix")
	if err != nil {
		panic("doc_prefix panic")
	}

	docUrl, err = cfg.GetValue("oss-info","doc_url")
	if err != nil {
		panic("doc_url panic")
	}
	return
}