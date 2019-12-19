package core

import (
    "fmt"
    "os"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "github.com/Unknwon/goconfig"
)

func OSS() {
    cfg, err := goconfig.LoadConfigFile("conf/app.ini")
    if err != nil {
        panic("panic")
    }

    endpoint, err := cfg.GetValue("oss","endpoint")
    bucketName, err := cfg.GetValue("oss","bucket")
    acc_key, err := cfg.GetValue("oss","access_key")
    acc_secret, err := cfg.GetValue("oss","access_secret")

    
    client, err := oss.New(endpoint, acc_key, acc_secret)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    objectName := "video/Into_The_Fire.mp4"

    bucket, err := client.Bucket(bucketName)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }


    props, err := bucket.GetObjectDetailedMeta(objectName)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }
    for k := range props {
        fmt.Println(k)
        fmt.Println(props[k])
    }
}
