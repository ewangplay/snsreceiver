/*
 1. GET request format
 http://sns.jiuzhilan.com/v1/weibo/appsecret?
     nonce=57155157
     &timestamp=1397022061823
     &echostr=dnPdpTZz85
     &signature=90e4c22c90a58f26526c2dd5b6c56c8822edeaa1

 2. POST request format
 http://sns.jiuzhilan.com/v1/weibo/appsecret?
     nonce=57155157
     &timestamp=1397022061823
     &signature=90e4c22c90a58f26526c2dd5b6c56c8822edeaa1

====================================================================
纯文本类型私信消息
{
    "type": "text",
    "receiver_id": 1902538057,
    "sender_id": 2489518277,
    "created_at": "Mon Jul 16 18:09:20 +0800 2012",
    "text": "私信或留言内容",
    "data": {}
}
说明：
type        string  text
receiver_id int64   消息的接收者
sender_id   int64   消息的发送者
created_at  string  消息创建时间
text        string  私信内容
data        string  消息内容，纯文本私信或留言为空

====================================================================
位置类型私信消息
{
    "type": "position",
    "receiver_id": 1902538057,
    "sender_id": 2489518277,
    "created_at": "Mon Jul 16 18:09:20 +0800 2012",
    "text": "我在这里: http://t.cn/zQgLLYO",
    "data": {
        "longitude": "116.308586",
        "latitude": "39.982525"
    }
}
说明：
type            string  position
receiver_id     int64   消息的接收者
sender_id       int64   消息的发送者
created_at      string  消息创建时间
text            string  原位置私信文本，没有时用默认文案“发送了一个位置”
data            string  消息内容
data:longitude  string  经度
data:latitude   string  纬度

====================================================================
语音类型私信消息
{
    "type": "voice",
    "receiver_id": 1902538057,
    "sender_id": 2489518277,
    "created_at": "Mon Jul 16 18:09:20 +0800 2012",
    "text": "发了一个语音消息",
    "data": {
        "vfid": 821804459,    // 发送者用此ID查看语音
        "tovfid": 821804469  // 接收者用此ID查看语音
    }
}
说明：
type        string  voice
receiver_id int64   消息的接收者
sender_id   int64   消息的发送者
created_at  string  消息创建时间
text        string  私信内容
data        string  消息内容，纯文本私信或留言为空
data:vfid   string  语音文件ID，发送者通过此ID读取语音
data:tovfid string  语音文件ID，接收者通过此ID读取语音

====================================================================
图片类型私信消息
{
    "type": "image",
    "receiver_id": 1902538057,
    "sender_id": 2489518277,
    "created_at": "Mon Jul 16 18:09:20 +0800 2012",
    "text": "发了一张图片",
    "data": {
        "vfid": 821804459,     // 发送者用此ID查看图片
        "tovfid": 821804469    // 接收者用此ID查看图片
    }
}
说明：
type        string  image
receiver_id int64   消息的接收者
sender_id   int64   消息的发送者
created_at  string  消息创建时间
text        string  私信内容
data        string  消息内容，纯文本私信或留言为空
data:vfid   string  图片ID，发送者通过此ID读取图片
data:tovfid string  图片ID，接收者通过此ID读取图片
*/
package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	WEIBO_CHECK_TOKEN_SQL string = "SELECT cid, account_id FROM jzl_weibo_account WHERE appsecret=? AND is_delete=0"
	GET_CUSTOMER_ID_SQL   string = "SELECT cid FROM jzl_weibo_account WHERE account_id=? AND is_delete=0"
)

type WeiboPrivateMsg struct {
	CustomerId  int
	AccountId   string
	Type        string
	Receiver_id int64
	Sender_id   int64
	Created_at  string
	Text        string
	Data        map[string]string
}

type WeiboProcessor struct {
}

func (this *WeiboProcessor) ProcessFunc(version int, resource string, method string, token string, params map[string]string, body []byte, result map[string]interface{}) error {

	if !CheckAuth(token, params) {
		//return fmt.Errorf("Check Auth Fail")
	}

	switch method {
	case "GET":
		//微博服务器通过GET请求建立连接
		var echostr string
		echostr, ok := params["echostr"]
		if !ok {
			LOG_ERROR("请求中不包含echostr信息")
			return fmt.Errorf("no echostr info")
		}
		if echostr == "" {
			LOG_ERROR("echostr信息为空")
			return fmt.Errorf("echostr info is null")
		}

		LOG_DEBUG("echostr: %v", echostr)

		LOG_INFO("签名信息验证通过，回复%v跟微博服务器建立连接", echostr)

		//所有验证通过，回复echostr建立连接
		result["echo"] = echostr

	case "POST":
		/*
			//根据token信息查找私信对应的cid和account_id
			cid, acc_id, err := this.CheckToken(token)
			if err != nil {
				LOG_ERROR("检查token失败: %v", err)
				return err
			}

			LOG_DEBUG("cid=%v, account_id=%v", cid, acc_id)
		*/

		//对接收到微信私信进行解析
		msg := &WeiboPrivateMsg{}
		err := json.Unmarshal(body, msg)
		if err != nil {
			LOG_ERROR("解析消息体失败: %v", err)
			return err
		}

		msg.AccountId = strconv.FormatInt(msg.Receiver_id, 10)
		msg.CustomerId, err = this.GetCustomerId(msg.AccountId)
		if err != nil {
			LOG_ERROR("根据微博账号[%v]获取客户ID失败: %v", err)
			return err
		}

		LOG_DEBUG("msg=%v", msg)

		//处理后的微信消息做JSON编码
		json_msg, err := json.Marshal(msg)
		if err != nil {
			LOG_ERROR("对微信私信做JSON编码失败: %v", err)
			return err
		}

		//微博服务器通过POST请求推送消息
		topic := "WeiboPrivateMsg"
		err = g_nsqProducer.Publish(topic, json_msg)
		if err != nil {
			LOG_ERROR("发布消息[%v]到NSQ主题[%v]失败，失败原因：%v", string(json_msg), topic, err)
			return err
		}

		LOG_INFO("发布消息[%v]到NSQ主题[%v]成功", string(json_msg), topic)

		//假如开发者无法保证在五秒内处理并回复，可以直接回复空串，
		//微博服务器不会对此作任何处理，并且不会发起重试
		result["echo"] = ""

	default:
		//其它请求不处理
		LOG_INFO("接收到不符合要求的请求[%v]，忽略掉", method)
		return nil
	}

	return nil
}

func (this *WeiboProcessor) CheckToken(token string) (customer_id int, account_id string, err error) {
	row, err := g_mysqladaptor.QueryFormat(WEIBO_CHECK_TOKEN_SQL, token)
	if err != nil {
		return
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&customer_id, &account_id)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("token %v invalid", token)
		return
	}

	return
}
func (this *WeiboProcessor) GetCustomerId(account_id string) (customer_id int, err error) {
	row, err := g_mysqladaptor.QueryFormat(GET_CUSTOMER_ID_SQL, account_id)
	if err != nil {
		return
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&customer_id)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("account_id[%v] not found", account_id)
		return
	}

	return
}
