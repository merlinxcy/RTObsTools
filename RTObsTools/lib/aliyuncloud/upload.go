package aliyuncloud

import (
	"gitlab.com/mjwhitta/win/wininet/http"
	"io/ioutil"
	"log"
	"testobs/common"
	"time"
)

func UploadObs(file_path string, dst_bucket string, conf *Config){
	var res *http.Response
	var b []byte
	var e error

	timeNow := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	var header = map[string]string{
		"Date": timeNow,
		"Content-Type": "text/plain",
		"Authorization": CalcSign(conf.AK, conf.SK, "PUT","text/plain",
			timeNow,"/" + conf.BucketName + "/" + dst_bucket),
	}
	http.DefaultClient.TLSClientConfig.InsecureSkipVerify = true
	req := http.NewRequest(http.MethodPut, conf.OBSURL+"/" + dst_bucket,GetFileContent(file_path, conf.FileKey))
	req.Headers = header
	if res, e = http.DefaultClient.Do(req); e != nil{
		panic(e)
	}
	if res.Body != nil {
		if b, e = ioutil.ReadAll(res.Body); e != nil {
			panic(e)
		}
	}
	log.Println(string(b))
	log.Println(res.Status)
}

func DownloadObs(src_bucket string, conf *Config){
	var res *http.Response
	var b []byte
	var e error

	timeNow := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	var header = map[string]string{
		"Date": timeNow,
		"Content-Type": "text/plain",
		"Authorization": CalcSign(conf.AK, conf.SK, "GET","text/plain",
			timeNow,"/" + conf.BucketName + "/" + src_bucket),
	}
	http.DefaultClient.TLSClientConfig.InsecureSkipVerify = true
	req := http.NewRequest(http.MethodGet, conf.OBSURL+"/" + src_bucket,nil)
	req.Headers = header
	if res, e = http.DefaultClient.Do(req); e != nil{
		panic(e)
	}
	if res.Body != nil {
		if b, e = ioutil.ReadAll(res.Body); e != nil {
			panic(e)
		}
	}
	log.Println("StatusCode: ", res.Status)
	//log.Println("Content", b)
	if b !=nil && res.StatusCode == 200{
		DownloadFileContent(src_bucket, b, conf.FileKey)
	}
}


func GetFileContent(file_path , asekey string) []byte{
	con, err := ioutil.ReadFile(file_path)
	if err != nil{
		panic(err)
	}
	ret := common.EncryptAES(con, []byte(asekey))
	//base64.StdEncoding.EncodeToString("")
	return ret
}

func DownloadFileContent(filename string, filecon []byte, asekey string){
	err := ioutil.WriteFile(filename, common.DecryptAES(filecon, []byte(asekey)), 0777)
	if err != nil{
		log.Println(err)
	}
}