package aliyuncloud

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func CalcSign(ak, sk , http_verb, content_type, date , resource string) string{
	string_to_sign := http_verb + "\n" + "\n"+
		content_type + "\n" +
		date + "\n" +
		resource

	mac := hmac.New(sha1.New, []byte(sk))
	_, err := mac.Write([]byte(string_to_sign))
	if err != nil{
		panic("hmac sha1 failed")
	}

	return "OSS "+ ak + ":" +base64.StdEncoding.EncodeToString(mac.Sum(nil))
}