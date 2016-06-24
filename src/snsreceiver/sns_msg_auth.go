/*
 1. 在开发者首次使用事件推送服务时，需要先通过一次校验来和微博服务器建立首次连接，具体来说：
     开发者提交信息后，微博消息服务器将发送GET请求到填写的URL上，GET请求携带四个参数：

     校验参数字段   字段类型    字段说明
     signature      string      微博加密签名，signature结合了开发者的token、和请求中的timestamp参数，nonce参数
     timestamp      string      时间戳
     nonce          string      随机数
     echostr        string      随机字符串

 2. signature参数的加密规则为：
     将开发者的token，timestamp参数，nonce参数进行字典排序后，将三个参数字符串拼接成一个字符串进行sha1加密 校验参数：
     appsercret=xyz123xyz        timestamp=1397022061823        nonce=57155157
     加密结果：
     拼接后的字符串为：139702206182357155157xyz123xyz
     sha1签名后的结果为：90e4c22c90a58f26526c2dd5b6c56c8822edeaa1
     验证url有效性请求的样例为： http://yoururl?nonce=57155157&timestamp=1397022061823&echostr=dnPdpTZz85&signature=90e4c22c90a58f26526c2dd5b6c56c8822edeaa1
     此时如果返回的是echostr的值（此样例中为dnPdpTZz85）则通过url验证。
*/
package main

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

func CheckAuth(token string, params map[string]string) bool {
	var signature string
	var timestamp string
	var nonce string
	var ok bool

	//get signature param
	signature, ok = params["signature"]
	if !ok {
		LOG_ERROR("请求中不包含signature信息")
		return false
	}
	if signature == "" {
		LOG_ERROR("signature信息为空")
		return false
	}

	LOG_DEBUG("signature: %v", signature)

	//get timestamp param
	timestamp, ok = params["timestamp"]
	if !ok {
		LOG_ERROR("请求中不包含timestamp信息")
		return false
	}
	if timestamp == "" {
		LOG_ERROR("timestamp信息为空")
		return false
	}

	LOG_DEBUG("timestamp: %v", timestamp)

	//get nonce param
	nonce, ok = params["nonce"]
	if !ok {
		LOG_ERROR("请求中不包含nonce信息")
		return false
	}
	if nonce == "" {
		LOG_ERROR("nonce信息为空")
		return false
	}

	LOG_DEBUG("nonce: %v", nonce)

	//check the signature
	var mysig string
	tmpArr := []string{token, timestamp, nonce}
	sort.Strings(tmpArr)
	tmpStr := strings.Join(tmpArr, "")
	sha1Data := sha1.Sum([]byte(tmpStr))
	for i := 0; i < sha1.Size; i++ {
		mysig += fmt.Sprintf("%x", sha1Data[i])
	}

	LOG_DEBUG("mysig: %v", mysig)

	if strings.TrimSpace(strings.ToLower(mysig)) != strings.TrimSpace(strings.ToLower(signature)) {
		LOG_ERROR("请求的签名信息不匹配")
		return false
	}

	return true
}
