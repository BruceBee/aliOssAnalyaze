/*
@Author : Bruce Bee
@Date   : 2019/12/17 10:17
@Email  : mzpy_1119@126.com
*/
package core

import (
    "../utils"
    "fmt"
    "github.com/Unknwon/goconfig"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "os"
)

// define interface
type Osser interface {
    ReturnSize() string
    ListFile()
}

// define struct of oss client
type OSS struct {
    client *oss.Client
}

// initialization
func InitOSS() Osser{

    act := Active_change{
        User: 1,
        Poster:  "22",
    }
    act_res := act.Query()
    fmt.Println(act_res)


    ban := Banner_nfo{
        GrpID: 1,
        PictureUrl:  "22",
    }
    ban_res := ban.Query()
    fmt.Println(ban_res)





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

// list 
func (o *OSS) ListFile() {
    cfg, _ := goconfig.LoadConfigFile("conf/app.ini")
    bucketName, _ := cfg.GetValue("oss","bucket")

    bucket, _ := o.client.Bucket(bucketName)

    marker := ""
    prefix := oss.Prefix("")    
   
   /**

    lsRes, err := bucket.ListObjects(oss.Marker(marker), prefix)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }


    fmt.Println(lsRes.Prefix)
    
    for _, object := range lsRes.Objects {
        fmt.Println("Bucket: ", object.Key)
        props, _ := bucket.GetObjectDetailedMeta(object.Key)

        s := utils.FormatSize(props["Content-Length"][0])
        fmt.Println(s)
    }
    **/


    deli := "/"   
 
    for {
        lsRes, err := bucket.ListObjects(oss.Marker(marker), prefix, oss.Delimiter(deli))
        // lsRes, err := bucket.ListObjects(oss.Marker(marker), prefix)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }

        for _, dirName := range lsRes.CommonPrefixes {
            fmt.Println("DIR: ", dirName)
        }
       
        // fmt.Println(lsRes.Prefix)

        for _, object := range lsRes.Objects {
            fmt.Println("FILE: ", object.Key)
            props, _ := bucket.GetObjectDetailedMeta(object.Key)
            
            //fmt.Println(object)
      
            s := utils.FormatSize(props["Content-Length"][0])
            fmt.Println("SIZE: ",s)
        }

        
        if lsRes.IsTruncated {
            marker = lsRes.NextMarker
        } else {
            break
        }
       
    }
    
}
