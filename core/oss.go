/*
@Author : Bruce Bee
@Date   : 2019/12/17 10:17
@Email  : mzpy_1119@126.com
*/

// The core of Project
// 1、 Define osser interface and implement three methods respectively.
//    `ReturnSize` is used for thread scheduling
//    `sizeCalc` for file size statistics and `ListFile` for display a list of files.
//
// 2、 OSS-based SDK define structures whose primary purpose is to use SDK objects to query files.
//
// 3、 Goroutine is used for multi-threaded file queries and SDK calls,
//    `chan` is used for communication between threads,
//    and the main thread waitsfor execution using the `WaitGroup` method of the `sync` package.

package core

import (
    "../utils"
    "os"
    "strconv"
    "sync"
    "time"
    "fmt"
    "github.com/Unknwon/goconfig"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Osser ...
type Osser interface {
    ReturnSize(int64) error
    sizeCalc(BaseInfo, string, map[string]map[string]int)
    ListFile()
}

// OSS ...
type OSS struct {
    client *oss.Client
}

// register ...
func register(groupID int64, r chan <- BaseInfo, wg *sync.WaitGroup){

    var registerList []func(groupID int64) []BaseInfo
    registerList  = append(registerList, QueryBanner)
    registerList  = append(registerList, QueryCard)
    registerList  = append(registerList, QueryCardChapter)
    registerList  = append(registerList, QueryCardQuestion)
    registerList  = append(registerList, QueryColumnAnswer)


    for _, f := range registerList {
        res := f(groupID)
        for _, obj := range res {
           r <- obj
        }
    }
    wg.Done()
    close(r)
}


// fileCalc ...
func fileCalc(groupID int64, fc <- chan BaseInfo, wg *sync.WaitGroup, o *OSS, total map[string]map[string]int){

    for {
       fileObj := <- fc
       if (fileObj.GrpID == 0){
           break
       }
       o.sizeCalc(fileObj, fileObj.TableName, total )
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

    partLine := partLine()

    totalData := map[string]map[string]int{}
    wg := &sync.WaitGroup{}
    ch := make(chan BaseInfo, 1000)
    wg.Add(2)
    go register(groupID, ch, wg)
    go fileCalc(groupID, ch, wg, o, totalData)

    time.Sleep(2 * time.Second)
    wg.Wait()

    for t := range totalData {
        ts := strconv.Itoa(totalData[t]["totalSize"])

        CreateFile(t, partLine + "\n")
        CreateFile(t, fmt.Sprintf("Total: FileCount: %d ; FileSize: %s .\n",totalData[t]["totalCount"], utils.FormatSize(ts) ))
    }
    return nil
}

// SizeCalc ...
func (o *OSS) sizeCalc(info BaseInfo, fileName string, total map[string]map[string]int){

    partLine := partLine()

    if (info.PicBucket != "" ) {
        fName := fmt.Sprintf("%s%s", fileName, "_Pic")

        if (total[info.TableName + "_Pic"] == nil) {
            subMapB := make(map[string]int)
            total[info.TableName + "_Pic"] = subMapB

            CreateFile(fName, fmt.Sprintf("GroupID: %d ; Bucket: %s ; Path: %s\n",info.GrpID, info.PicBucket,info.PicPrefix ))
            CreateFile(fName,partLine + "\n")

        }

        bucket, err := o.client.Bucket(info.PicBucket)
        if err != nil {
            fmt.Println("BucketError:", err)
            //os.Exit(-1)
        }else {
            props, err := bucket.GetObjectDetailedMeta(info.PicPrefix + info.PicURL)
            if err != nil {
                fmt.Println("ObjectError:", err)
                //os.Exit(-1)
            }else {

                Cont := utils.FormatSize(props["Content-Length"][0])
                ContentLength, _ := strconv.Atoi(props["Content-Length"][0])

                total[info.TableName+"_Pic"]["totalSize"] += ContentLength
                total[info.TableName+"_Pic"]["totalCount"] ++

                fmt.Printf("%s | %s\n", Cont, info.PicURL)
                CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, info.PicURL))
            }

        }
    }


    if (info.VoiceBucket != "" ){
        fName :=  fmt.Sprintf("%s%s", fileName,  "_Voice")

        if (total[info.TableName + "_Voice"] == nil) {
            subMapB := make(map[string]int)
            total[info.TableName + "_Voice"] = subMapB

            CreateFile(fName, fmt.Sprintf("GroupID: %d ; Bucket: %s ; Path: %s\n",info.GrpID, info.VoiceBucket, info.VoicePrefix ))
            CreateFile(fName, partLine + "\n")

        }

        bucket, err := o.client.Bucket(info.VoiceBucket)
        if err != nil {
            fmt.Println("BucketError:", err)
            //os.Exit(-1)
        }else {
            props, err := bucket.GetObjectDetailedMeta(info.VoicePrefix + info.VoiceURL)
            if err != nil {
                fmt.Println("ObjectError:", err)
                //os.Exit(-1)
            }else {
                Cont := utils.FormatSize(props["Content-Length"][0])
                ContentLength, _ := strconv.Atoi(props["Content-Length"][0])

                total[info.TableName + "_Voice"]["totalSize"] += ContentLength
                total[info.TableName + "_Voice"]["totalCount"] ++

                fmt.Printf("%s | %s\n", Cont, info.VoiceURL)
                CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, info.VoiceURL))
            }
        }
    }


    if (info.VideoBucket != "" ){
        fName :=  fmt.Sprintf("%s%s", fileName, "_Video")

        if (total[info.TableName + "_Video"] == nil) {
            subMapB := make(map[string]int)
            total[info.TableName + "_Video"] = subMapB

            CreateFile(fName, fmt.Sprintf("GroupID: %d ; Bucket: %s ; Path: %s\n",info.GrpID, info.VideoBucket, info.VideoPrefix ))
            CreateFile(fName, partLine + "\n")

        }

        bucket, err := o.client.Bucket(info.VideoBucket)
        if err != nil {
            fmt.Println("BucketError:", err)
            //os.Exit(-1)
        }else {
            props, err := bucket.GetObjectDetailedMeta(info.VideoPrefix + info.VideoURL)
            if err != nil {
                fmt.Println("ObjectError:", err)
                //os.Exit(-1)
            }else {
                Cont := utils.FormatSize(props["Content-Length"][0])
                ContentLength, _ := strconv.Atoi(props["Content-Length"][0])

                total[info.TableName + "_Video"]["totalSize"] += ContentLength
                total[info.TableName + "_Video"]["totalCount"] ++

                fmt.Printf("%s | %s\n", Cont, info.VideoURL)
                CreateFile(fName, fmt.Sprintf("%s | %s \n", Cont, info.VideoURL))
            }
        }
    }

    /*
    if (info.DocBucket != "" ){
        fName :=  fmt.Sprintf("%s%s", fileName, "_Doc")
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


// partLint
func partLine() (partline string) {

    for i := 0; i < 60; i++ {
        partline += "-"
    }
    return
}