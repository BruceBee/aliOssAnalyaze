/*
@Author : Bruce Bee
@Date   : 2019/12/17 10:17
@Email  : mzpy_1119@126.com
*/

// Package core ...
package core

import (
    "../utils"
    "fmt"
    "github.com/Unknwon/goconfig"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "os"
    "strconv"
    "strings"
)

// Osser ...
type Osser interface {
    ReturnSize(int64) error
    ListFile()
}

// OSS ...
type OSS struct {
    client *oss.Client
}

// Register ...
func Register(group_id int64) (*BaseInfo, []string){

    b, banRes := QueryBanner(group_id)
    //c, cardRes := QueryCard(group_id)
    //for _, x := range cardRes {
    //    fmt.Println(x)
    //}
    return b, banRes
}

// InitOSS initialization
func InitOSS() Osser{

    cfg, err := goconfig.LoadConfigFile("conf/app.ini")
    if err != nil {
        panic("panic")
    }

    endPoint, err := cfg.GetValue("oss","endpoint")
    accKey, err := cfg.GetValue("oss","access_key")
    accSecret, err := cfg.GetValue("oss","access_secret")

    c, err := oss.New(endPoint, accKey, accSecret)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    return &OSS{
        client:c,
    }
}


// ReturnSize get file size
func (o *OSS) ReturnSize(group_id int64) error {

    // partLint
    partLine := "-"
    for i := 0; i < 60; i++ {
        partLine += "-"
    }

    ban, banRes := Register(group_id)

    fileName := []string{}
    fileName = append(fileName, file_path)
    fileName = append(fileName, ban.TableName)


    if (ban.PictureBucket != "" ){
        fName :=  fmt.Sprintf("%s%s", strings.Join(fileName, ""), "_Pic")
        totalSize := 0
        totalCount := 0

        bucket, err := o.client.Bucket(ban.PictureBucket)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }

        CreateFile(fName, fmt.Sprintf("GroupID: %d ; Bucket: %s ; Path: %s\n",ban.GrpID, ban.PictureBucket,ban.PicturePrefix ))
        CreateFile(fName,partLine + "\n")

        for _, b := range banRes {
            props, err := bucket.GetObjectDetailedMeta(ban.PicturePrefix + b)
            if err != nil {
                fmt.Println("Error:", err)
                os.Exit(-1)
            }

            Cont := utils.FormatSize(props["Content-Length"][0])
            ContentLength, _ :=  strconv.Atoi(props["Content-Length"][0])

            totalSize += ContentLength
            totalCount += 1

            fmt.Printf("%s | %s\n", Cont, b)
            CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, b))
        }

        t := strconv.Itoa(totalSize)

        CreateFile(fName,partLine + "\n")
        CreateFile(fName,fmt.Sprintf("Totol: FileCount: %d ; FileSize: %s .\n",totalCount, utils.FormatSize(t) ))

    }

    if (ban.VoicesBucket != "" ){
        fName :=  fmt.Sprintf("%s%s", strings.Join(fileName, ""), "_Voice")
        totalSize := 0
        totalCount := 0

        bucket, err := o.client.Bucket(ban.VoicesBucket)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }

        CreateFile(fName, fmt.Sprintf("GroupID: %d ; Bucket: %s ; Path: %s\n",ban.GrpID, ban.VoicesBucket,ban.VoicesPrefix ))
        CreateFile(fName,partLine + "\n")

        for _, b := range banRes {
            props, err := bucket.GetObjectDetailedMeta(ban.VoicesPrefix + b)
            if err != nil {
                fmt.Println("Error:", err)
                os.Exit(-1)
            }

            Cont := utils.FormatSize(props["Content-Length"][0])
            ContentLength, _ :=  strconv.Atoi(props["Content-Length"][0])

            totalSize += ContentLength
            totalCount += 1

            fmt.Printf("%s | %s\n", Cont, b)
            CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, b))
        }

        t := strconv.Itoa(totalSize)

        CreateFile(fName,partLine + "\n")
        CreateFile(fName,fmt.Sprintf("Totol: FileCount: %d ; FileSize: %s .\n",totalCount, utils.FormatSize(t) ))

    }

    if (ban.VideosBucket != "" ){
        fName :=  fmt.Sprintf("%s%s", strings.Join(fileName, ""), "_Video")
        totalSize := 0
        totalCount := 0

        bucket, err := o.client.Bucket(ban.VideosBucket)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }

        CreateFile(fName, fmt.Sprintf("GroupID: %d ; Bucket: %s ; Path: %s\n",ban.GrpID, ban.VideosBucket,ban.VideosPrefix ))
        CreateFile(fName,partLine + "\n")

        for _, b := range banRes {
            props, err := bucket.GetObjectDetailedMeta(ban.VideosPrefix + b)
            if err != nil {
                fmt.Println("Error:", err)
                os.Exit(-1)
            }

            Cont := utils.FormatSize(props["Content-Length"][0])
            ContentLength, _ :=  strconv.Atoi(props["Content-Length"][0])

            totalSize += ContentLength
            totalCount += 1

            fmt.Printf("%s | %s\n", Cont, b)
            CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, b))
        }

        t := strconv.Itoa(totalSize)

        CreateFile(fName,partLine + "\n")
        CreateFile(fName,fmt.Sprintf("Totol: FileCount: %d ; FileSize: %s .\n",totalCount, utils.FormatSize(t) ))

    }

    if (ban.DocBucket != "" ){
        fName :=  fmt.Sprintf("%s%s", strings.Join(fileName, ""), "_Doc")
        totalSize := 0
        totalCount := 0

        bucket, err := o.client.Bucket(ban.DocBucket)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }

        CreateFile(fName, fmt.Sprintf("GroupID: %d ; Bucket: %s ; Path: %s\n",ban.GrpID, ban.DocBucket,ban.DocPrefix ))
        CreateFile(fName,partLine + "\n")

        for _, b := range banRes {
            props, err := bucket.GetObjectDetailedMeta(ban.DocPrefix + b)
            if err != nil {
                fmt.Println("Error:", err)
                os.Exit(-1)
            }

            Cont := utils.FormatSize(props["Content-Length"][0])
            ContentLength, _ :=  strconv.Atoi(props["Content-Length"][0])

            totalSize += ContentLength
            totalCount += 1

            fmt.Printf("%s | %s\n", Cont, b)
            CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, b))
        }

        t := strconv.Itoa(totalSize)

        CreateFile(fName,partLine + "\n")
        CreateFile(fName,fmt.Sprintf("Totol: FileCount: %d ; FileSize: %s .\n",totalCount, utils.FormatSize(t) ))

    }

    return nil
}

// ListFile ...
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
