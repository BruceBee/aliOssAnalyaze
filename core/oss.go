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
    "sync"
    "time"
)

// Osser ...
type Osser interface {
    ReturnSize(int64) error
    SizeCalc(BaseInfo, []string, map[string]map[string]int)
    ListFile()

}


// OSS ...
type OSS struct {
    client *oss.Client
}

// Register ...
func Register(groupID int64, r chan <- BaseInfo, wg *sync.WaitGroup){
    ban := QueryBanner(groupID)
    for _, b := range ban {
        r <- b
    }

    card := QueryCard(groupID)
    for _, c := range card {
        r <- c
    }


    wg.Done()
    close(r)

}


// FileCalc ...
func FileCalc(groupID int64, fc <- chan BaseInfo, wg *sync.WaitGroup, o *OSS){

    totalData := map[string]map[string]int{}

    for {
       fileObj := <- fc
       if (fileObj.GrpID == 0){
           break
       }
       //f := fmt.Sprintf("GroupID: %d; PicURL: %s; VoiceURL: %s; VideoURL: %s; TABLE:%s ", fileObj.GrpID, fileObj.PicURL, fileObj.VoiceURL, fileObj.VideoURL, fileObj.TableName)
       //fmt.Println(f)

        fileName := []string{}
        fileName = append(fileName, filePath)
        fileName = append(fileName, fileObj.TableName)

        o.SizeCalc(fileObj, fileName, totalData )
    }
    wg.Done()
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
func (o *OSS) ReturnSize(groupID int64) error {

    wg := &sync.WaitGroup{}

    ch := make(chan BaseInfo, 1000)

    wg.Add(2)

    go Register(groupID, ch, wg)
    go FileCalc(groupID, ch, wg, o)

    time.Sleep(2 * time.Second)
    wg.Wait()

    return nil
}

// SizeCalc ...
func (o *OSS) SizeCalc(info BaseInfo, fileName []string, total map[string]map[string]int){

    // partLint
    partLine := "-"
    for i := 0; i < 60; i++ {
        partLine += "-"
    }

    if (info.PicBucket != "" ) {
        fName := fmt.Sprintf("%s%s", strings.Join(fileName, ""), "_Pic")

        bucket, err := o.client.Bucket(info.PicBucket)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }

        props, err := bucket.GetObjectDetailedMeta(info.PicPrefix + info.PicURL)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }

        Cont := utils.FormatSize(props["Content-Length"][0])
        ContentLength, _ := strconv.Atoi(props["Content-Length"][0])

        subMapB := make(map[string]int)
        subMapB["totalSize"] = ContentLength
        subMapB["totalCount"] = 1
        total[fName] = subMapB

        fmt.Printf("%s | %s\n", Cont, info.PicURL)
        CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, info.PicURL))
    }



    if (info.VoiceBucket != "" ){
        fName :=  fmt.Sprintf("%s%s", strings.Join(fileName, ""), "_Voice")

        bucket, err := o.client.Bucket(info.VoiceBucket)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }

        props, err := bucket.GetObjectDetailedMeta(info.VoicePrefix + info.VoiceURL)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }

        Cont := utils.FormatSize(props["Content-Length"][0])
        ContentLength, _ :=  strconv.Atoi(props["Content-Length"][0])

        subMapB := make(map[string]int)
        subMapB["totalSize"] = ContentLength
        subMapB["totalCount"] = 1
        total[fName] = subMapB

        fmt.Printf("%s | %s\n", Cont, info.VoiceURL)
        CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, info.VoiceURL))
    }

    /*

    if (info.VideoBucket != "" ){
        fName :=  fmt.Sprintf("%s%s", strings.Join(fileName, ""), "_Video")
        totalSize := 0
        totalCount := 0

        bucket, err := o.client.Bucket(ban.VideoBucket)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }

        CreateFile(fName, fmt.Sprintf("GroupID: %d ; Bucket: %s ; Path: %s\n",ban.GrpID, ban.VideoBucket,ban.VideoPrefix ))
        CreateFile(fName,partLine + "\n")

        for _, b := range banRes {
            props, err := bucket.GetObjectDetailedMeta(ban.VideoPrefix + b)
            if err != nil {
                fmt.Println("Error:", err)
                os.Exit(-1)
            }

            Cont := utils.FormatSize(props["Content-Length"][0])
            ContentLength, _ :=  strconv.Atoi(props["Content-Length"][0])

            totalSize += ContentLength
            totalCount++

            fmt.Printf("%s | %s\n", Cont, b)
            CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, b))
        }

        t := strconv.Itoa(totalSize)

        CreateFile(fName,partLine + "\n")
        CreateFile(fName,fmt.Sprintf("Total: FileCount: %d ; FileSize: %s .\n",totalCount, utils.FormatSize(t) ))

    }

    if (info.DocBucket != "" ){
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
            totalCount++

            fmt.Printf("%s | %s\n", Cont, b)
            CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, b))
        }

        t := strconv.Itoa(totalSize)

        CreateFile(fName,partLine + "\n")
        CreateFile(fName,fmt.Sprintf("Total: FileCount: %d ; FileSize: %s .\n",totalCount, utils.FormatSize(t) ))

    }

     */
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
