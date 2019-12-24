/*
@Author : Bruce Bee
@Date   : 2019/12/17 10:17
@Email  : mzpy_1119@126.com
*/
package main

import (
    "./core"
    "flag"
    "fmt"
)

func main() {

    var groupID int64
    flag.Int64Var(&groupID, "g", 0,"组织ID")
    flag.Parse()

    if (groupID == 0){
        fmt.Printf("Usage of ./aliOssAnalyaze:\n  -g int\n        组织ID\n")
        return
    }

    //cfg, _ := goconfig.LoadConfigFile("conf/app.ini")
    //bucketName, _ := cfg.GetValue("oss","bucket")

     // core.InitDB()
     a := core.InitOSS()
     a.ReturnSize(groupID)
     //a.ListFile()
}
