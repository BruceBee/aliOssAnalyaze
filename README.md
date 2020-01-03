## aliOssAnalyaze

![](https://img.shields.io/badge/Go-1.13.1-brightgreen.svg) ![](https://img.shields.io/badge/License-MIT-orange.svg)


## Introduce

In some cases, we use the services of cloud platforms for file storage, such as aliyun, AWS and tencent cloud. Inevitably, we will scatter the file data everywhere.

Therefore, we need to summarize and count the same files according to some rules, such as I need to count the files uploaded by a user, what are the names of these files, where they are stored and so on.

Here we use a database approach to query the rules:

- First, we can query the path information of the file according to the rules, including bucket, directory, file name and so on.
- By calling the SDK interface of OSS of aliyun, calculate the actual size of the file, upload time, file permission and other information.
- Summary of the number of statistical files, total size, etc.
- Output to a file with a custom filename based on the file information

Of course, you can also adjust the code according to your own rules.



## Precondition
You must ensure that you have obtained the keys of aliyun OSS, including `access_key` and `access_secret`, and that the permissions of the keys are available.

You also need to fill in the database connection information and the OSS key to the `conf/app.ini`.

## Configure
In the configuration file, the items in `os-info` respectively represent four file types, `pic` for picture, `voice` for voice, `video` for video, and `doc` for document.

Meanwhile, in the `utils\utils.go` module, we set the file suffix type to include these `jpg、png、gif、mp3、silk、mp4、m3u8、txt、rtf、doc、docx、xls、xlsx、ppt、pptx、pdf`，and you can adjust according to the actual situation.

In `core\module`, By default, we set a module corresponding to an SQL query and tried to parse it. You can cut it according to your own needs. My parsing method may not be completely suitable for your business scenario.

By default, we set a module corresponding to an SQL query and tried to parse it. You can cut it according to your own needs. My parsing method may not be completely suitable for your business scenario.

It should be noted that each module may produce files of the above four types. We will display files of the same type in the same file and calculate the total capacity of files of the same class.

Four different types of files are named as follows:

    `ModuleFileName_Pic` 、 `ModuleFileName_Voice`、`ModuleFileName_Video` or `ModuleFileName_Doc`
    
    


## Usage
```shell script
> git clone git@github.com:BruceBee/aliOssAnalyaze.git
> cd aliOssAnalyaze
> go build
> ./aliOssAnalyaze -h

Usage of ./aliOssAnalyaze:
  -g int
      Group_ID

> ./aliOssAnalyaze -g 123

```
 
