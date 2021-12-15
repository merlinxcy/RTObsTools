package config

import (
	"testobs/lib/aliyuncloud"
	"testobs/lib/huaweicloud"
)

var HuaweiCloudConfig *huaweicloud.Config

var AliyunCloudConfig *aliyuncloud.Config

func init(){
	HuaweiCloudConfig = &huaweicloud.Config{
		AK:         "",
		SK:         "",
		BucketName: "kali8",
		OBSURL:     "",
		FileKey:    "",
	}

	AliyunCloudConfig = &aliyuncloud.Config{
		AK:         "",
		SK:         "",
		BucketName: "kali8",
		OBSURL:     "",
		FileKey:    "",
	}
}
