package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testobs/config"
	"testobs/lib/aliyuncloud"
	"testobs/lib/huaweicloud"
	"time"
)

var Mode string
var FilePath string
var BucketName string

func init(){
	flag.StringVar(&Mode, "m", "", "模式：upload/download")
	flag.StringVar(&FilePath, "p", "", "地址：path")
	flag.StringVar(&BucketName, "b", "", "桶：Bucket")
}

func main(){
	flag.Parse()
	fmt.Println("[+]Time: ",time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
	conf := config.HuaweiCloudConfig
	fmt.Println(FilePath, BucketName)
	if Mode == "upload" && FilePath != ""{
		base_name := filepath.Base(FilePath)
		if _, err := os.Stat(FilePath);err !=nil{
			fmt.Println("file not exists")
			os.Exit(-1)
		}
		huaweicloud.UploadObs(FilePath, base_name, conf)
	} else if Mode == "download" && BucketName != ""{
		huaweicloud.DownloadObs(BucketName, conf)
	} else {
		loggg := `用法：
上传：obs.exe -m upload -p d:\1.txt
下载：obs.exe -m download -b 1.txt
`
		log.Fatal(loggg)
	}
	//huaweicloud.UploadObs("d:\\1.txt", "1.txt", &huaweicloud.Config{
	//	AK:         "",
	//	SK:         "",
	//	BucketName: "kali8",
	//	OBSURL:     "",
	//	FileKey: "",
	//})

	//huaweicloud.DownloadObs("1.txt", conf)

}


func mai1n(){
	conf := config.AliyunCloudConfig
	aliyuncloud.UploadObs("D:\\1.txt", "1.txt", conf)
	aliyuncloud.DownloadObs("1.txt",conf)
}