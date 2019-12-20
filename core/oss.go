package core

import (
    "fmt"
    "os"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "github.com/Unknwon/goconfig"
)

// define interface
type Osser interface {
    ReturnSize() string
}

// define struct of oss client
type OSS struct {
    client *oss.Client
}

// initialization
func InitOSS() Osser{
    cfg, err := goconfig.LoadConfigFile("conf/app.ini")
    if err != nil {
        panic("panic")
    }

    endpoint, err := cfg.GetValue("oss","endpoint")
    acc_key, err := cfg.GetValue("oss","access_key")
    acc_secret, err := cfg.GetValue("oss","access_secret")

    c, err := oss.New(endpoint, acc_key, acc_secret)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    return &OSS{
        client:c,
    }
}

// get file size
func (o *OSS) ReturnSize() string {
    
    cfg, err := goconfig.LoadConfigFile("conf/app.ini")
    bucketName, err := cfg.GetValue("oss","bucket")
    objectName := "video/Into_The_Fire.mp4"

    bucket, err := o.client.Bucket(bucketName)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    props, err := bucket.GetObjectDetailedMeta(objectName)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }
    return props["Content-Length"][0]
}

