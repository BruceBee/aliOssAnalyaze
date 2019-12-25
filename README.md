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
 
